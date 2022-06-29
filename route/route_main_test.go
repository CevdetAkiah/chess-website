package route

import (
	"go-projects/chess/service"
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
	testServ = service.DbService{
		Db:             testDb,
		UserService:    testUserService,
		SessionService: testSessService,
	}
}

func TestRequest(t *testing.T) {
	mux.HandleFunc("/", Request(&testServ))

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}
