package fields

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hooks up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type CharFieldTestSuite struct{}

var _ = Suite(&CharFieldTestSuite{})

func (s *CharFieldTestSuite) TestCleanSuccess(c *C) {
	f := NewCharField("description")
	f.SetValue("Testing 1, 2, 3")

	cleanedValue, error := f.Clean()
	c.Check(cleanedValue, Equals, "Testing 1, 2, 3")
	c.Check(error, IsNil)
	c.Check(f.CleanedValue(), Equals, "Testing 1, 2, 3")
}

func (s *CharFieldTestSuite) TestMinLength(c *C) {
	f := NewCharField("description")
	f.MinLength = 20
	f.SetValue("Testing 1, 2, 3")

	cleanedValue, err := f.Clean()

	c.Check(err.Error(), Equals,
		"The value must have a minimum length of 20 characters.")
	c.Check(cleanedValue, IsNil)
}

func (s *CharFieldTestSuite) TestMaxLength(c *C) {
	f := NewCharField("description")
	f.MaxLength = 10
	f.SetValue("Testing 1, 2, 3")

	cleanedValue, err := f.Clean()

	c.Check(cleanedValue, IsNil)
	c.Check(err.Error(), Equals,
		"The value must have a maximum length of 10 characters.")
}

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
