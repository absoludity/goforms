package fields

import (
	"errors"
	"testing"
)

type RegexTestData []struct {
	regex string
	in    string
	max   int
	min   int
	out   interface{}
	err   error
}

var cleanRegexCases = RegexTestData{
	{"a.c", "abc", 0, 0, "abc", nil},
	{"a.c", "abz", 0, 0, nil, errors.New("The input 'abz' did not match 'a.c'.")},
	{"a*c", "abzdsadsadc", 5, 0, nil, errors.New("The value must have a maximum length of 5 characters.")},
	{"a*c", "ac", 5, 3, nil, errors.New("The value must have a minimum length of 3 characters.")},
}

func TestCleanRegex(t *testing.T) {
	for i, tt := range cleanRegexCases {
		f := NewRegexField(Defaults{
			"MatchString": tt.regex,
			"Min":         tt.min,
			"Max":         tt.max,
		})

		cleanedValue, err := f.Clean(tt.in)

		if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
			t.Errorf("%d. Clean() for regex %q with value %q => (%v, %q), expected (%v, %q).", i, tt.regex, tt.in, cleanedValue, err, tt.out, tt.err)
		}
	}
}
