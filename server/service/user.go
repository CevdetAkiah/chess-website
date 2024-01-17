package service

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Uuid      string
	Name      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	GameID    string `json:"gameID"`
	CreatedAt time.Time
}

// NewUser returns a user object with a hashed password
func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: HashPw(password), CreatedAt: time.Now()}
}

// Authenticate OKs the user for login
func (u *User) Authenticate(password string) (ok bool) {
	return u.CheckPw(password)
}

// DecodeJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func (u *User) DecodeJSON(r *http.Request) error {
	d := json.NewDecoder(r.Body)
	return d.Decode(u)
}

// CheckPW compares the given password against the hashed user password in the database
func (u *User) CheckPw(formPw string) (ok bool) {

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(formPw)) == nil {
		ok = true
	} else {
		ok = false
	}
	return
}

func (u *User) CreateUUID() {
	u.Uuid = uuid.NewV4().String()
}

// HashPw hashes the provided password
func HashPw(text string) (cryptext string) {
	b, _ := bcrypt.GenerateFromPassword([]byte(text), 4)
	cryptext = string(b)

	return
}
