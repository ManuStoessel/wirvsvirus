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
