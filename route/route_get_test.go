package route

import (
	"testing"
)

// TODO: write a dummy serv struct
func TestIndex(t *testing.T) {
	mux.HandleFunc("/", Request(&testServ))

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}

	// if strings.Contains(body, "/signup") == false {
	// 	t.Errorf("Body does not contain login")
	// }
}

func TestErrorPage(t *testing.T) {
	// mux.HandleFunc("/errors", ErrorPage)
	mux.HandleFunc("/errors", Request(&testServ))

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}

func TestSignup(t *testing.T) {
	mux.HandleFunc("/signup", Request(&testServ))

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}

}

func TestLogin(t *testing.T) {
	mux.HandleFunc("/login", Request(&testServ))

	if writer.Code != 200 {
		t.Errorf("Response code is %d, expected %d", writer.Code, 200)
	}
}
