package notification

import (
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"reflect"
	"testing"
)

func TestSenderRegistry_checkStatus(t *testing.T) {
	tests := []*model.SentNotification{
		{
			SendingStatus: map[string]interface{}{
				"sender1": map[string]interface{}{
					"sent": true,
				},
			},
		},
		{
			SendingStatus: map[string]interface{}{
				"sender1": map[string]interface{}{
					"sent": false,
				},
			},
		},
		{
			SendingStatus: nil,
		},
	}

	results := []bool{true, false, false}

	sr := NewSenderRegistry()

	for i, test := range tests {
		if sr.checkStatus(test, "sender1") != results[i] {
			t.Errorf("Expected %v, got %v", results[i], sr.checkStatus(test, "sender1"))
		}
	}
}

func TestSenderRegistry_addError(t *testing.T) {
	tests := []*model.SentNotification{
		{
			SendingStatus: map[string]interface{}{
				"sender1": map[string]interface{}{
					"sent":   false,
					"errors": []string{"error0"},
				},
			},
		},
		{
			SendingStatus: nil,
		},
	}

	results := [][]string{{"error0", "error1"}, {"error1"}}

	sr := NewSenderRegistry()

	for i, test := range tests {
		sr.addError(test, "sender1", errors.New("error1"))
		if !reflect.DeepEqual(results[i], test.SendingStatus["sender1"].(map[string]interface{})["errors"].([]string)) {
			t.Errorf("Expected %v, got %v", results[i], test.SendingStatus["sender1"].(map[string]interface{})["errors"].([]string))
		}
	}

}

func TestSenderRegistry_setSuccess(t *testing.T) {
	tests := []*model.SentNotification{
		{
			SendingStatus: map[string]interface{}{
				"sender1": map[string]interface{}{
					"sent":   false,
					"errors": []string{"error0"},
				},
			},
		},
		{
			SendingStatus: nil,
		},
	}

	sr := NewSenderRegistry()

	for _, test := range tests {
		sr.setSuccess(test, "sender1")
		if !test.SendingStatus["sender1"].(map[string]interface{})["sent"].(bool) {
			t.Errorf("Expected true, got %v", test.SendingStatus["sender1"].(map[string]interface{})["sent"].(bool))
		}
	}
}
