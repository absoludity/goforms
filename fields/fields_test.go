package fields

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hooks up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type RegexFieldTestSuite struct{}

var _ = Suite(&RegexFieldTestSuite{})

func (s *RegexFieldTestSuite) TestCleanSuccess(c *C) {
	f := NewRegexField("alphabet", "a.c")
	f.SetValue("abc")

	cleanedValue, err := f.Clean()

	c.Check(err, IsNil)
	c.Check(cleanedValue, Equals, "abc")
	c.Check(f.CleanedValue(), Equals, "abc")
}

func (s *RegexFieldTestSuite) TestCleanInvalid(c *C) {
	f := NewRegexField("alphabet", "a.c")
	f.SetValue("abz")

	cleanedValue, err := f.Clean()

	c.Check(err.Error(), Equals, "The input 'abz' did not match 'a.c'.")
	c.Check(f.CleanedValue(), IsNil)
	c.Check(cleanedValue, IsNil)
}
