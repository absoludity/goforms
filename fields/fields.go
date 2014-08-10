// Copyright 2012 The GoForms Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Extended by MartinBrugnara (2014) <martin@martin-dev.eu>
//

/*
Package goforms/fields provides various types of fields for validating
and cleaning individual form values.
*/
package fields

import (
	"errors"
)

// A field is able to clean a string value to the correct type
// or indicate a relevant error, as well as determine whether it is
// required.
type Field interface {
	Clean(string) (interface{}, ValidationError)
	CleanArray([]string) (interface{}, ValidationError)
	IsRequired() bool
	IsArray() bool
}

type ValidationError interface {
	Error() string
}

// The types that are embedded/shared with more specific field types.
type BaseField struct {
	Required bool
	Array    bool
}

// Returns whether data for this field is required.
func (f BaseField) IsRequired() bool {
	return f.Required
}

// Returns true if this field support array.
func (f BaseField) IsArray() bool {
	return f.Array
}

// Returns true if this field support array.
func (f BaseField) CleanArray([]string) (interface{}, ValidationError) {
	return nil, ValidationError(errors.New("Not yet implemnted feature"))
}

// Specify the default attributes of a field.
type Defaults map[string]interface{}
