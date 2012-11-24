package forms

import (
	"goforms/fields"
	. "launchpad.net/gocheck"
	"testing"
)

// Hooks up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type FormTestSuite struct {
	egForm *Form
}

var _ = Suite(&FormTestSuite{})

// MakeForm is a helper method to make a form with optional form data.
func (s *FormTestSuite) MakeForm(data FormData) *Form {
	egForm := NewForm(FormFields{
        "description":    fields.CharField{MaxLength: 10},
		"purchase_count": fields.IntegerField{},
		"notused":        fields.CharField{},
	})

	if data != nil {
        egForm.Data = data
	}
	return egForm
}

func (s *FormTestSuite) TestIsValidTrue(c *C) {
	var formData = FormData{
		"description":    []string{"short desc"},
		"purchase_count": []string{"24"},
		"ignored":        []string{"ignore me"},
	}
	myForm := s.MakeForm(formData)

	c.Check(myForm.IsValid(), Equals, true)
	c.Check(myForm.CleanedData, DeepEquals, map[string]interface{}{
		"description":    "short desc",
		"purchase_count": 24,
		"notused":        "",
	})
}

func (s *FormTestSuite) TestIsValidFalse(c *C) {
	var formData = FormData{
		"description":    []string{"this is too long"},
		"purchase_count": []string{"2a4"},
	}
	myForm := s.MakeForm(formData)

	c.Check(myForm.IsValid(), Equals, false)
	c.Check(myForm.Errors, DeepEquals, map[string]string{
		"purchase_count": "The value must be a valid integer.",
		"description": "The value must have a maximum length of " +
			"10 characters.",
	})
}
