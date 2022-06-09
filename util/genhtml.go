package util

import (
	"fmt"
	"html/template"
	"net/http"
)

func InitHTML(w http.ResponseWriter, filename string, data ...interface{}) error {
	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../../templates/%s.html", filename)))
	err := tpl.Execute(w, data)
	return err
}
