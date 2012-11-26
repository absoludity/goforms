package fields

import (
    "errors"
    "testing"
)

type RegexTestData []struct {
    regex string
    in string
    out interface{}
    err error
}

var cleanRegexCases = RegexTestData{
    {"a.c", "abc", "abc", nil},
    {"a.c", "abz", nil, errors.New("The input 'abz' did not match 'a.c'.")},
}

func TestCleanRegex(t *testing.T) {
    for i, tt := range cleanRegexCases {
        f := NewRegexField(Defaults{"MatchString": tt.regex})

        cleanedValue, err := f.Clean(tt.in)

        if !ErrorsEquivalent(err, tt.err) || cleanedValue != tt.out {
            t.Errorf("%d. Clean() for regex %q with value %q => (%v, %q), expected (%v, %q).", i, tt.regex, tt.in, cleanedValue, err, tt.out, tt.err)
        }
    }
}
