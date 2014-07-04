package fields

import (
	"errors"
	"fmt"
	"regexp"
)

type RegexField struct {
	BaseField
	MatchString string
	Max         int
	Min         int
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
func (f RegexField) Clean(value string) (interface{}, ValidationError) {
	// Copy and paste from charfield.go
	// Ensure value is between max and min length,
	if f.Max != 0 && len(value) > f.Max {
		return nil, errors.New(fmt.Sprint(
			"The value must have a maximum length of ",
			f.Max, " characters."))
	}
	if len(value) < f.Min {
		return nil, errors.New(fmt.Sprint(
			"The value must have a minimum length of ",
			f.Min, " characters."))
	}

	// original code
	matches, err := regexp.MatchString("^"+f.MatchString+"$", value)
	if err != nil {
		return nil, errors.New(
			"The regexp could not be compiled.")
	}
	if !matches {
		return nil, errors.New(fmt.Sprint(
			"The input '", value, "' did not match '",
			f.MatchString, "'."))
	}
	return value, nil
}

// Create and initialise the new fields with the given defaults.
func NewRegexField(defaults Defaults) RegexField {
	field := RegexField{}
	for fieldName, value := range defaults {
		switch fieldName {
		case "Required":
			if v, ok := value.(bool); ok {
				field.Required = v
			}
		case "MatchString":
			if v, ok := value.(string); ok {
				field.MatchString = v
			}
		}
	}
	return field
}
