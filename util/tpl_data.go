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
	Data      interface{}
	CssSrc    string
}

// hashCSS creates a new css file from the source file and current nano second for the purpose of cache busting.
func hashCSS() string {
	// remove any old css files from hashCSS
	files, err := ioutil.ReadDir("../static/hashCSS")
	if len(files) != 0 {
		for _, file := range files {
			fmt.Println(file.Name())
			path := "../static/hashCSS/" + file.Name()
			os.Remove(path)
		}
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

func templateData(r *http.Request, data ...interface{}) TemplateData {
	// get updated css file name for cache busting purposes
	cssFileName := hashCSS()
	return TemplateData{
		CSRFToken: nosurf.Token(r), // CSRFToken nosurf checks against
		Data:      data,
		CssSrc:    cssFileName, // caching workaround for the CSS file. // TODO: disable this when in production
	}
}
