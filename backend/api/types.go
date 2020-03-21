package api

import (
	"github.com/ManuStoessel/wirvsvirus/backend/user"
)

// User defines the user entity on the api
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	PassHash string `json:"passhash"`
}

// Donation defines the donation entity on the api
type Donation struct {
	ID         string    `json:"id"`
	ReceiverID string    `json:"receiverid"`
	Receiver   user.User `json:"receiver"`
	Amount     float32   `json:"amount"`
}
