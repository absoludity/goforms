package fields

import (
	"errors"
	"testing"
)

type CharTestData []struct {
	in  string
	min int
	max int
	out interface{}
    err error
}

var cleanTestCases = CharTestData{
	{"Testing 1, 2, 3", 0, 0, "Testing 1, 2, 3", nil},
	{"1234", 1, 5, "1234", nil},
	{"1234", 4, 4, "1234", nil},
	{"1234", 6, 0, nil, errors.New("The value must have a minimum length of 6 characters.")},
	{"1234", 0, 3, nil, errors.New("The value must have a maximum length of 3 characters.")},
}

func ErrorsEquivalent(e1, e2 error) bool {
    if e1 == nil && e2 == nil {
        return true
    }

    if e1 == nil || e2 == nil {
        return false
    }

    if e1.Error() == e2.Error() {
        return true
    }

    return false
}

func TestClean(t *testing.T) {
	for i, tt := range cleanTestCases {
        f := CharField{MinLength: tt.min, MaxLength: tt.max}
		f.SetValue(tt.in)

		cleanedValue, err := f.Clean()

        if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
            t.Errorf("%d. Clean() with (value, min, max)=(%q, %v, %v) => (%q, %q), expected (%q, %q).", i, tt.in, tt.min, tt.max, cleanedValue, err, tt.out, tt.err)
        }
	}
}
