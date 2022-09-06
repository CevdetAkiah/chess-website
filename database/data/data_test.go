package data

import (
	"database/sql"
	"fmt"
	"go-projects/chess/service"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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

var (
	mockServ service.DbService
)

func setUp() {
	mux = http.NewServeMux()
	writer = httptest.NewRecorder()

	testDb, err = sql.Open("postgres", "user=cevdet dbname=website password=cevdet sslmode=disable")
	if err != nil {
		err = fmt.Errorf("\nCannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	query, err := ioutil.ReadFile("../testpsql-setup/setup")
	// query,
	if err != nil {
		panic(err)
	}
	if _, err := testDb.Exec(string(query)); err != nil {
		panic(err)
	}

	mockServ = service.DbService{
		Db:             testDb,
		UserService:    testUserAccess{},
		SessionService: testSessionAccess{},
	}

}

// Tests Encrypt function
func TestEncrypt(t *testing.T) {
	testText := "test"

	cryptText := Encrypt(testText)
	if cryptText == testText {
		t.Fail()
		t.Logf("cryptText: %s \t testText: %s", cryptText, testText)
	}

	if bcrypt.CompareHashAndPassword([]byte(cryptText), []byte(testText)) != nil {
		t.Fail()
		t.Logf("testText: %s \t cryptText: %s", testText, cryptText)
	}
}

// Tests CreateUUID function
func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()

	if uuid == "" {
		t.Fail()
	}
}

// Tests AssignCookie function
func TestAssignCookie(t *testing.T) {
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
// func TestDeleteCookie(t *testing.T) {
// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/logout", nil)

// 	testCookie := &http.Cookie{
// 		Name:     "session",
// 		Value:    CreateUUID(),
// 		HttpOnly: true,
// 	}
// 	t.Log("HERE")

// 	http.SetCookie(writer, testCookie)
// 	t.Log("HERE after set cookie")

// 	session := DeleteCookie(http.ResponseWriter(writer), request)
// 	t.Log("HERE after set delete")
// 	fmt.Println(session)

// 	cookie, err := request.Cookie("session")
// 	if err != http.ErrNoCookie {
// 		t.Errorf("Got %s, wanted %d", cookie.Value, http.ErrNoCookie)
// 	}

// }

// Tests AuthSession function
func TestAuthSession(t *testing.T) {
	writer := httptest.NewRecorder()
	form := url.Values{}
	// create form value
	form.Add("password", "12345")

	// add form value to request body
	request, _ := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create testuser and add to testdb
	user := service.User{
		Name:     "testuser",
		Email:    "test@email.com",
		Password: "12345",
	}
	user.Password = Encrypt(user.Password)
	err := mockServ.NewUser(&user)

	// get the user back from testdb
	user, err = mockServ.UserByEmail("test@email.com")

	// if AuthSession can match form pw to testdb pw then the user is given a session and AuthSession is behaving
	err = AuthSession(writer, request, user, mockServ)
	require.NoError(t, err)

	sess, err := SessionById(user.Id)
	require.NoError(t, err)

	require.Equal(t, user.Uuid, sess.Uuid)
}
