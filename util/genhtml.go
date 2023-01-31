package util

import (
	"bytes"
	"fmt"
	"go-projects/chess/service"
	"html/template"
	"net/http"
)

func InitHTML(w http.ResponseWriter, r *http.Request, filename string, DBAccess service.DbService, errMsg string) {
	var buf bytes.Buffer
	// Gather the data for insertion into the templates
	TplData := templateData(r, errMsg)

	// Parse both the html page and layout
	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../templates/%s.page.html", filename)))
	if CheckLogin(r, DBAccess) {
		tpl.ParseFiles("../templates/nav-loggedin.layout.html")
	} else {
		tpl.ParseFiles("../templates/nav-loggedout.layout.html")
	}

	// Write the template to the buffer first
	err := tpl.Execute(&buf, TplData)
	// Handle the error if any
	if err != nil {
		RouteError(w, r, err, "InitHTML", "Initialize template")
	}
	// Write the buffer to the writer
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
	return
}

func CheckLogin(r *http.Request, DBAccess service.DbService) (ok bool) {
	cookie, err := r.Cookie("session")
	if err == nil {
		ok, err = DBAccess.SessionService.CheckSession(cookie.Value)
	} else {
		ok = false
	}
	return
}
