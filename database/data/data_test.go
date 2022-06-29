package data

import (
	"go-projects/chess/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncrypt(t *testing.T) {
	testText := "test"

	cryptText := Encrypt(testText)
	if cryptText == testText {
		t.Fail()
		t.Logf("cryptText: %s \t testText: %s", cryptText, testText)
	}

	if Encrypt(testText) != cryptText {
		t.Fail()
		t.Logf("testText: %s \t cryptText: %s", testText, cryptText)
	}
}

func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()

	if uuid == "" {
		t.Fail()
	}
}

func TestAssignCookie(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	testSess := service.Session{Uuid: "test session"}
	AssignCookie(writer, request, testSess)

	testCookie := http.Cookie{
		Name:     "session",
		Value:    testSess.Uuid,
		HttpOnly: true,
	}

	if writer.Header().Get("Set-Cookie") != testCookie.String() {
		t.Errorf("Writer cookie is %s \t wanted %s", writer.Header().Get("Set-Cookie"), testCookie.String())
	}

	if writer.Code != 302 {
		t.Errorf("Response code is %v", writer.Code)
	}

}

//TODO: fix TestDeleteSession
func TestDeleteSession(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)

	testCookie := http.Cookie{
		Name:     "session",
		Value:    "test",
		HttpOnly: true,
	}
	http.SetCookie(writer, &testCookie)
	DeleteSession(writer, request, nil)

	cookie, err := request.Cookie("session")
	if err != http.ErrNoCookie {
		t.Errorf("Expected %d, got %s", http.ErrNoCookie, cookie.String())
	}
}

// TODO: write a test for AuthSession
