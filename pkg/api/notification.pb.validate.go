// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: notification.proto

package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateNotificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNotificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNotificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNotificationRequestMultiError, or nil if none found.
func (m *CreateNotificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNotificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for ConditionId

	// no validation rules for Config

	if len(errors) > 0 {
		return CreateNotificationRequestMultiError(errors)
	}

	return nil
}

// CreateNotificationRequestMultiError is an error wrapping multiple validation
// errors returned by CreateNotificationRequest.ValidateAll() if the
// designated constraints aren't met.
type CreateNotificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNotificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNotificationRequestMultiError) AllErrors() []error { return m }

// CreateNotificationRequestValidationError is the validation error returned by
// CreateNotificationRequest.Validate if the designated constraints aren't met.
type CreateNotificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNotificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNotificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNotificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNotificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNotificationRequestValidationError) ErrorName() string {
	return "CreateNotificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNotificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNotificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNotificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNotificationRequestValidationError{}

// Validate checks the field values on GetNotificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNotificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNotificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNotificationRequestMultiError, or nil if none found.
func (m *GetNotificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNotificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return GetNotificationRequestMultiError(errors)
	}

	return nil
}

// GetNotificationRequestMultiError is an error wrapping multiple validation
// errors returned by GetNotificationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetNotificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNotificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNotificationRequestMultiError) AllErrors() []error { return m }

// GetNotificationRequestValidationError is the validation error returned by
// GetNotificationRequest.Validate if the designated constraints aren't met.
type GetNotificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotificationRequestValidationError) ErrorName() string {
	return "GetNotificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotificationRequestValidationError{}

// Validate checks the field values on GetNotificationsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNotificationsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNotificationsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNotificationsRequestMultiError, or nil if none found.
func (m *GetNotificationsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNotificationsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return GetNotificationsRequestMultiError(errors)
	}

	return nil
}

// GetNotificationsRequestMultiError is an error wrapping multiple validation
// errors returned by GetNotificationsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetNotificationsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNotificationsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNotificationsRequestMultiError) AllErrors() []error { return m }

// GetNotificationsRequestValidationError is the validation error returned by
// GetNotificationsRequest.Validate if the designated constraints aren't met.
type GetNotificationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotificationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotificationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotificationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotificationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotificationsRequestValidationError) ErrorName() string {
	return "GetNotificationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotificationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotificationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotificationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotificationsRequestValidationError{}

// Validate checks the field values on NotificationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotificationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationResponseMultiError, or nil if none found.
func (m *NotificationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ConditionId

	// no validation rules for Config

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NotificationResponseMultiError(errors)
	}

	return nil
}

// NotificationResponseMultiError is an error wrapping multiple validation
// errors returned by NotificationResponse.ValidateAll() if the designated
// constraints aren't met.
type NotificationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationResponseMultiError) AllErrors() []error { return m }

// NotificationResponseValidationError is the validation error returned by
// NotificationResponse.Validate if the designated constraints aren't met.
type NotificationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationResponseValidationError) ErrorName() string {
	return "NotificationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationResponseValidationError{}

// Validate checks the field values on NotificationListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotificationListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationListResponseMultiError, or nil if none found.
func (m *NotificationListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NotificationListResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotificationListResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationListResponseValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return NotificationListResponseMultiError(errors)
	}

	return nil
}

// NotificationListResponseMultiError is an error wrapping multiple validation
// errors returned by NotificationListResponse.ValidateAll() if the designated
// constraints aren't met.
type NotificationListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationListResponseMultiError) AllErrors() []error { return m }

// NotificationListResponseValidationError is the validation error returned by
// NotificationListResponse.Validate if the designated constraints aren't met.
type NotificationListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationListResponseValidationError) ErrorName() string {
	return "NotificationListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationListResponseValidationError{}