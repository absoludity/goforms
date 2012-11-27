package fields

import (
	"errors"
	"strconv"
)

type IntegerField struct {
	BaseField
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
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

// Create and initialise the new fields with the given defaults.
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
