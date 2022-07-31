package entity

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name    string
	Email   string
	Tel     string
	Gender  string
	Address string
	// Department string
}
