package model

import "gorm.io/gorm"

type Contestant struct {
	gorm.Model
	Contestname          string  `gorm:"type:varchar(30);not null"`
	Username             string  `gorm:"type:varchar(100);not null"`
	Rank                 int     `gorm:"type:int;not null"`
	Finish_time          int64   `gorm:"type:bigint"`
	Data_region          string  `gorm:"type:varchar(5)"`
	Attend               bool    `gorm:"type:tinyint"`
	AttendedContestCount int     `gorm:"type:int"`
	Score                int     `gorm:"type:int"`
	Rating               float64 `gorm:"type:double"`
	PredictedRating      float64 `gorm:"type:double"`
}
