package models

import "time"

//Inventory stores the information of survivor's inventory
type Inventory struct {
	//Book keeping
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	//Data
	SurvivorID int64  `gorm:"type:bigint" json:"survivor_id"`
	Items      string `gorm:"type:text" json:"items"`
}
