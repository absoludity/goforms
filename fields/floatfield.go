package fields

import (
	"errors"
	"strconv"
	"strings"
)

type FloatField struct {
	BaseField
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
func (f FloatField) Clean(value string) (interface{}, ValidationError) {
	cleaned_value, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return nil, errors.New("The value must be a valid float.")
	}
	return cleaned_value, nil
}

func (f FloatField) IsRequired() bool {
	return f.Required
}

// Create and initialise the new fields with the given defaults.
func NewFloatField(defaults Defaults) FloatField {
	field := FloatField{}
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
