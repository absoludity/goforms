package fields

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hooks up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type IntegerFieldTestSuite struct{}

var _ = Suite(&IntegerFieldTestSuite{})

func (s *IntegerFieldTestSuite) TestCleanSuccess(c *C) {
	f := NewIntegerField("num_purchases")
	f.SetValue("12345")

	cleanedValue, error := f.Clean()
	c.Check(error, IsNil)
	c.Check(cleanedValue, Equals, 12345)
	c.Check(f.CleanedValue(), Equals, 12345)
}

func (s *IntegerFieldTestSuite) TestCleanInvalid(c *C) {
	f := NewIntegerField("num_purchases")
	f.SetValue("a12345")

	cleanedValue, err := f.Clean()
	c.Check(err.Error(), Equals, "The value must be a valid integer.")
	c.Check(cleanedValue, IsNil)
}

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
