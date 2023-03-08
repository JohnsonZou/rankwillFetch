package model

import "gorm.io/gorm"

type Contest struct {
	gorm.Model
	TitleSlug     string `gorm:"type:varchar(30);not null"`
	StartTime     int64  `gorm:"type:bigint"`
	ContestantNum int    `gorm:"type:int"`
}
