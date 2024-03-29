// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: orderHistory.proto

package pb

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

// define the regex for a UUID once up-front
var _order_history_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Order) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OrderMultiError, or nil if none found.
func (m *Order) ValidateAll() error {
	return m.validate(true)
}

func (m *Order) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for State

	if all {
		switch v := interface{}(m.GetOrderHistoryDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderValidationError{
					field:  "OrderHistoryDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderValidationError{
					field:  "OrderHistoryDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrderHistoryDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderValidationError{
				field:  "OrderHistoryDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderMultiError(errors)
	}

	return nil
}

// OrderMultiError is an error wrapping multiple validation errors returned by
// Order.ValidateAll() if the designated constraints aren't met.
type OrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderMultiError) AllErrors() []error { return m }

// OrderValidationError is the validation error returned by Order.Validate if
// the designated constraints aren't met.
type OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderValidationError) ErrorName() string { return "OrderValidationError" }

// Error satisfies the builtin error interface
func (e OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderValidationError{}

// Validate checks the field values on OrderHistoryDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderHistoryDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderHistoryDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderHistoryDetailMultiError, or nil if none found.
func (m *OrderHistoryDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderHistoryDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Amount

	if len(errors) > 0 {
		return OrderHistoryDetailMultiError(errors)
	}

	return nil
}

// OrderHistoryDetailMultiError is an error wrapping multiple validation errors
// returned by OrderHistoryDetail.ValidateAll() if the designated constraints
// aren't met.
type OrderHistoryDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderHistoryDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderHistoryDetailMultiError) AllErrors() []error { return m }

// OrderHistoryDetailValidationError is the validation error returned by
// OrderHistoryDetail.Validate if the designated constraints aren't met.
type OrderHistoryDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderHistoryDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderHistoryDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderHistoryDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderHistoryDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderHistoryDetailValidationError) ErrorName() string {
	return "OrderHistoryDetailValidationError"
}

// Error satisfies the builtin error interface
func (e OrderHistoryDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderHistoryDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderHistoryDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderHistoryDetailValidationError{}

// Validate checks the field values on CreditLimit with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreditLimit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreditLimit with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreditLimitMultiError, or
// nil if none found.
func (m *CreditLimit) ValidateAll() error {
	return m.validate(true)
}

func (m *CreditLimit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Amount

	if len(errors) > 0 {
		return CreditLimitMultiError(errors)
	}

	return nil
}

// CreditLimitMultiError is an error wrapping multiple validation errors
// returned by CreditLimit.ValidateAll() if the designated constraints aren't met.
type CreditLimitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreditLimitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreditLimitMultiError) AllErrors() []error { return m }

// CreditLimitValidationError is the validation error returned by
// CreditLimit.Validate if the designated constraints aren't met.
type CreditLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreditLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreditLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreditLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreditLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreditLimitValidationError) ErrorName() string { return "CreditLimitValidationError" }

// Error satisfies the builtin error interface
func (e CreditLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreditLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreditLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreditLimitValidationError{}

// Validate checks the field values on GetOrderHistoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderHistoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderHistoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderHistoryRequestMultiError, or nil if none found.
func (m *GetOrderHistoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderHistoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		err = GetOrderHistoryRequestValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOrderHistoryRequestMultiError(errors)
	}

	return nil
}

func (m *GetOrderHistoryRequest) _validateUuid(uuid string) error {
	if matched := _order_history_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetOrderHistoryRequestMultiError is an error wrapping multiple validation
// errors returned by GetOrderHistoryRequest.ValidateAll() if the designated
// constraints aren't met.
type GetOrderHistoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderHistoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderHistoryRequestMultiError) AllErrors() []error { return m }

// GetOrderHistoryRequestValidationError is the validation error returned by
// GetOrderHistoryRequest.Validate if the designated constraints aren't met.
type GetOrderHistoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderHistoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderHistoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderHistoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderHistoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderHistoryRequestValidationError) ErrorName() string {
	return "GetOrderHistoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderHistoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderHistoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderHistoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderHistoryRequestValidationError{}

// Validate checks the field values on GetOrderHistoryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderHistoryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderHistoryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderHistoryReplyMultiError, or nil if none found.
func (m *GetOrderHistoryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderHistoryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CustomerId

	for idx, item := range m.GetOrders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetOrderHistoryReplyValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetOrderHistoryReplyValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetOrderHistoryReplyValidationError{
					field:  fmt.Sprintf("Orders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetCreditLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetOrderHistoryReplyValidationError{
					field:  "CreditLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetOrderHistoryReplyValidationError{
					field:  "CreditLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreditLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetOrderHistoryReplyValidationError{
				field:  "CreditLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetOrderHistoryReplyMultiError(errors)
	}

	return nil
}

// GetOrderHistoryReplyMultiError is an error wrapping multiple validation
// errors returned by GetOrderHistoryReply.ValidateAll() if the designated
// constraints aren't met.
type GetOrderHistoryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderHistoryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderHistoryReplyMultiError) AllErrors() []error { return m }

// GetOrderHistoryReplyValidationError is the validation error returned by
// GetOrderHistoryReply.Validate if the designated constraints aren't met.
type GetOrderHistoryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderHistoryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderHistoryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderHistoryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderHistoryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderHistoryReplyValidationError) ErrorName() string {
	return "GetOrderHistoryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderHistoryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderHistoryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderHistoryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderHistoryReplyValidationError{}
