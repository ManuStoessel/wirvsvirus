package user

import "testing"

var testEmail string = "info@alexanderwagner.eu"
var testEmail2 string = "test@example.de"
var testName string = "Alex"

func Test(t *testing.T) {
	user := &User{Username: testName, Email: testEmail}
	Create(user)

	id := user.ID

	readUser := Read(id)

	if readUser.ID != user.ID {
		t.Errorf("Read user should be created user")
	}

	user.Email = testEmail2

	Update(user)

	readUser = Read(id)

	if readUser.Email != testEmail2 {
		t.Errorf("Mail should be changed")
	}

	Delete(user)

	readUser = Read(id)

	if readUser.ID != "" {
		t.Errorf("User should not exist anymore")
	}
}
