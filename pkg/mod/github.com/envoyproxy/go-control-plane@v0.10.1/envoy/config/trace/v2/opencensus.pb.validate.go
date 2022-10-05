// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v2/opencensus.proto

package envoy_config_trace_v2

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

// Validate checks the field values on OpenCensusConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OpenCensusConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OpenCensusConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OpenCensusConfigMultiError, or nil if none found.
func (m *OpenCensusConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *OpenCensusConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTraceConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "TraceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "TraceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTraceConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "TraceConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StdoutExporterEnabled

	// no validation rules for StackdriverExporterEnabled

	// no validation rules for StackdriverProjectId

	// no validation rules for StackdriverAddress

	if all {
		switch v := interface{}(m.GetStackdriverGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "StackdriverGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "StackdriverGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStackdriverGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "StackdriverGrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ZipkinExporterEnabled

	// no validation rules for ZipkinUrl

	// no validation rules for OcagentExporterEnabled

	// no validation rules for OcagentAddress

	if all {
		switch v := interface{}(m.GetOcagentGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "OcagentGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenCensusConfigValidationError{
					field:  "OcagentGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOcagentGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "OcagentGrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OpenCensusConfigMultiError(errors)
	}
	return nil
}

// OpenCensusConfigMultiError is an error wrapping multiple validation errors
// returned by OpenCensusConfig.ValidateAll() if the designated constraints
// aren't met.
type OpenCensusConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpenCensusConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpenCensusConfigMultiError) AllErrors() []error { return m }

// OpenCensusConfigValidationError is the validation error returned by
// OpenCensusConfig.Validate if the designated constraints aren't met.
type OpenCensusConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenCensusConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenCensusConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenCensusConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenCensusConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenCensusConfigValidationError) ErrorName() string { return "OpenCensusConfigValidationError" }

// Error satisfies the builtin error interface
func (e OpenCensusConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenCensusConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenCensusConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenCensusConfigValidationError{}
