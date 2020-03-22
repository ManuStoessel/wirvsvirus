package entity

import (
	"github.com/google/uuid"
)

type Company struct {
	ID               string `json:"id"`
	UserID           string `json:"userid"`
	User             User   `json:"user"`
	BusinessNr       string `json:"businessnr"`
	Name             string `json:"name"`
	Town             string `json:"town"`
	Street           string `json:"street"`
	Business         string `json:"business"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortdescription"`
	PaypalButtonId   string `json:"paypalbuttonid"`
}

func (company *Company) ListAll() []Company {
	var companies []Company
	db.Find(&companies)

	return companies
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
