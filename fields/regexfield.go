package fields

import (
    "errors"
    "fmt"
    "regexp"
)

type RegexField struct {
	CharField
	MatchString string
}

func NewRegexField(name string, matchString string) *RegexField {
	field := RegexField{MatchString: matchString}
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
	return f.value, nil
}
