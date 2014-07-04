// Copyright 2012 The GoForms Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package goforms/forms enables form data validation, cleaning and error collection, similar to Django forms. From forms_test.go:

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

	personForm.Data = urls.Values{
		"name": {"Michael Nelson"},
		"age":  {"37"},
	}

	if personForm.IsValid() {
		doStuffWithCleanedData()
		// personForm.CleanedData contains cleaned data
		// (ie. an int for age in this case):
		// {"name": "Michael Nelson", "age": 37}
		// so that for required fields you can safely do:
		// var age int = personForm.CleanedData["age"].(int)
	} else {
		doStuffWithErrors()
		// If personForm.Data = urls.Values{"age": {"Not a number"}},
		// then personForm.Errors would be {
		//	 "name": "This field is required.",
		//	 "age":  "The value ust be a valid integer."
		// }
	}

You can see more examples in goforms/forms/forms_test.go.

For another form-data processing library, see github.com/gorilla/schema, which fills structs with form data using struct tags.

*/
package forms

import (
	//"github.com/absoludity/goforms/fields"
	"github.com/MartinBrugnara/goforms/fields"
	"net/url"
)

// A collection of fields used on a Form.
type FormFields map[string]fields.Field

// A form brings together a collection of fields, the form data
// to be validated against those fields, as well as the
// generated cleaned data or the collection of errors.
type Form struct {
	Fields      FormFields
	Data        url.Values
	CleanedData map[string]interface{}
	Errors      map[string]string
}

// Contains the cleaned data after a call to IsValid().
type CleanedData map[string]interface{}

// Returns whether the given data validates against the form's
// fields. If successful, CleanedData will be populated,
// otherwise Errors will be populated.
func (f *Form) IsValid() bool {
	isValid := true
	cleanedData := CleanedData{}
	errors := map[string]string{}

	for fieldName, field := range f.Fields {
		dataValues, fieldHasData := f.Data[fieldName]
		dataValue := ""
		switch fieldHasData {
		case true:
			if vlen := len(dataValues); vlen == 1 {
				dataValue = dataValues[0]
				break // do not fallthrough
			} else if vlen > 1 {
				errors[fieldName] = "Multiple value availabele for this field"
				isValid = false
				break // do not fallthrough
			}
			// No value available -> check if is required
			fallthrough
		case false:
			if field.IsRequired() {
				errors[fieldName] = "This field is required."
				isValid = false
			}
			continue
		}
		cleanedValue, err := field.Clean(dataValue)
		if err == nil {
			cleanedData[fieldName] = cleanedValue
		} else {
			errors[fieldName] = err.Error()
			isValid = false
		}
	}

	if isValid {
		f.CleanedData = cleanedData
	} else {
		f.Errors = errors
	}
	return isValid
}
