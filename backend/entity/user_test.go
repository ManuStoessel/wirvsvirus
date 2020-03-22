package entity

import (
	"testing"
)

func TestUser(t *testing.T) {
	var testEmail string = "info@alexanderwagner.eu"
	var testEmail2 string = "test@example.de"

	user := &User{Email: testEmail}
	user.Create()

	id := user.ID

	readUser := &User{}
	readUser = readUser.Read(id)

	if readUser.ID != user.ID {
		t.Errorf("Read user should be created user")
	}

	user.Email = testEmail2

	user.Update()

	readUser = readUser.Read(id)

	if readUser.Email != testEmail2 {
		t.Errorf("Mail should be changed")
	}

	user.Delete()

	readUser = readUser.Read(id)

	if readUser.ID != "" {
		t.Errorf("User should not exist anymore")
	}
}
