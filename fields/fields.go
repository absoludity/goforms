// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

type Field interface {
	Clean(string) (interface{}, ValidationError)
}

type ValidationError interface {
	Error() string
}
