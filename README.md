# GoForms - form data validation, cleaning and error reporting

The goforms library is a proof-of-concept for a data validation, cleaning and
error collecting using [Go](http://golang.org), in a similar style to Django's django.forms
library. It enables thin handlers like this:

```go
func my_post_handler(w http.ResponseWriter, r *http.Request){
    myForm := MyCustomForm()
    myForm.Data = r.Form

    if myForm.IsValid(){
        // Do something with myForm.CleanedData, which is the request's form
        // data after being cleaned to the correct types etc.
    } else {
        // Re-render a template, passing in egForm.Errors which is a map
        // of errors for each field on the form.
    }
```

## Installation and tests
To install goforms into your current Go workspace:
```
$ go get github.com/absoludity/goforms/forms
```

You can then run the tests with
```
$ go test github.com/absoludity/goforms/fields github.com/absoludity/goforms/forms
```

## Example
### Define your form
```go
	formFields := FormFields{
		"name": fields.NewCharField(fields.Defaults{
			"Required": true,
		}),
		"age": fields.NewIntegerField(fields.Defaults{
			"Required": false,
		}),
		"about": fields.NewCharField(fields.Defaults{
			"Max": 10,
		}),
	}
	personForm := Form{Fields: formFields}
```

### Use the form in your handler
Using the [http.Request](http://golang.org/pkg/net/http/#Request) objects Form (r.Form):
```go
	personForm.Data = r.Form

	if personForm.IsValid() {
		doStuffWithCleanedData()
		// personForm.CleanedData contains cleaned data
		// (ie. an int for age in this case):
		// {"name": "Michael Nelson", "age": 37}
		// so that for required fields you can safely do:
		// var age int = personForm.CleanedData["age"].(int)
	} else {
		doStuffWithErrors()
		// If personForm.Data = urls.Values{"age": {"Not a number"}},
		// then personForm.Errors would be {
		//	 "name": "This field is required.",
		//	 "age":  "The value ust be a valid integer."
		// }
	}
```

## Notes
 * Originally discussed on [this golang-nuts post](http://goo.gl/pFh6I).

## TODO
 * Remove Form.Data, and instead provide data to Form.IsValid()
 * Perhaps make Form.CleanedData and Form.Errors private, providing getters.
 * Add defaults to fields so that a value is always included, or even remove Required, assuming that a field is required if there's no default?.
 * Enable custom error messages.
 * Update field tests to use nicer in/out for test tables like form_tests.

## License
Copyright 2012 The GoForms Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
