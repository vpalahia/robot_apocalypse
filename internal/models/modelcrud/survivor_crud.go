package modelcrud

import (
	"errors"
	"fmt"
	"log"

	"github.com/vpalahia/robot_apocalypse/internal/models"
	dbg "github.com/vpalahia/robot_apocalypse/pkg/dbg"
	"gorm.io/gorm"
)

//SurvivorCRUD interface for managing Survivor records in the database
type SurvivorCRUD struct{}

//GetAll will find all the survivors records based on infect flag in the database and return them in an array.
//ASSUMPTION: Infected flag is true for infected survivor and false for non-infected
func (sCRUD *SurvivorCRUD) GetAll(infectFlag bool) ([]models.Survivor, error) {
	var survivors []models.Survivor
	db := dbg.DB
	if err := db.Debug().Where("infected = ? ", infectFlag).Find(&survivors).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find survivors: %w", err)
	}
	return survivors, nil
}

//Create will create the survivors records in the database and return them in an array
func (sCRUD *SurvivorCRUD) Create(survivors []models.Survivor) ([]models.Survivor, error) {
	db := dbg.DB
	result := db.Create(&survivors)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create survivors: %w", err)
	}

	if result.RowsAffected == 0 {
		log.Println("zero survivors created")
		return nil, nil
	}
	return survivors, nil
}

//Update will update the survivorS records in the database and return them in an array
func (sCRUD *SurvivorCRUD) Update(survivor *models.Survivor) error {
	db := dbg.DB
	result := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(survivor).Where("id = ?", survivor.ID).Updates(survivor)
	err := result.Error
	if err != nil {
		return fmt.Errorf("failed to update survivor: %w", err)
	}

	if result.RowsAffected == 0 {
		log.Println("zero survivor update")
		return nil
	}
	return nil
}
