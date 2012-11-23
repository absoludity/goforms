package fields

import (
    "errors"
    "testing"
)

type IntegerTestData []struct {
	in  string
	out interface{}
    err error
}

var cleanIntegerCases = IntegerTestData{
    {"123", 123, nil},
    {"-260", -260, nil},
    {"A260", nil, errors.New("The value must be a valid integer.")},
}

func TestCleanInteger(t *testing.T) {
    for i, tt := range cleanIntegerCases {
        // XXX Why even have fieldname - do Django forms have a reason? (ie. as opposed to the key of the form).
        f := NewIntegerField("fieldname")
        f.SetValue(tt.in)

        cleanedValue, err := f.Clean()

        if !ErrorsEquivalent(err, tt.err) ||
            cleanedValue != tt.out ||
            f.CleanedValue() != tt.out {
            t.Errorf("%d. Clean() after SetValue(%v) => (%v, %q), expected (%v, %q).", i, tt.in, cleanedValue, err, tt.out, tt.err)
        }

    }
}
