// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Field interface {
	Name() string
	SetValue(string)
	Value() string
	Clean() (interface{}, ValidationError)
	CleanedValue() interface{}
}

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
func (f *IntegerField) Clean() (interface{}, ValidationError) {
	cleaned_value, error := strconv.Atoi(f.value)
	if error != nil {
		return nil, errors.New(
			"The value must be a valid integer.")
	}
	f.cleaned_value = cleaned_value
	return f.cleaned_value, nil
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

func (f *RegexField) Clean() (interface{}, ValidationError) {
	matches, err := regexp.MatchString("^"+f.MatchString+"$", f.value)
	if err != nil {
		return nil, errors.New(
			"The regexp could not be compiled.")
	}
	if !matches {
		return nil, errors.New(fmt.Sprint(
			"The input '", f.value, "' did not match '",
			f.MatchString, "'."))
	}
	f.cleaned_value = f.value
	return f.cleaned_value, nil
}
