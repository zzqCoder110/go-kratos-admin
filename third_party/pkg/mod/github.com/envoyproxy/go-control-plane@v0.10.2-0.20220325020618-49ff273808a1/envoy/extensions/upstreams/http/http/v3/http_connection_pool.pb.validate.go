// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/upstreams/http/http/v3/http_connection_pool.proto

package httpv3

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

// Validate checks the field values on HttpConnectionPoolProto with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HttpConnectionPoolProto) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpConnectionPoolProto with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HttpConnectionPoolProtoMultiError, or nil if none found.
func (m *HttpConnectionPoolProto) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpConnectionPoolProto) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return HttpConnectionPoolProtoMultiError(errors)
	}

	return nil
}

// HttpConnectionPoolProtoMultiError is an error wrapping multiple validation
// errors returned by HttpConnectionPoolProto.ValidateAll() if the designated
// constraints aren't met.
type HttpConnectionPoolProtoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpConnectionPoolProtoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpConnectionPoolProtoMultiError) AllErrors() []error { return m }

// HttpConnectionPoolProtoValidationError is the validation error returned by
// HttpConnectionPoolProto.Validate if the designated constraints aren't met.
type HttpConnectionPoolProtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpConnectionPoolProtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpConnectionPoolProtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpConnectionPoolProtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpConnectionPoolProtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpConnectionPoolProtoValidationError) ErrorName() string {
	return "HttpConnectionPoolProtoValidationError"
}

// Error satisfies the builtin error interface
func (e HttpConnectionPoolProtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpConnectionPoolProto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpConnectionPoolProtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpConnectionPoolProtoValidationError{}
