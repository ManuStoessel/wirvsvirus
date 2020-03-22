package entity

import (
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
var dbStringDefault string = "root:my-secret-pw@/wirvsvirus?charset=utf8&parseTime=True&loc=Local"

// Question: The db should be closed some time, or not?
func Initialize(dbString string) {
	if dbString == "" {
		dbString = dbStringDefault
	}

	if db == nil {
		// TODO: Should be a one liner
		//dbobject, err := gorm.Open("mysql", )
		dbobject, err := gorm.Open("mysql", dbString)
		db = dbobject

		if err != nil {
			panic("failed to connect database")
		}

		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Company{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Image{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Donation{})
	}
}
