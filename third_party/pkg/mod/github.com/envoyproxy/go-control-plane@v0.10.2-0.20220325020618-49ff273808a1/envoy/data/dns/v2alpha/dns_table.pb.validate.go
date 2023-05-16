// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/dns/v2alpha/dns_table.proto

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

// Validate checks the field values on DnsTable with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DnsTable) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsTable with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DnsTableMultiError, or nil
// if none found.
func (m *DnsTable) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsTable) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ExternalRetryCount

	if len(m.GetVirtualDomains()) < 1 {
		err := DnsTableValidationError{
			field:  "VirtualDomains",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetVirtualDomains() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DnsTableValidationError{
						field:  fmt.Sprintf("VirtualDomains[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DnsTableValidationError{
						field:  fmt.Sprintf("VirtualDomains[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("VirtualDomains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetKnownSuffixes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DnsTableValidationError{
						field:  fmt.Sprintf("KnownSuffixes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DnsTableValidationError{
						field:  fmt.Sprintf("KnownSuffixes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("KnownSuffixes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DnsTableMultiError(errors)
	}

	return nil
}

// DnsTableMultiError is an error wrapping multiple validation errors returned
// by DnsTable.ValidateAll() if the designated constraints aren't met.
type DnsTableMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsTableMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsTableMultiError) AllErrors() []error { return m }

// DnsTableValidationError is the validation error returned by
// DnsTable.Validate if the designated constraints aren't met.
type DnsTableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTableValidationError) ErrorName() string { return "DnsTableValidationError" }

// Error satisfies the builtin error interface
func (e DnsTableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTableValidationError{}

// Validate checks the field values on DnsTable_AddressList with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsTable_AddressList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsTable_AddressList with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsTable_AddressListMultiError, or nil if none found.
func (m *DnsTable_AddressList) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsTable_AddressList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAddress()) < 1 {
		err := DnsTable_AddressListValidationError{
			field:  "Address",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetAddress() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 3 {
			err := DnsTable_AddressListValidationError{
				field:  fmt.Sprintf("Address[%v]", idx),
				reason: "value length must be at least 3 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return DnsTable_AddressListMultiError(errors)
	}

	return nil
}

// DnsTable_AddressListMultiError is an error wrapping multiple validation
// errors returned by DnsTable_AddressList.ValidateAll() if the designated
// constraints aren't met.
type DnsTable_AddressListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsTable_AddressListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsTable_AddressListMultiError) AllErrors() []error { return m }

// DnsTable_AddressListValidationError is the validation error returned by
// DnsTable_AddressList.Validate if the designated constraints aren't met.
type DnsTable_AddressListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_AddressListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_AddressListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_AddressListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_AddressListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_AddressListValidationError) ErrorName() string {
	return "DnsTable_AddressListValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_AddressListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_AddressList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_AddressListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_AddressListValidationError{}

// Validate checks the field values on DnsTable_DnsEndpoint with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsTable_DnsEndpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsTable_DnsEndpoint with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsTable_DnsEndpointMultiError, or nil if none found.
func (m *DnsTable_DnsEndpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsTable_DnsEndpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.EndpointConfig.(type) {

	case *DnsTable_DnsEndpoint_AddressList:

		if all {
			switch v := interface{}(m.GetAddressList()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DnsTable_DnsEndpointValidationError{
						field:  "AddressList",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DnsTable_DnsEndpointValidationError{
						field:  "AddressList",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddressList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsEndpointValidationError{
					field:  "AddressList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := DnsTable_DnsEndpointValidationError{
			field:  "EndpointConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return DnsTable_DnsEndpointMultiError(errors)
	}

	return nil
}

// DnsTable_DnsEndpointMultiError is an error wrapping multiple validation
// errors returned by DnsTable_DnsEndpoint.ValidateAll() if the designated
// constraints aren't met.
type DnsTable_DnsEndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsTable_DnsEndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsTable_DnsEndpointMultiError) AllErrors() []error { return m }

// DnsTable_DnsEndpointValidationError is the validation error returned by
// DnsTable_DnsEndpoint.Validate if the designated constraints aren't met.
type DnsTable_DnsEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsEndpointValidationError) ErrorName() string {
	return "DnsTable_DnsEndpointValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsEndpointValidationError{}

// Validate checks the field values on DnsTable_DnsVirtualDomain with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsTable_DnsVirtualDomain) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsTable_DnsVirtualDomain with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsTable_DnsVirtualDomainMultiError, or nil if none found.
func (m *DnsTable_DnsVirtualDomain) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsTable_DnsVirtualDomain) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 2 {
		err := DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_DnsTable_DnsVirtualDomain_Name_Pattern.MatchString(m.GetName()) {
		err := DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEndpoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsTable_DnsVirtualDomainValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsTable_DnsVirtualDomainValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetAnswerTtl(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DnsTable_DnsVirtualDomainValidationError{
				field:  "AnswerTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := DnsTable_DnsVirtualDomainValidationError{
					field:  "AnswerTtl",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return DnsTable_DnsVirtualDomainMultiError(errors)
	}

	return nil
}

// DnsTable_DnsVirtualDomainMultiError is an error wrapping multiple validation
// errors returned by DnsTable_DnsVirtualDomain.ValidateAll() if the
// designated constraints aren't met.
type DnsTable_DnsVirtualDomainMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsTable_DnsVirtualDomainMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsTable_DnsVirtualDomainMultiError) AllErrors() []error { return m }

// DnsTable_DnsVirtualDomainValidationError is the validation error returned by
// DnsTable_DnsVirtualDomain.Validate if the designated constraints aren't met.
type DnsTable_DnsVirtualDomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsVirtualDomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsVirtualDomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsVirtualDomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsVirtualDomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsVirtualDomainValidationError) ErrorName() string {
	return "DnsTable_DnsVirtualDomainValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsVirtualDomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsVirtualDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsVirtualDomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsVirtualDomainValidationError{}

var _DnsTable_DnsVirtualDomain_Name_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
