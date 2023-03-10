// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order.proto

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
var _order_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OrderDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderDetailMultiError, or
// nil if none found.
func (m *OrderDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAmount() != 0 {

		if m.GetAmount() < 1 {
			err := OrderDetailValidationError{
				field:  "Amount",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return OrderDetailMultiError(errors)
	}

	return nil
}

// OrderDetailMultiError is an error wrapping multiple validation errors
// returned by OrderDetail.ValidateAll() if the designated constraints aren't met.
type OrderDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderDetailMultiError) AllErrors() []error { return m }

// OrderDetailValidationError is the validation error returned by
// OrderDetail.Validate if the designated constraints aren't met.
type OrderDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderDetailValidationError) ErrorName() string { return "OrderDetailValidationError" }

// Error satisfies the builtin error interface
func (e OrderDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderDetailValidationError{}

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		err = CreateOrderRequestValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOrderDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOrderRequestValidationError{
					field:  "OrderDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOrderRequestValidationError{
					field:  "OrderDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrderDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrderRequestValidationError{
				field:  "OrderDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

func (m *CreateOrderRequest) _validateUuid(uuid string) error {
	if matched := _order_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderReplyMultiError, or nil if none found.
func (m *CreateOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateOrderReplyMultiError(errors)
	}

	return nil
}

// CreateOrderReplyMultiError is an error wrapping multiple validation errors
// returned by CreateOrderReply.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderReplyMultiError) AllErrors() []error { return m }

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on GetOrderStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderStatusRequestMultiError, or nil if none found.
func (m *GetOrderStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetOrderStatusRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOrderStatusRequestMultiError(errors)
	}

	return nil
}

func (m *GetOrderStatusRequest) _validateUuid(uuid string) error {
	if matched := _order_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetOrderStatusRequestMultiError is an error wrapping multiple validation
// errors returned by GetOrderStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type GetOrderStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderStatusRequestMultiError) AllErrors() []error { return m }

// GetOrderStatusRequestValidationError is the validation error returned by
// GetOrderStatusRequest.Validate if the designated constraints aren't met.
type GetOrderStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderStatusRequestValidationError) ErrorName() string {
	return "GetOrderStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderStatusRequestValidationError{}

// Validate checks the field values on GetOrderStatusReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderStatusReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderStatusReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderStatusReplyMultiError, or nil if none found.
func (m *GetOrderStatusReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderStatusReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for State

	if len(errors) > 0 {
		return GetOrderStatusReplyMultiError(errors)
	}

	return nil
}

// GetOrderStatusReplyMultiError is an error wrapping multiple validation
// errors returned by GetOrderStatusReply.ValidateAll() if the designated
// constraints aren't met.
type GetOrderStatusReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderStatusReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderStatusReplyMultiError) AllErrors() []error { return m }

// GetOrderStatusReplyValidationError is the validation error returned by
// GetOrderStatusReply.Validate if the designated constraints aren't met.
type GetOrderStatusReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderStatusReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderStatusReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderStatusReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderStatusReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderStatusReplyValidationError) ErrorName() string {
	return "GetOrderStatusReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderStatusReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderStatusReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderStatusReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderStatusReplyValidationError{}
