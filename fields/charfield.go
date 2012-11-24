package fields

import (
    "errors"
    "fmt"
)

type CharField struct {
	MaxLength int
	MinLength int
}

// Clean verifies the validity of the given value and prepares the cleaned
// value, returning an error for invalid data.
func (f CharField) Clean(value string) (interface{}, ValidationError) {
	// Ensure value is between max and min length,
	// Might be worth a Cleanable interface?
	if f.MaxLength != 0 && len(value) > f.MaxLength {
		return nil, errors.New(fmt.Sprint(
			"The value must have a maximum length of ",
			f.MaxLength, " characters."))
	}
	if len(value) < f.MinLength {
		return nil, errors.New(fmt.Sprint(
			"The value must have a minimum length of ",
			f.MinLength, " characters."))
	}

	return value, nil
}
