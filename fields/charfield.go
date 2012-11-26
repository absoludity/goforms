package fields

import (
	"errors"
	"fmt"
)

type CharField struct {
	BaseField
	MaxLength int
	MinLength int
}

// Clean verifies the validity of the given value and prepares the cleaned
// value, returning an error for invalid data.
func (f CharField) Clean(value string) (interface{}, ValidationError) {
	// Ensure value is between max and min length,
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

// This is a helper for creating fields, so that users of goforms/fields
// don't need to do things like:
// "purchaseCount": fields.IntegerField{fields.BaseField{Required: true}},
// More details here:
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/FS_H0SiEioA
// If that ever changes, we should be able to use simple
// struct literals.
func NewCharField(defaults Defaults) CharField {
    field := CharField{}
    for fieldName, value := range defaults {
        switch fieldName {
        case "Required":
            if v, ok := value.(bool); ok {
                field.Required = v
            }
        case "MinLength":
            if v, ok := value.(int); ok {
                field.MinLength = v
            }
        case "MaxLength":
            if v, ok := value.(int); ok {
                field.MaxLength = v
            }
        }
    }
    return field
}
