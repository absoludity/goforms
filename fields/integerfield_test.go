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
        f := NewIntegerField(Defaults{})

        cleanedValue, err := f.Clean(tt.in)

        if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
            t.Errorf("%d. Clean() after SetValue(%v) => (%v, %q), expected (%v, %q).", i, tt.in, cleanedValue, err, tt.out, tt.err)
        }

    }
}
