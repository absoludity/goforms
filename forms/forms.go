// Package forms enables collecting form fields for validation and
// error collection.
package forms

import (
	"goforms/fields"
)

type FormFields map[string]fields.Field
// FormData just defines the type used in http.Request.Form
type FormData map[string][]string

type Form struct {
	Fields      FormFields
    Data        FormData
	CleanedData map[string]interface{}
	Errors      map[string]string
}


// IsValid verifies the validity of all the field values
// and collects the errors.
func (f *Form) IsValid() bool {
	isValid := true
	cleanedData := map[string]interface{}{}
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
