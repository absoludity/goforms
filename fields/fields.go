// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

type Field interface {
	Name() string
	SetValue(string)
	Value() string
	Clean() ValidationError
	CleanedValue() interface{}
}

// Can this just be a Stringer?
type ValidationError interface {
	Error() string
}

// BaseField contains items common to all form fields.
type BaseField struct {
	name          string
	value         string
	cleaned_value interface{}
}

// CleanedValue returns the value prepared during clean.
func (f *BaseField) CleanedValue() interface{} {
	return f.cleaned_value
}
func (f *BaseField) Name() string {
	return f.name
}
func (f *BaseField) SetValue(value string) {
	f.value = value
}
func (f *BaseField) Value() string {
	return f.value
}

type CharField struct {
	BaseField
	MaxLength int
	MinLength int
}

// Clean verifies the validity of the given value and prepares the cleaned
// value, returning an error for invalid data.
func (f *CharField) Clean() (error ValidationError) {
	// Ensure value is between max and min length,
	// Might be worth a Cleanable interface?
	if f.MaxLength != 0 && len(f.value) > f.MaxLength {
		return errors.New(fmt.Sprint(
			"The value must have a maximum length of ",
			f.MaxLength, " characters."))
	}
	if len(f.value) < f.MinLength {
		return errors.New(fmt.Sprint(
			"The value must have a minimum length of ",
			f.MinLength, " characters."))
	}

	f.cleaned_value = f.value
	return nil
}

func NewCharField(name string) *CharField {
	field := CharField{}
	field.name = name
	return &field
}

type IntegerField struct {
	BaseField
}

func NewIntegerField(name string) *IntegerField {
	field := IntegerField{}
	field.name = name
	return &field
}

// Clean verifies the validity of the given value and prepares the cleaned
// value, returning an error for invalid data.
func (f *IntegerField) Clean() (error ValidationError) {
	f.cleaned_value, error = strconv.Atoi(f.value)
	if error != nil {
		return errors.New(
			"The value must be a valid integer.")
	}
	return nil
}

type RegexField struct {
	CharField
	MatchString string
}

func NewRegexField(name string, matchString string) *RegexField {
	field := RegexField{MatchString: matchString}
	field.name = name
	return &field
}

func (f *RegexField) Clean() (error ValidationError) {
	matches, err := regexp.MatchString("^"+f.MatchString+"$", f.value)
	if err != nil {
		return errors.New(
			"The regexp could not be compiled.")
	}
	if !matches {
		return errors.New(fmt.Sprint(
			"The input '", f.value, "' did not match '",
			f.MatchString, "'."))
	}
	f.cleaned_value = f.value
	return nil
}
