package forms

import (
	"github.com/absoludity/goforms/fields"
	"net/url"
	"testing"
)

func MakeForm(data string) Form {
	formFields := FormFields{
		"name": fields.NewCharField(fields.Defaults{
			"Required": true,
		}),
		"age": fields.NewIntegerField(fields.Defaults{
			"Required": false,
		}),
		"about": fields.NewCharField(fields.Defaults{
			"Max": 10,
		}),
	}
	personForm := Form{Fields: formFields}
	personForm.Data, _ = url.ParseQuery(data)

	return personForm
}

type TestErrorData map[string]string

type FormTestData []struct {
	in  string
	out CleanedData
	err TestErrorData
}

var FormTestCases = FormTestData{
	// Values are cleaned (name, age) and extra data is ignored (ignore).
	{
		"name=Alpha+Beta&age=24&ignore=ignore+me",
		CleanedData{
			"name": "Alpha Beta",
			"age":  24,
		},
		nil,
	},
	// Invalid data results in collected errors (age, about).
	{

		"name=Alpha+Beta&age=This+is+not+a+number&about=This+is+too+long.&ignore=ignore+me",
		nil,
		TestErrorData{
			"about": "The value must have a maximum length of 10 characters.",
			"age":   "The value must be a valid integer.",
		},
	},
	// A lack of data for a required field is an error (name).
	{
		"about=I+like+Go",
		nil,
		TestErrorData{
			"name": "This field is required.",
		},
	},
	// Empty data for a field does not error (age). [Required False]
	// (Not sure if this is possible, but test anyway.)
	{
		"name=Alpha+Beta&age=",
		CleanedData{
			"name": "Alpha Beta",
		},
		nil,
	},
	// Test empty data on required fields.
	{
		"name=&age=24&ignore=ignore+me",
		nil,
		TestErrorData{
			"name": "This field is required.",
		},
	},
	// Test error on multiple values.
	{
		"name=Ciccio&age=24&ignore=ignore+me&name=Barocco",
		nil,
		TestErrorData{
			"name": "Too many values for this field.",
		},
	},
}

func TestIsValid(t *testing.T) {
	for i, tt := range FormTestCases {
		myForm := MakeForm(tt.in)

		CheckFormValidity(t, i, &myForm)

		CheckFormOutput(t, i, &myForm)
	}
}

func CheckFormValidity(t *testing.T, testCaseIndex int, f *Form) {
	isValid := f.IsValid()

	if isValid && f.CleanedData == nil {
		t.Errorf("%d. isValid() returned true and form has no cleaned data.", testCaseIndex)
	}
	if !isValid && f.Errors == nil {
		t.Errorf("%d. isValid() returned False and form has no errors.", testCaseIndex)
	}
}

func CheckFormOutput(t *testing.T, testCaseIndex int, f *Form) {
	tt := FormTestCases[testCaseIndex]
	// XXX see reflect.DeepEqual
	if len(tt.out) != len(f.CleanedData) {
		t.Errorf("%d. Expected %d entries in CleanedData, got %d.", testCaseIndex, len(tt.out), len(f.CleanedData))
	}
	for key, expected := range tt.out {
		actual, ok := f.CleanedData[key]
		switch {
		case !ok:
			t.Errorf("%d. Key %q not present in CleanedData.", testCaseIndex, key)
		case actual != expected:
			t.Errorf("%d. %q=>%v found in CleanedData. Expected %q=>%v.", testCaseIndex, key, actual, key, expected)
		}
	}
	// XXX see reflect.DeepEqual
	if len(tt.err) != len(f.Errors) {
		t.Errorf("%d. Expected %d entries in Errors, got %d. Errors=>%v.", testCaseIndex, len(tt.err), len(f.Errors), f.Errors)
	}
	for key, expected := range tt.err {
		actual, ok := f.Errors[key]
		switch {
		case !ok:
			t.Errorf("%d. Error with key %q not present in form Errors.", testCaseIndex, key)
		case actual != expected:
			t.Errorf("%d. %q=>%q found in Errors. Expected %q=>%v.", testCaseIndex, key, actual, key, expected)
		}
	}
}
