package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/justinas/nosurf"
)

type TemplateData struct {
	CSRFToken string
	ErrMsg    string
	CssSrc    string
}

// hashCSS creates a new css file from the source file and current nano second for the purpose of cache busting.

// TODO: error handlin	g
func hashCSS() string {
	// remove any old css files from hashCSS dir
	files, err := ioutil.ReadDir("../static/hashCSS")
	if len(files) != 0 {
		for _, file := range files {
			os.Remove(fmt.Sprintf("../static/hashCSS/%s", file.Name()))
		}
	}
	if err != nil {
		panic(err)
	}

	// Open source css file
	original, err := os.Open("../static/css/index.css")
	if err != nil {
		panic(err)
	}
	defer original.Close()
	// Make new css file
	name := fmt.Sprintf("../static/hashCSS/%d.index.css", time.Now().Nanosecond())
	new, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer new.Close()
	_, err = io.Copy(new, original)
	if err != nil {
		panic(err)
	}

	return new.Name()
}

// templateData compiles data to be inserted into the html templates
func templateData(r *http.Request, eMsg string) TemplateData {
	// get updated css file name for cache busting purposes
	cssFileName := hashCSS()
	token := nosurf.Token(r)

	return TemplateData{
		CSRFToken: token, // CSRFToken nosurf checks against
		ErrMsg:    eMsg,
		CssSrc:    cssFileName, // caching workaround for the CSS file. // TODO: disable this when in production
	}
}
