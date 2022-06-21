package route

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
)

func setUp() {
	mux = http.NewServeMux()
	writer = httptest.NewRecorder()
}

func TestIndex(t *testing.T) {
	setUp()
	mux.HandleFunc("/", Index)
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestErrorPage(t *testing.T) {
	setUp()
	mux.HandleFunc("/errors", ErrorPage)
	request, _ := http.NewRequest("GET", "/errors", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestSignup(t *testing.T) {
	setUp()
	mux.HandleFunc("/signup", ErrorPage)
	request, _ := http.NewRequest("GET", "/signup", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}
