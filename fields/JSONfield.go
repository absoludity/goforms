package fields

import (
	"encoding/json"
)

type JSONField struct {
	BaseField
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
func (f JSONField) Clean(value string) (interface{}, ValidationError) {
	var js interface{}
	e := json.Unmarshal([]byte(value), &js)
	return js, e
}

// Create and initialise the new fields with the given defaults.
func NewJSONField(defaults Defaults) CharField {
	// This is a helper for creating fields, so that users of goforms/fields
	// don't need to do things like:
	// "purchaseCount": fields.IntegerField{fields.BaseField{Required: true}},
	// More details here:
	// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/FS_H0SiEioA
	// If that ever changes, we should be able to use simple
	// struct literals.
	field := CharField{}
	for fieldName, value := range defaults {
		switch fieldName {
		case "Required":
			if v, ok := value.(bool); ok {
				field.Required = v
			}
		}
	}
	return field
}
