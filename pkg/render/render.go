package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zkz/bookings/pkg/config"
	"github.com/zkz/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// parseTemplate parses the teplates, render them and send them to the browser
func RenderTemplate(w http.ResponseWriter, templt string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config when in production mode
		tc = app.TemplateCache

	} else {
		// Create or rebuild the template cache every time when in developpment mode
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[templt]
	if !ok {
		log.Fatalf("the %s is not available \n", templt)
	}

	// we could write the templt directly to w but we shose to use a buf ot finetune
	// the error handling and knwoing exacly what happend; is it the template that is
	// not available or we cant write to the ResponseWriter?
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	err := t.Execute(buf, td) // overriding the `err` variable instead of redefining it
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w) // overriding the `err` variable instead of redefining it
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		// the line below parses the page and store it in a template
		ts, err := template.New(name).ParseFiles(page) // ts stands for template set
		if err != nil {
			return myCache, err
		}

		// get all the layouts in the ./templates dir
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			// parse the template defininitions and associate them with ts
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			// now ts has everithing it needs to render the page and the layout!
			myCache[name] = ts
		}
	}
	return myCache, nil
}
