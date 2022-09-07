package util

import (
	"bytes"
	"fmt"
	"go-projects/chess/service"
	"html/template"
	"net/http"
	"time"
)

func InitHTML(w http.ResponseWriter, r *http.Request, filename string, loggedIn bool, serv service.DbService, errMsg string) {
	var buf bytes.Buffer
	// Gather the data for insertion into the templates
	TplData := templateData(r, loggedIn, errMsg)
	// Parse both the html page and layout
	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../templates/%s.page.html", filename)))
	tpl.ParseFiles("../templates/nav.layout.html")
	// Write the template to the buffer first
	err := tpl.Execute(&buf, TplData)
	// Handle the error if any
	if err != nil {
		ErrHandler("InitHTML", "Initialize template ", time.Now(), w, r)
	}
	// Write the buffer to the writer
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
	return
}
