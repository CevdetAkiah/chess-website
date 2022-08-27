package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

type TemplateData struct {
	CSRFToken string
	data      interface{}
}

func InitHTML(w http.ResponseWriter, r *http.Request, filename string, data ...interface{}) {
	var buf bytes.Buffer

	TplData := TemplateData{
		CSRFToken: nosurf.Token(r), // CSRFToken nosurf checks against
		data:      data,
	}

	tpl := template.Must(template.ParseFiles(fmt.Sprintf("../templates/%s.html", filename)))

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
