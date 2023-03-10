// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: customer.proto

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

// Validate checks the field values on CreateCustomerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCustomerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCustomerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCustomerRequestMultiError, or nil if none found.
func (m *CreateCustomerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCustomerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if m.GetAmount() != 0 {

		if m.GetAmount() < 1 {
			err := CreateCustomerRequestValidationError{
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
		return CreateCustomerRequestMultiError(errors)
	}

	return nil
}

// CreateCustomerRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCustomerRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCustomerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCustomerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCustomerRequestMultiError) AllErrors() []error { return m }

// CreateCustomerRequestValidationError is the validation error returned by
// CreateCustomerRequest.Validate if the designated constraints aren't met.
type CreateCustomerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCustomerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCustomerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCustomerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCustomerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCustomerRequestValidationError) ErrorName() string {
	return "CreateCustomerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCustomerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCustomerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCustomerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCustomerRequestValidationError{}

// Validate checks the field values on CreateCustomerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCustomerReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCustomerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCustomerReplyMultiError, or nil if none found.
func (m *CreateCustomerReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCustomerReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CustomerId

	if len(errors) > 0 {
		return CreateCustomerReplyMultiError(errors)
	}

	return nil
}

// CreateCustomerReplyMultiError is an error wrapping multiple validation
// errors returned by CreateCustomerReply.ValidateAll() if the designated
// constraints aren't met.
type CreateCustomerReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCustomerReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCustomerReplyMultiError) AllErrors() []error { return m }

// CreateCustomerReplyValidationError is the validation error returned by
// CreateCustomerReply.Validate if the designated constraints aren't met.
type CreateCustomerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCustomerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCustomerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCustomerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCustomerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCustomerReplyValidationError) ErrorName() string {
	return "CreateCustomerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCustomerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCustomerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCustomerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCustomerReplyValidationError{}
