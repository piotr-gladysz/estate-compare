package notification

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
)

// NotificationChannel is an interface for sending notifications by single channel
type NotificationChannel interface {

	// SendNotification sends notification to the user and returns error if any
	SendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) error

	// GetName returns name of the channel for identification
	GetName() string
}

type ChannelRegistry struct {
	channels []NotificationChannel
}

func NewChannelRegistry() *ChannelRegistry {
	return &ChannelRegistry{channels: make([]NotificationChannel, 0)}
}

func (r *ChannelRegistry) Register(channel NotificationChannel) {
	r.channels = append(r.channels, channel)
}

func (r *ChannelRegistry) SendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) (*model.SentNotification, error) {

	allSuccess := true

	for _, channel := range r.channels {

		if r.checkStatus(sentNotif, channel.GetName()) {
			continue
		}

		err := channel.SendNotification(ctx, sentNotif, offer)
		if err != nil {
			r.addError(sentNotif, channel.GetName(), err)
			allSuccess = false
		} else {
			r.setSuccess(sentNotif, channel.GetName())
		}
	}

	sentNotif.SentSuccessfully = allSuccess

	return sentNotif, nil
}

func (r *ChannelRegistry) checkStatus(sentNotif *model.SentNotification, channelName string) bool {

	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if channelStatus, ok := sentNotif.SendingStatus[channelName]; ok {
		if status, ok := channelStatus.(map[string]interface{}); ok {
			if sent, ok := status["sent"].(bool); ok && sent {
				return true
			}
		}
	}
	return false
}

func (r *ChannelRegistry) setSuccess(sentNotif *model.SentNotification, channelName string) *model.SentNotification {
	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if _, ok := sentNotif.SendingStatus[channelName].(map[string]any); !ok {
		sentNotif.SendingStatus[channelName] = map[string]any{
			"sent":   false,
			"errors": []string{},
		}
	}
	sentNotif.SendingStatus[channelName] = map[string]interface{}{"sent": true}

	return sentNotif
}

func (r *ChannelRegistry) addError(sentNotif *model.SentNotification, channelName string, err error) *model.SentNotification {
	if sentNotif.SendingStatus == nil {
		sentNotif.SendingStatus = make(map[string]any)
	}

	if _, ok := sentNotif.SendingStatus[channelName].(map[string]any); !ok {
		sentNotif.SendingStatus[channelName] = map[string]any{
			"sent":   false,
			"errors": []string{},
		}
	}
	errorsArr := sentNotif.SendingStatus[channelName].(map[string]any)["errors"].([]string)
	errorsArr = append(errorsArr, err.Error())
	sentNotif.SendingStatus[channelName].(map[string]any)["errors"] = errorsArr

	return sentNotif

}
