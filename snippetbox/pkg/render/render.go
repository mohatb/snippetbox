package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) *template.Template {

	tmpl = "ui/html/" + tmpl

	ts, err := template.ParseFiles(tmpl, "ui/html/base.layout.tmpl", "ui/html/footer.partial.tmpl")
	if err != nil {
		log.Println(err.Error())
	}

	return ts

}
