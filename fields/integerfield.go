package fields

import (
    "errors"
    "strconv"
)

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

