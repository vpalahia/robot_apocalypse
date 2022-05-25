package models

//Robot store the robo information
type Robot struct {
	Model            string `gorm:"type:text" json:"model"`
	SerialNumber     string `gorm:"type:text" json:"serialNumber"`
	ManufacturedDate string `gorm:"type:text" json:"manufacturedDate"`
	Category         string `gorm:"type:text" json:"category"`
}
