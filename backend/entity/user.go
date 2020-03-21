package entity

import "github.com/google/uuid"

type User struct {
	ID       string
	Username string
	Email    string
}

func (user *User) Create() {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	db.Create(&user)
}

func (user *User) Update() {
	db.Save(&user)
}

func (user *User) Delete() {
	db.Delete(&user)
}

func (user *User) Read(id string) *User {
	user = &User{}

	db.Where("id = ?", id).First(&user)

	return user
}
