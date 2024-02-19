package notification

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
)

type Sender interface {
	SendNotification(ctx context.Context, notification *model.Notification, offer *model.Offer) error
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

func (r *SenderRegistry) SendNotification(ctx context.Context, notification *model.Notification, offer *model.Offer) error {
	for _, sender := range r.senders {
		if err := sender.SendNotification(ctx, notification, offer); err != nil {
			return err
		}
	}
	return nil
}
