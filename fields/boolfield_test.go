package fields

import (
	"errors"
	"testing"
)

type BoolTestData []struct {
	in  string
	out interface{}
	err error
}

var cleanBoolCases = BoolTestData{
	{"0", false, nil},
	{"1", true, nil},
	{"f", false, nil},
	{"t", true, nil},
	{"false", false, nil},
	{"true", true, nil},
	{"False", false, nil},
	{"True", true, nil},
	{"FALSE", false, nil},
	{"TRUE", true, nil},
	{"bla", nil, errors.New("The value is not boolean")},
	{"", nil, errors.New("The value is not boolean")},
}

func TestCleanBool(t *testing.T) {
	for i, tt := range cleanBoolCases {
		f := NewBoolField(Defaults{})
		cleanedValue, err := f.Clean(tt.in)

		if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
			t.Errorf("%d. Clena() with (value)=(%s) => (%q, %q), expected (%q, %q).", i, tt.in, cleanedValue, err, tt.out, tt.err)
		}
	}
}
