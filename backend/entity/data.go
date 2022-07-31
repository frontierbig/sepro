package entity

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model

	Name    string
	DataLv1 string
	DataLv2 string
	DataLv3 string
}
