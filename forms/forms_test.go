package forms

import (
	"goforms/fields"
	"net/url"
	"testing"
)

func MakeForm(data url.Values) Form {
	egForm := Form{
		Fields: FormFields{
			"description":   fields.CharField{MaxLength: 10},
            "purchaseCount": fields.IntegerField{Required: true},
			"otherField":    fields.CharField{},
		},
		Data: data,
	}

	return egForm
}

type TestErrorData map[string]string

type FormTestData []struct {
	in  url.Values
	out CleanedData
	err TestErrorData
}

var FormTestCases = FormTestData{
	{
		url.Values{
			"description":   {"short desc"},
			"purchaseCount": {"24"},
			"ignore":        {"ignore me"},
		},
		CleanedData{
			"description":   "short desc",
			"purchaseCount": 24,
			"otherField":    "",
		},
		nil,
	},
	{
		url.Values{
			"description":   {"This is too long."},
			"purchaseCount": {"abc123"},
			"ignore":        {"ignore me"},
		},
		nil,
		TestErrorData{
			"description":   "The value must have a maximum length of 10 characters.",
			"purchaseCount": "The value must be a valid integer.",
		},
	},
	{
		url.Values{
			"description":   {"short desc"},
			"ignore":        {"ignore me"},
		},
		nil,
		TestErrorData{
			"purchaseCount": "This field is required.",
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


func TestRegexRequired(t *testing.T) {
	egForm := Form{
		Fields: FormFields{
			"description":   fields.CharField{},
            "requiredField": fields.RegexField{Required: true},
		},
		Data: url.Values{
            "description": {"whohoo"},
        },
	}

    isValid := egForm.IsValid()

    if isValid {
        t.Errorf("Form with missing required field should not be valid.")
    }
    if len(egForm.Errors) != 1 {
        t.Errorf("Expected 1 validation error, got %d.", len(egForm.Errors))
    }
    err, ok := egForm.Errors["requiredField"]
    expectedError := "This field is required."
    switch {
    case !ok:
        t.Errorf("Required field was not required.")
    case err != expectedError:
        t.Errorf("Expected %q, got %q.", expectedError, err)
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
