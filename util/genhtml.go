package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func InitHTML(w http.ResponseWriter, filename string, data ...interface{}) {
	var buf bytes.Buffer

	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../../templates/%s.html", filename)))

	// Write the template to the buffer first
	err := tpl.Execute(&buf, data)
	// Handle the error if any
	ErrHandler(err, "IndexHTML", "Initialize template", time.Now(), w)
	// Write the buffer to the writer
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
	return
}
