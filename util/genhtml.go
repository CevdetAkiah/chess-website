package util

import (
	"bytes"
	"fmt"
	"go-projects/chess/service"
	"html/template"
	"net/http"
)

func InitHTML(w http.ResponseWriter, r *http.Request, filename string, loggedIn bool, serv service.DbService, errMsg string) {
	var buf bytes.Buffer
	// Gather the data for insertion into the templates
	TplData := templateData(r, errMsg)

	// Parse both the html page and layout
	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../templates/%s.page.html", filename)))
	if loggedIn {
		tpl.ParseFiles("../templates/nav-loggedin.layout.html")
	} else {
		tpl.ParseFiles("../templates/nav-loggedout.layout.html")
	}

	// Write the template to the buffer first
	err := tpl.Execute(&buf, TplData)
	// Handle the error if any
	if err != nil {
		SendError(err)
		url := fmt.Sprintf("/errors?fname=%s&op=%s", "InitHTML", "Initialize template")
		http.Redirect(w, r, url, 303)
	}
	// Write the buffer to the writer
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
	return
}
