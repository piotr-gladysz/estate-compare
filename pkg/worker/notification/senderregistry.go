package notification

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
)

type Sender interface {
	SendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) error
	GetName() string
}

type SenderRegistry struct {
	senders []Sender
}

func NewSenderRegistry() *SenderRegistry {
	return &SenderRegistry{senders: make([]Sender, 0)}
}

func (r *SenderRegistry) Register(sender Sender) {
	r.senders = append(r.senders, sender)
}

func (r *SenderRegistry) SendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) (*model.SentNotification, error) {

	allSuccess := true

	for _, sender := range r.senders {

		if r.checkStatus(sentNotif, sender.GetName()) {
			continue
		}

		err := sender.SendNotification(ctx, sentNotif, offer)
		if err != nil {
			r.addError(sentNotif, sender.GetName(), err)
			allSuccess = false
		} else {
			r.setSuccess(sentNotif, sender.GetName())
		}
	}

	sentNotif.SentSuccessfully = allSuccess

	return sentNotif, nil
}

func (r *SenderRegistry) checkStatus(sentNotif *model.SentNotification, senderName string) bool {

	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if senderStatus, ok := sentNotif.SendingStatus[senderName]; ok {
		if status, ok := senderStatus.(map[string]interface{}); ok {
			if sent, ok := status["sent"].(bool); ok && sent {
				return true
			}
		}
	}
	return false
}

func (r *SenderRegistry) setSuccess(sentNotif *model.SentNotification, senderName string) *model.SentNotification {
	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if _, ok := sentNotif.SendingStatus[senderName].(map[string]any); !ok {
		sentNotif.SendingStatus[senderName] = map[string]any{
			"sent":   false,
			"errors": []string{},
		}
	}
	sentNotif.SendingStatus[senderName] = map[string]interface{}{"sent": true}

	return sentNotif
}

func (r *SenderRegistry) addError(sentNotif *model.SentNotification, senderName string, err error) *model.SentNotification {
	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if _, ok := sentNotif.SendingStatus[senderName].(map[string]any); !ok {
		sentNotif.SendingStatus[senderName] = map[string]any{
			"sent":   false,
			"errors": []string{},
		}
	}
	errorsArr := sentNotif.SendingStatus[senderName].(map[string]any)["errors"].([]string)
	errorsArr = append(errorsArr, err.Error())
	sentNotif.SendingStatus[senderName].(map[string]any)["errors"] = errorsArr

	return sentNotif

}
