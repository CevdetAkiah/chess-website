package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// CheckSession checks if the session is active using the given uuid
// func (sa SessionAccess) CheckSession(uuid string) (active bool) {
// 	var err error
// 	err = Db.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE uuid = $1", uuid).Scan(&active)
// 	if err != nil {
// 		active = false
// 		return
// 	}
// 	return
// }

func TestCheckSession(t *testing.T) {
	u := User{
		Name:     "test",
		Email:    "tests@email.com",
		Password: "1234",
	}
	err := mockServ.NewUser(&u)
	require.NoError(t, err)
	// TODO: continue this. None of the tests are working right now, for some reason
}
