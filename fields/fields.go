// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

import (
	"errors"
	"fmt"
	"regexp"
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
