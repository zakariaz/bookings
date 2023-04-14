package models

import "github.com/zakariaz/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // we put curly braces because we are declaring interface as a type
	CSRFToken string                 // google CSRF if you are confused
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
