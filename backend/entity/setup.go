package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("Sepro.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&User{})

	db = database

	Role1 := Role{
		Name: "Admin",
	}
	db.Model(&Role{}).Create(&Role1)

	Role2 := Role{
		Name: "Doctor",
	}
	db.Model(&Role{}).Create(&Role2)

	Role3 := Role{
		Name: "Nurse",
	}
	db.Model(&Role{}).Create(&Role3)

	Role4 := Role{
		Name: "User",
	}
	db.Model(&Role{}).Create(&Role4)

}
