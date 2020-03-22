package api

import "github.com/ManuStoessel/wirvsvirus/backend/entity"

// UserList represents a list of users
type UserList struct {
	Count int           `json:"count"`
	Users []entity.User `json:"users"`
}

// DonationList represents a list of users
type DonationList struct {
	Count     int               `json:"count"`
	Donations []entity.Donation `json:"donations"`
}

// CompanyList represents a list of users
type CompanyList struct {
	Count     int              `json:"count"`
	Companies []entity.Company `json:"companies"`
}

// ImageList represents a list of users
type ImageList struct {
	Count  int            `json:"count"`
	Images []entity.Image `json:"images"`
}
