package models

import (
	"database/sql"
	"time"
)

//Survivor stores the users details, it's infect flag, location and inventories
type Survivor struct {
	//Book keeping
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	//Data
	Name         string         `gorm:"type:text" json:"name"`
	Age          int64          `gorm:"type:bigint" json:"age,omitempty"`
	Gender       string         `gorm:"type:text" json:"gender,omitempty"`
	Infected     bool           `gorm:"type:bool" json:"infected"`   //true if the survivor is infected and false if non-infected
	Reporter     *sql.NullInt64 `gorm:"-" json:"reporter,omitempty"` //this field will detect that survivor should be marked as infected or not
	LastLocation *LastLocation  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"last_location,omitempty"`
	Inventory    *Inventory     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"inventory,omitempty"`
}
