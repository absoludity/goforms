package fields

import (
	"errors"
	"strconv"
)

type IntegerField struct {
	BaseField
}

// Clean verifies the validity of the given value and prepares the cleaned
// value, returning an error for invalid data.
func (f IntegerField) Clean(value string) (interface{}, ValidationError) {
	cleaned_value, error := strconv.Atoi(value)
	if error != nil {
		return nil, errors.New(
			"The value must be a valid integer.")
	}
	return cleaned_value, nil
}

func (f IntegerField) IsRequired() bool {
	return f.Required
}

func NewIntegerField(defaults Defaults) IntegerField {
	field := IntegerField{}
	for fieldName, value := range defaults {
		switch fieldName {
		case "Required":
			if v, ok := value.(bool); ok {
				field.Required = v
			}
		}
	}
    return field
}
