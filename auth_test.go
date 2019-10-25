package auth

import (
	"testing"
)

func TestGetPasswrod(t *testing.T) {
	secret := "asdaasdh"
	t.Log(len(secret))

	code, raiming, err := GetPassword(secret)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("code:%d \n raiming:%d", code, raiming)
	t.Fail()
}
