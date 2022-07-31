package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Tel      string
	Gender   string
	Brithday time.Time

	// Facility []Facility `gorm:"foreignKey:UserID"`

	RoleID *uint
	Role   Role
}

type Role struct {
	gorm.Model
	Name string
	User []User `gorm:"foreignKey:RoleID"`
}
