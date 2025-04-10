// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: filmclip/v1/filmclip.proto

package v1

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

// Validate checks the field values on UploadImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadImageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadImageRequestMultiError, or nil if none found.
func (m *UploadImageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadImageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UploadImageRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetBase64()) < 1 {
		err := UploadImageRequestValidationError{
			field:  "Base64",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UploadImageRequestMultiError(errors)
	}

	return nil
}

// UploadImageRequestMultiError is an error wrapping multiple validation errors
// returned by UploadImageRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadImageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadImageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadImageRequestMultiError) AllErrors() []error { return m }

// UploadImageRequestValidationError is the validation error returned by
// UploadImageRequest.Validate if the designated constraints aren't met.
type UploadImageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadImageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadImageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadImageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadImageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadImageRequestValidationError) ErrorName() string {
	return "UploadImageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadImageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadImageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadImageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadImageRequestValidationError{}

// Validate checks the field values on UploadImageReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UploadImageReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadImageReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadImageReplyMultiError, or nil if none found.
func (m *UploadImageReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadImageReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return UploadImageReplyMultiError(errors)
	}

	return nil
}

// UploadImageReplyMultiError is an error wrapping multiple validation errors
// returned by UploadImageReply.ValidateAll() if the designated constraints
// aren't met.
type UploadImageReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadImageReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadImageReplyMultiError) AllErrors() []error { return m }

// UploadImageReplyValidationError is the validation error returned by
// UploadImageReply.Validate if the designated constraints aren't met.
type UploadImageReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadImageReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadImageReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadImageReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadImageReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadImageReplyValidationError) ErrorName() string { return "UploadImageReplyValidationError" }

// Error satisfies the builtin error interface
func (e UploadImageReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadImageReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadImageReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadImageReplyValidationError{}

// Validate checks the field values on GenClipScriptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenClipScriptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenClipScriptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenClipScriptRequestMultiError, or nil if none found.
func (m *GenClipScriptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GenClipScriptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetClothingImage()) < 1 {
		err := GenClipScriptRequestValidationError{
			field:  "ClothingImage",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetModelImage()) < 1 {
		err := GenClipScriptRequestValidationError{
			field:  "ModelImage",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Prompt

	if len(errors) > 0 {
		return GenClipScriptRequestMultiError(errors)
	}

	return nil
}

// GenClipScriptRequestMultiError is an error wrapping multiple validation
// errors returned by GenClipScriptRequest.ValidateAll() if the designated
// constraints aren't met.
type GenClipScriptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenClipScriptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenClipScriptRequestMultiError) AllErrors() []error { return m }

// GenClipScriptRequestValidationError is the validation error returned by
// GenClipScriptRequest.Validate if the designated constraints aren't met.
type GenClipScriptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenClipScriptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenClipScriptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenClipScriptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenClipScriptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenClipScriptRequestValidationError) ErrorName() string {
	return "GenClipScriptRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GenClipScriptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenClipScriptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenClipScriptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenClipScriptRequestValidationError{}

// Validate checks the field values on GenClipScriptReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenClipScriptReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenClipScriptReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenClipScriptReplyMultiError, or nil if none found.
func (m *GenClipScriptReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GenClipScriptReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Content

	for idx, item := range m.GetScenes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GenClipScriptReplyValidationError{
						field:  fmt.Sprintf("Scenes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GenClipScriptReplyValidationError{
						field:  fmt.Sprintf("Scenes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GenClipScriptReplyValidationError{
					field:  fmt.Sprintf("Scenes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GenClipScriptReplyMultiError(errors)
	}

	return nil
}

// GenClipScriptReplyMultiError is an error wrapping multiple validation errors
// returned by GenClipScriptReply.ValidateAll() if the designated constraints
// aren't met.
type GenClipScriptReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenClipScriptReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenClipScriptReplyMultiError) AllErrors() []error { return m }

// GenClipScriptReplyValidationError is the validation error returned by
// GenClipScriptReply.Validate if the designated constraints aren't met.
type GenClipScriptReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenClipScriptReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenClipScriptReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenClipScriptReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenClipScriptReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenClipScriptReplyValidationError) ErrorName() string {
	return "GenClipScriptReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GenClipScriptReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenClipScriptReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenClipScriptReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenClipScriptReplyValidationError{}

// Validate checks the field values on SceneScript with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SceneScript) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SceneScript with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SceneScriptMultiError, or
// nil if none found.
func (m *SceneScript) ValidateAll() error {
	return m.validate(true)
}

func (m *SceneScript) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Description

	// no validation rules for Dialogue

	// no validation rules for Actions

	// no validation rules for ShotType

	if len(errors) > 0 {
		return SceneScriptMultiError(errors)
	}

	return nil
}

// SceneScriptMultiError is an error wrapping multiple validation errors
// returned by SceneScript.ValidateAll() if the designated constraints aren't met.
type SceneScriptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SceneScriptMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SceneScriptMultiError) AllErrors() []error { return m }

// SceneScriptValidationError is the validation error returned by
// SceneScript.Validate if the designated constraints aren't met.
type SceneScriptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SceneScriptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SceneScriptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SceneScriptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SceneScriptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SceneScriptValidationError) ErrorName() string { return "SceneScriptValidationError" }

// Error satisfies the builtin error interface
func (e SceneScriptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSceneScript.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SceneScriptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SceneScriptValidationError{}
