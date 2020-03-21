package user

import (
	"github.com/google/uuid"
)

// User type
type User struct {
	ID       string
	Username string
	Email    string
}

// Create a user
func Create(user *User) (*User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// Persist
}

// Read get user by id
func Read(id string) (*User, error) {
	return nil, nil
}

// Update a user
func Update(user *User) (*User, error) {

}

// Delete a user
func Delete(user *User) error {

}
