// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: watch_url.proto

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

// Validate checks the field values on AddUrlRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddUrlRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddUrlRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddUrlRequestMultiError, or
// nil if none found.
func (m *AddUrlRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddUrlRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	// no validation rules for IsList

	if len(errors) > 0 {
		return AddUrlRequestMultiError(errors)
	}

	return nil
}

// AddUrlRequestMultiError is an error wrapping multiple validation errors
// returned by AddUrlRequest.ValidateAll() if the designated constraints
// aren't met.
type AddUrlRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddUrlRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddUrlRequestMultiError) AllErrors() []error { return m }

// AddUrlRequestValidationError is the validation error returned by
// AddUrlRequest.Validate if the designated constraints aren't met.
type AddUrlRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUrlRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUrlRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUrlRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUrlRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUrlRequestValidationError) ErrorName() string { return "AddUrlRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddUrlRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUrlRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUrlRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUrlRequestValidationError{}

// Validate checks the field values on GetUrlsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUrlsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUrlsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUrlsRequestMultiError,
// or nil if none found.
func (m *GetUrlsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUrlsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return GetUrlsRequestMultiError(errors)
	}

	return nil
}

// GetUrlsRequestMultiError is an error wrapping multiple validation errors
// returned by GetUrlsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUrlsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUrlsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUrlsRequestMultiError) AllErrors() []error { return m }

// GetUrlsRequestValidationError is the validation error returned by
// GetUrlsRequest.Validate if the designated constraints aren't met.
type GetUrlsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUrlsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUrlsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUrlsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUrlsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUrlsRequestValidationError) ErrorName() string { return "GetUrlsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUrlsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUrlsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUrlsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUrlsRequestValidationError{}

// Validate checks the field values on SetStateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetStateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetStateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetStateRequestMultiError, or nil if none found.
func (m *SetStateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetStateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IsDisabled

	if len(errors) > 0 {
		return SetStateRequestMultiError(errors)
	}

	return nil
}

// SetStateRequestMultiError is an error wrapping multiple validation errors
// returned by SetStateRequest.ValidateAll() if the designated constraints
// aren't met.
type SetStateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetStateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetStateRequestMultiError) AllErrors() []error { return m }

// SetStateRequestValidationError is the validation error returned by
// SetStateRequest.Validate if the designated constraints aren't met.
type SetStateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetStateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetStateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetStateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetStateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetStateRequestValidationError) ErrorName() string { return "SetStateRequestValidationError" }

// Error satisfies the builtin error interface
func (e SetStateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetStateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetStateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetStateRequestValidationError{}

// Validate checks the field values on UrlResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UrlResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UrlResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UrlResponseMultiError, or
// nil if none found.
func (m *UrlResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UrlResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	// no validation rules for IsList

	// no validation rules for Id

	// no validation rules for IsDisabled

	// no validation rules for Created

	// no validation rules for Updated

	if len(errors) > 0 {
		return UrlResponseMultiError(errors)
	}

	return nil
}

// UrlResponseMultiError is an error wrapping multiple validation errors
// returned by UrlResponse.ValidateAll() if the designated constraints aren't met.
type UrlResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UrlResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UrlResponseMultiError) AllErrors() []error { return m }

// UrlResponseValidationError is the validation error returned by
// UrlResponse.Validate if the designated constraints aren't met.
type UrlResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UrlResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UrlResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UrlResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UrlResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UrlResponseValidationError) ErrorName() string { return "UrlResponseValidationError" }

// Error satisfies the builtin error interface
func (e UrlResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUrlResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UrlResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UrlResponseValidationError{}

// Validate checks the field values on UrlListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UrlListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UrlListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UrlListResponseMultiError, or nil if none found.
func (m *UrlListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UrlListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUrls() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UrlListResponseValidationError{
						field:  fmt.Sprintf("Urls[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UrlListResponseValidationError{
						field:  fmt.Sprintf("Urls[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UrlListResponseValidationError{
					field:  fmt.Sprintf("Urls[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return UrlListResponseMultiError(errors)
	}

	return nil
}

// UrlListResponseMultiError is an error wrapping multiple validation errors
// returned by UrlListResponse.ValidateAll() if the designated constraints
// aren't met.
type UrlListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UrlListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UrlListResponseMultiError) AllErrors() []error { return m }

// UrlListResponseValidationError is the validation error returned by
// UrlListResponse.Validate if the designated constraints aren't met.
type UrlListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UrlListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UrlListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UrlListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UrlListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UrlListResponseValidationError) ErrorName() string { return "UrlListResponseValidationError" }

// Error satisfies the builtin error interface
func (e UrlListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUrlListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UrlListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UrlListResponseValidationError{}
