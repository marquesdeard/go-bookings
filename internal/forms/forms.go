package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors , otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initialzes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required, Checks input value whether exist or not
func (f *Form) Required(fields ...string /*veriatic function*/) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// MinLength, Checks input value length greater than given condition
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long.", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "This field is not an email.")
	}
}
