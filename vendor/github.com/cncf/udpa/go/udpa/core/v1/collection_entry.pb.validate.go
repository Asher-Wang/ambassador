// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: udpa/core/v1/collection_entry.proto

package udpa_core_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _collection_entry_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CollectionEntry with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CollectionEntry) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ResourceSpecifier.(type) {

	case *CollectionEntry_Locator:

		if v, ok := interface{}(m.GetLocator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CollectionEntryValidationError{
					field:  "Locator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CollectionEntry_InlineEntry_:

		if v, ok := interface{}(m.GetInlineEntry()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CollectionEntryValidationError{
					field:  "InlineEntry",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return CollectionEntryValidationError{
			field:  "ResourceSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// CollectionEntryValidationError is the validation error returned by
// CollectionEntry.Validate if the designated constraints aren't met.
type CollectionEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectionEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectionEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectionEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectionEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectionEntryValidationError) ErrorName() string { return "CollectionEntryValidationError" }

// Error satisfies the builtin error interface
func (e CollectionEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectionEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectionEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectionEntryValidationError{}

// Validate checks the field values on CollectionEntry_InlineEntry with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CollectionEntry_InlineEntry) Validate() error {
	if m == nil {
		return nil
	}

	if !_CollectionEntry_InlineEntry_Name_Pattern.MatchString(m.GetName()) {
		return CollectionEntry_InlineEntryValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[0-9a-zA-Z_\\\\-\\\\.~:]+$\"",
		}
	}

	// no validation rules for Version

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CollectionEntry_InlineEntryValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CollectionEntry_InlineEntryValidationError is the validation error returned
// by CollectionEntry_InlineEntry.Validate if the designated constraints
// aren't met.
type CollectionEntry_InlineEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectionEntry_InlineEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectionEntry_InlineEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectionEntry_InlineEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectionEntry_InlineEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectionEntry_InlineEntryValidationError) ErrorName() string {
	return "CollectionEntry_InlineEntryValidationError"
}

// Error satisfies the builtin error interface
func (e CollectionEntry_InlineEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectionEntry_InlineEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectionEntry_InlineEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectionEntry_InlineEntryValidationError{}

var _CollectionEntry_InlineEntry_Name_Pattern = regexp.MustCompile("^[0-9a-zA-Z_\\-\\.~:]+$")
