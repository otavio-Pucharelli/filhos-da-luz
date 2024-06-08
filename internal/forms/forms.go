package forms

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0

}

// New initializes a new form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields in a form
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		if strings.TrimSpace(x) == "" {
			f.Errors.Add(field, "Este campo não pode estar em branco")
		}
	}
}

// Has checks if a form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// IsEmail checks if a form field is a valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Este não é um email válido")
	}

}
