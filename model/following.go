package model

import "gorm.io/gorm"

type Following struct {
	gorm.Model
	Username   string `gorm:"type:varchar(30);not null"`
	Lcusername string `gorm:"type:varchar(100);not null"`
}
