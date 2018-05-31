package publicbio

import (
	"html/template"
	"io"
	"log"
)

var profileTmpl *template.Template

const templatesDir = "../../templates/"

func init() {
	profileTmpl = template.Must(template.New("profile").ParseFiles(templatesDir + "profile.tmpl"))
}

// renderTemplate retrieves the given template and renders it to the given io.Writer.
// If something goes wrong, the error is logged and returned.
func renderTemplate(w io.Writer, tmpl string, data interface{}) error {
	err := profileTmpl.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		log.Printf("[ERROR] Error rendering %s: %s\n", tmpl, err)
	}

	return err
}
