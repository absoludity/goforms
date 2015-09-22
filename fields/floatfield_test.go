package fields

import (
    "errors"
    "testing"
)

type FloatTestData []struct {
	in  string
	out interface{}
    err error
}

var cleanFloatCases = FloatTestData{
    {"123.0001", 123.0001, nil},
    {"-260.0000000001", -260.0000000001, nil},
    {"A260", nil, errors.New("The value must be a valid float.")},
}

func TestCleanFloat(t *testing.T) {
    for i, tt := range cleanFloatCases {
        f := NewFloatField(Defaults{})

        cleanedValue, err := f.Clean(tt.in)

        if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
            t.Errorf("%d. Clean() after SetValue(%v) => (%v, %q), expected (%v, %q).", i, tt.in, cleanedValue, err, tt.out, tt.err)
        }
    }
}
