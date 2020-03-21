package entity

import (
	"testing"
)

func Test(t *testing.T) {
	testCompany := &Company{Name: "Test"}

	testCompany.Create()

	id := testCompany.ID

	if id == "" {
		t.Errorf("Company id should not be empty")
	}

	readCompany := &Company{}
	readCompany = readCompany.Read(id)

	if readCompany.ID != id {
		t.Errorf("Company could not be read from db")
	}
}

var testEmail string = "info@alexanderwagner.eu"
var testEmail2 string = "test@example.de"
var testName string = "Alex"

func TestUser(t *testing.T) {
	user := &User{Username: testName, Email: testEmail}
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
