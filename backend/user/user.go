package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User type
type User struct {
	ID       string
	Username string
	Email    string
}

var db *gorm.DB

// Question: The db should be closed some time, or not?
func init() {
	if db == nil {
		// TODO: Should be a one liner
		dbobject, err := gorm.Open("mysql", "root:my-secret-pw@/wirvsvirus?charset=utf8&parseTime=True&loc=Local")
		db = dbobject

		if err != nil {
			panic("failed to connect database")
		}

		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	}
}

// Create a user
func Create(user *User) (*User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	db.Create(&user)

	return user, nil
}

// Read get user by id
func Read(id string) (user *User, err error) {
	user = &User{}

	db.Where("id = ?", id).First(&user)

	return
}

// Update a user
func Update(user *User) (*User, error) {
	return nil, nil
}

// Delete a user
func Delete(user *User) error {
	return nil
}
