// This is an example of using the goforms library to simplify
// validating and cleaning POSTed data in requests.
// $ gorun example.go
// Input data:  map[name:[Harold] age:[24]]
// Valid, cleaned data is:  map[name:Harold age:24]
// Input data:  map[name:[Pheobe] age:[Eighty]]
// Invalid data, errors are:  map[age:The value must be a valid integer.]
// Input data:  map[name:[A] age:[12]]
// Invalid data, errors are:  map[name:The value must have a minimum length of 2
// characters.]
// Input data:  map[name:[Areallylongname] age:[12]]
// Invalid data, errors are:  map[name:The value must have a maximum length of 10
// characters.]

package main

import (
	"fmt"
	"launchpad.net/goforms"
)

func NewRegistrationForm() *forms.Form {
	// Define the fields and create the example form.
	nameField := forms.NewCharField("name")
	nameField.MinLength = 2
	nameField.MaxLength = 10
	ageField := forms.NewIntegerField("age")

	return forms.NewForm(nameField, ageField)
}

func main() {
	egForm := NewRegistrationForm()

	// Create a batch of example form data which would normally be
	// automatically included with the http post requests.
	multiple_posts := []forms.FormData{
		forms.FormData{
			"name": []string{"Harold"},
			"age":  []string{"24"},
		},
		forms.FormData{
			"name": []string{"Pheobe"},
			"age":  []string{"Eighty"},
		},
		forms.FormData{
			"name": []string{"A"},
			"age":  []string{"12"},
		},
		forms.FormData{
			"name": []string{"Areallylongname"},
			"age":  []string{"12"},
		},
	}

	for _, data := range multiple_posts {
		egForm.SetFormData(data)
		fmt.Println("Input data: ", data)

		if egForm.IsValid() {
			fmt.Println("Valid, cleaned data is: ", egForm.CleanedData)
			// At this point the cleaned data is of the correct type
			// (ie. CleanedData['age'] is an integer, rather than string)
			// and can be safely stored or used as needed.
		} else {
			fmt.Println("Invalid data, errors are: ", egForm.Errors)
			// At this point, the map of errors can be passed to a template
			// which renders them in the context of the form.
		}
	}
}
