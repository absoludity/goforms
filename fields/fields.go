// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

type Field interface {
	Clean(string) (interface{}, ValidationError)
    IsRequired() bool
}

type ValidationError interface {
	Error() string
}

type BaseField struct {
	Required bool
}

func (f BaseField) IsRequired() bool {
    return f.Required
}

type Defaults map[string]interface{}
