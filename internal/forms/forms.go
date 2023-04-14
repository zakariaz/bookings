package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a new Form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}), // which is equivalent to errors{}
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	} else {
		return true
	}

}
