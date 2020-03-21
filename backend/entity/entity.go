package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Entity interface {
	Create()
	Read()
	Update()
	Delete()
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

		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Company{})
	}
}

type Company struct {
	ID   string
	Name string
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
