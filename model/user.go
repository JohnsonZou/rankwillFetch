package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(30);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
