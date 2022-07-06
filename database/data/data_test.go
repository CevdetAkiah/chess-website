package data

import (
	"fmt"
	"go-projects/chess/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
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

//TODO: fix TestDeleteCookie
func TestDeleteCookie(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/logout", nil)

	testCookie := &http.Cookie{
		Name:     "session",
		Value:    "test",
		HttpOnly: true,
	}
	t.Log("HERE")

	http.SetCookie(writer, testCookie)
	t.Log("HERE after set cookie")

	session := DeleteCookie(writer, request)
	fmt.Println(session)

	t.Log("HERE after set delete")

	cookie, err := request.Cookie("session")
	if err != http.ErrNoCookie {
		t.Errorf("Got %s, wanted %d", cookie.Value, http.ErrNoCookie)
	}

}

// TODO: write a test for AuthSession
