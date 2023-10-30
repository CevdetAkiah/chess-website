package route

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	request := httptest.NewRequest("GET", "/", nil)
	index(writer, request, testServ)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
	body := writer.Body.String()
	if strings.Contains(body, "/signup") == false {
		t.Errorf("Body does not contain the signup url")
	}
}

func TestErrorPage(t *testing.T) {
	// mux.HandleFunc("/errors", ErrorPage)
	request := httptest.NewRequest("GET", "/errors", nil)
	errorPage(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}

	body := writer.Body.String()
	if strings.Contains(body, "ERROR") == false {
		t.Errorf("Body does not contain the title ERROR")
	}
}

func TestSignup(t *testing.T) {
	request := httptest.NewRequest("GET", "/signup", nil)
	signup(writer, request, testServ)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}

	body := writer.Body.String()
	if strings.Contains(body, "Register") == false {
		t.Errorf("Body does not contain Register")
	}

}

func TestLogin(t *testing.T) {
	request := httptest.NewRequest("GET", "/login", nil)
	login(writer, request, testServ)

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
	body := writer.Body.String()
	if strings.Contains(body, "Login") == false {
		t.Errorf("Body does not contain Login")
	}
}
