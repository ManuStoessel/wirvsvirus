package entity

import (
	"github.com/google/uuid"
)

// Donation type
type Donation struct {
	ID         string  `json:"id"`
	ReceiverID string  `json:"receiverid"`
	Receiver   User    `json:"receiver"`
	Amount     float32 `json:"amount"`
}

func (donation *Donation) Create() {
	if donation.ID == "" {
		donation.ID = uuid.New().String()
	}

	db.Create(&donation)
}

func (donation *Donation) Update() {
	db.Save(&donation)
}

func (donation *Donation) Delete() {
	db.Delete(&donation)
}

func (donation *Donation) Read(id string) *Donation {
	donation = &Donation{}

	db.Where("id = ?", id).First(&donation)

	return donation
}

func (donation *Donation) ReadByUser(userID string) []Donation {
	var donations []Donation
	db.Where("receiver_id = ?", userID).Find(&donations)

	return donations
}

func (donation *Donation) ListAll() []Donation {
	var donations []Donation
	db.Find(&donations)

	return donations
}
