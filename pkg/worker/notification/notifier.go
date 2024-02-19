package notification

import (
	"context"
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/condition"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"log/slog"
)

var AllNotificationsFailedError = errors.New("all notifications failed")

type Notifier struct {
	conditionRepo        db.ConditionRepository
	notificationRepo     db.NotificationRepository
	sentNotificationRepo db.SentNotificationRepository

	conditionRegistry *condition.Registry
	senderRegistry    *SenderRegistry
}

func NewNotifier(d db.DB, conditionRegistry *condition.Registry, senderRegistry *SenderRegistry) *Notifier {
	return &Notifier{
		conditionRepo:        d.GetConditionRepository(),
		notificationRepo:     d.GetNotificationRepository(),
		sentNotificationRepo: d.GetSentNotificationRepository(),
		conditionRegistry:    conditionRegistry,
		senderRegistry:       senderRegistry,
	}
}

func (n *Notifier) TrySendNotification(ctx context.Context, offer *model.Offer, action model.OfferAction) error {

	skip := int64(0)
	succeed := false

	for {
		notifs, total, err := n.notificationRepo.FindAll(ctx, 100, skip)

		if err != nil {
			return err
		}

		skip += 100
		if skip > total {
			break
		}

		for _, notif := range notifs {
			sentNotif, err := n.processNotification(ctx, notif, offer, action)

			if err != nil {
				slog.Error("failed to process notification", "error", err.Error())
				continue
			}

			err = n.sentNotificationRepo.Insert(ctx, sentNotif, notif, offer)

			if err != nil {
				slog.Error("failed to insert sent notification", "error", err.Error())
				continue
			}

			if sentNotif != nil {
				if n.sendNotification(ctx, sentNotif, offer) {
					succeed = true
				}
			} else {
				succeed = true
			}
		}
	}

	if !succeed {
		return AllNotificationsFailedError
	}

	return nil
}

func (n *Notifier) processNotification(ctx context.Context, notif *model.Notification, offer *model.Offer, action model.OfferAction) (*model.SentNotification, error) {
	cond, err := n.conditionRepo.FindById(ctx, notif.ConditionId)
	if err != nil {
		return nil, err
	}

	wrapper, err := n.conditionRegistry.Get(ctx, cond)
	if err != nil {
		slog.Error("failed to get condition wrapper", "error", err.Error())
		return nil, err
	}

	return wrapper.CheckOffer(ctx, offer, action, notif.Config)
}

func (n *Notifier) sendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) bool {
	sentNotif, err := n.senderRegistry.SendNotification(ctx, sentNotif, offer)

	updateErr := n.sentNotificationRepo.Update(ctx, sentNotif)

	if updateErr != nil {
		slog.Error("failed to update sent notification", "error", updateErr.Error())
		return false
	}

	if err != nil {
		slog.Error("failed to send notification", "error", err.Error())
		return false
	}

	return sentNotif.SentSuccessfully
}
