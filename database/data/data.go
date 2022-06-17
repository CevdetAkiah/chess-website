package data

import (
	"crypto/sha1"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Encrypt a password
func Encrypt(text string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	return
}

// CreateUUID to store in a cookie
func CreateUUID() string {
	sID := uuid.NewV4()
	return sID.String()
}
