package donation

import (
	"github.com/ManuStoessel/wirvsvirus/backend/v1/user"
	"github.com/google/uuid"
)

// Donation type
type Donation struct {
	ID       string
	Receiver *user.User
	Amount   float32
}

// Create a user
func Create(donation *Donation) (*Donation, error) {
	if donation.ID == "" {
		donation.ID = uuid.New().String()
	}

	// Persist
}

// Read get user by id
func Read(id string) (*Donation, error) {
	return nil, nil
}

// Update a user
func Update(donation *Donation) (*Donation, error) {

}

// Delete a user
func Delete(donation *Donation) error {

}
