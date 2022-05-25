package models

import (
	"time"
)

//LastLocation stores the survivor's location
type LastLocation struct {
	//Book keeping
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	//Data
	SurvivorID int64  `gorm:"type:bigint" json:"survivor_id"`
	Latitude   string `gorm:"type:text" json:"latitude"`
	Longitude  string `gorm:"type:text" json:"longitude"`
}
