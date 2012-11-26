// Package forms enables collecting form fields for validation and
// error collection.
package forms

import (
	"goforms/fields"
	"net/url"
)

type FormFields map[string]fields.Field

type Form struct {
	Fields      FormFields
    Data        url.Values
	CleanedData map[string]interface{}
	Errors      map[string]string
}

type CleanedData map[string]interface{}

// IsValid verifies the validity of all the field values
// and collects the errors.
func (f *Form) IsValid() bool {
	isValid := true
	cleanedData := CleanedData{}
	errors := map[string]string{}

	for fieldName, field := range f.Fields {
        dataValues, ok := f.Data[fieldName]
        dataValue := ""
        if ok {
            if len(dataValues) == 0 {
                // Is this possible? Need to add case to table-test for forms when created.
                continue
            } else {
                dataValue = dataValues[0]
            }
        }
        if !ok && field.IsRequired() {
            errors[fieldName] = "This field is required."
            isValid = false
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

func NewForm(formFields FormFields) *Form {
    form := Form{Fields: formFields}
    return &form
}
