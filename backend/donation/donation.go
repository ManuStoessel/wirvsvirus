package donation

import (
	"github.com/ManuStoessel/wirvsvirus/backend/user"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Donation type
type Donation struct {
	ID         string
	ReceiverID string
	Receiver   user.User
	Amount     float32
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

		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Donation{})
	}
}

// Create a user
func Create(donation *Donation) {
	if donation.ID == "" {
		donation.ID = uuid.New().String()
	}

	db.Create(&donation)
}

// Read get user by id
func Read(id string) (donation *Donation) {
	donation = &Donation{}

	db.Where("id = ?", id).First(&donation)

	return
}

// ReadByUser get all donations by user
func ReadByUser(userID string) (donations []Donation) {
	db.Where("receiver_id = ?", userID).Find(&donations)

	return
}

// Update a user
func Update(donation *Donation) {
	db.Save(&donation)
}

// Delete a user
func Delete(donation *Donation) {
	db.Delete(&donation)
}
