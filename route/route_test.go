package route

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	writer = httptest.NewRecorder()
}

func TestIndex(t *testing.T) {
	mux.HandleFunc("/", Index)
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestErrorPage(t *testing.T) {
	mux.HandleFunc("/errors", ErrorPage)
	request, _ := http.NewRequest("GET", "/errors", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestSignup(t *testing.T) {
	mux.HandleFunc("/signup", ErrorPage)
	request, _ := http.NewRequest("GET", "/signup", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestLogin(t *testing.T) {
	mux.HandleFunc("/login", SignupAccount)
	request, _ := http.NewRequest("GET", "/login", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

// TODO: write a test for the SignupAccount handler. Will need to learn about dependency injection.

// TODO: write a test for the Authenticate handler
