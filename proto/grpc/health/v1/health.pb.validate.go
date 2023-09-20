// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: grpc/health/v1/health.proto

package healthv1

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

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckRequestMultiError, or nil if none found.
func (m *HealthCheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	if len(errors) > 0 {
		return HealthCheckRequestMultiError(errors)
	}

	return nil
}

// HealthCheckRequestMultiError is an error wrapping multiple validation errors
// returned by HealthCheckRequest.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckRequestMultiError) AllErrors() []error { return m }

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckResponseMultiError, or nil if none found.
func (m *HealthCheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return HealthCheckResponseMultiError(errors)
	}

	return nil
}

// HealthCheckResponseMultiError is an error wrapping multiple validation
// errors returned by HealthCheckResponse.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckResponseMultiError) AllErrors() []error { return m }

// HealthCheckResponseValidationError is the validation error returned by
// HealthCheckResponse.Validate if the designated constraints aren't met.
type HealthCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckResponseValidationError) ErrorName() string {
	return "HealthCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckResponseValidationError{}
