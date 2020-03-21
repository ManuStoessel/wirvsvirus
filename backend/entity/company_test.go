package entity

import (
	"testing"
)

func TestCompany(t *testing.T) {
	var testEmail string = "info@alexanderwagner.eu"

	user := &User{Email: testEmail}
	user.Create()

	testCompany := &Company{Name: "Test", UserID: user.ID}

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

	testCompany.Name = "Test2"
	testCompany.BusinessNr = "123"

	testCompany.Update()

	readCompany = readCompany.Read(id)

	if readCompany.Name != "Test2" || readCompany.BusinessNr != "123" {
		t.Errorf("Company update failed")
	}

	testCompany.Delete()

	readCompany = readCompany.Read(id)

	if readCompany.ID != "" {
		t.Errorf("Company should not exist anymore")
	}
}
