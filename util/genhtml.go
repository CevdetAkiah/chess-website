package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func InitHTML(w http.ResponseWriter, r *http.Request, filename string, data ...interface{}) {
	var buf bytes.Buffer

	// Gather the data for insertion into the templates
	TplData := templateData(r, data)
	// Parse both the html page and layout
	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../templates/%s.page.html", filename)))
	tpl.ParseFiles("../templates/nav.layout.html")

	// Write the template to the buffer first
	err := tpl.Execute(&buf, TplData)
	// Handle the error if any
	if err != nil {
		ErrHandler(err, "InitHTML", "Initialize template ", time.Now(), w)
	}
	// Write the buffer to the writer
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
	return
}
