package entity

import "github.com/google/uuid"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageURL string `json:"imageurl"`
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

func (user *User) ListAll() []User {
	var users []User
	db.Find(&users)

	return users
}
