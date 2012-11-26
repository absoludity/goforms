package fields

import (
	"errors"
	"fmt"
	"regexp"
)

type RegexField struct {
	BaseField
	MatchString string
}

func (f RegexField) Clean(value string) (interface{}, ValidationError) {
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
