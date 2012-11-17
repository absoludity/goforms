// Package forms enables collecting form fields for validation and
// error collection.
package forms

import (
	"goforms/fields"
)

type Form struct {
	Errors      map[string]string
	Fields      map[string]fields.Field
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
		cleanedValue, err := field.Clean()
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

func NewForm(formFields ...fields.Field) *Form {
	form := Form{}
	form.Fields = make(map[string]fields.Field)
	for _, field := range formFields {
		form.Fields[field.Name()] = field
	}
	return &form
}
