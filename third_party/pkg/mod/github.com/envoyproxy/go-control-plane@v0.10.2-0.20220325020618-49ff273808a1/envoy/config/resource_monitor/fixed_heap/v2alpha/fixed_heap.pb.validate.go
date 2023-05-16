// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto

package v2alpha

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

// Validate checks the field values on FixedHeapConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FixedHeapConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FixedHeapConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FixedHeapConfigMultiError, or nil if none found.
func (m *FixedHeapConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FixedHeapConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMaxHeapSizeBytes() <= 0 {
		err := FixedHeapConfigValidationError{
			field:  "MaxHeapSizeBytes",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FixedHeapConfigMultiError(errors)
	}

	return nil
}

// FixedHeapConfigMultiError is an error wrapping multiple validation errors
// returned by FixedHeapConfig.ValidateAll() if the designated constraints
// aren't met.
type FixedHeapConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FixedHeapConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FixedHeapConfigMultiError) AllErrors() []error { return m }

// FixedHeapConfigValidationError is the validation error returned by
// FixedHeapConfig.Validate if the designated constraints aren't met.
type FixedHeapConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FixedHeapConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FixedHeapConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FixedHeapConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FixedHeapConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FixedHeapConfigValidationError) ErrorName() string { return "FixedHeapConfigValidationError" }

// Error satisfies the builtin error interface
func (e FixedHeapConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFixedHeapConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FixedHeapConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FixedHeapConfigValidationError{}
