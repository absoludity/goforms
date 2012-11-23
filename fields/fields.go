// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

type Field interface {
	SetValue(string)
	Value() string
	Clean() (interface{}, ValidationError)
	CleanedValue() interface{}
}

type ValidationError interface {
	Error() string
}

// BaseField contains items common to all form fields.
type BaseField struct {
	name          string
	value         string
	cleaned_value interface{}
}

// CleanedValue returns the value prepared during clean.
func (f *BaseField) CleanedValue() interface{} {
	return f.cleaned_value
}
func (f *BaseField) SetValue(value string) {
	f.value = value
}
func (f *BaseField) Value() string {
	return f.value
}
