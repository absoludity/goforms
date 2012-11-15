// Package forms enables collecting form fields for validation and
// error collection.
package forms

import (
    "goforms/forms/fields"
)

type Form struct {
	Errors      map[string]string
	Fields      map[string]Field
	CleanedData map[string]interface{}
}

// FormData just defines the type used in http.Request.Form
type FormData map[string][]string

// IsValid verifies the validity of all the field values
// and collects the errors.
func (f *Form) IsValid() bool {
	isValid := true
	cleanedData := map[string]interface{}{}
	errors := map[string]string{}

	for fieldName, field := range f.Fields {
		err := field.Clean()
		if err == nil {
			cleanedData[fieldName] = field.CleanedValue()
		} else {
			errors[fieldName] = err.String()
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

func (f *Form) SetFormData(data FormData) {
	for fieldName, values := range data {
		field := f.Fields[fieldName]
		if field != nil {
			// For the moment, just handle single-valued
			// params.
			field.SetValue(values[0])
		}
	}
}

func NewForm(fields ...Field) *Form {
	form := Form{}
	form.Fields = make(map[string]Field)
	for i := 0; i < len(fields); i++ {
		form.Fields[fields[i].Name()] = fields[i]
	}
	return &form
}
