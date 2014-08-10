package fields

import (
	"errors"
	"strconv"
)

type BoolField struct {
	BaseField
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
func (f BoolField) Clean(value string) (interface{}, ValidationError) {
	// Ensure value is between max and min length,
	if v, e := strconv.ParseBool(value); e != nil {
		return nil, errors.New("The value is not boolean")
	} else {
		return v, nil
	}
}

// Create and initialise the new fields with the given defaults.
func NewBoolField(defaults Defaults) BoolField {
	field := BoolField{}
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
