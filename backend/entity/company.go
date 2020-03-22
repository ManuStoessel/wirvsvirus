package entity

import (
	"github.com/google/uuid"
)

type Company struct {
	ID               string
	UserID           string
	User             User
	BusinessNr       string
	Name             string
	Town             string
	Street           string
	Business         string
	Description      string
	ShortDescription string
	PaypalButtonId   string
}

func (company *Company) Create() {
	if company.ID == "" {
		company.ID = uuid.New().String()
	}

	db.Create(&company)
}

func (company *Company) Update() {
	db.Save(&company)
}

func (company *Company) Delete() {
	db.Delete(&company)
}

func (company *Company) Read(id string) *Company {
	company = &Company{}

	db.Where("id = ?", id).First(&company)

	return company
}
