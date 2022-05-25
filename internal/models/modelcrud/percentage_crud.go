package modelcrud

import (
	"fmt"

	"github.com/vpalahia/robot_apocalypse/internal/models"
	dbg "github.com/vpalahia/robot_apocalypse/pkg/dbg"
)

//PercentageCRUD interface for managing Survivors records in the database
type PercentageCRUD struct{}

//GetAll will find percentage for all the infected and non-infected survivors in the database and return them in an array
func (spCRUD *PercentageCRUD) GetAll(infectFlag bool) (float64, error) {
	var survivors []models.Survivor
	var err error
	db := dbg.DB

	// Get all records
	result := db.Debug().Find(&survivors)
	err = result.Error
	if err != nil {
		return 0, fmt.Errorf("failed to find total survivors: %w", err)
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("zero survivors: %w", err)
	}
	totalSurvivors := result.RowsAffected

	//Get all infected or non-infected survivors
	affectedResult := db.Debug().Where("infected = ? ", infectFlag).Find(&survivors)
	err = affectedResult.Error
	if err != nil {
		return 0, fmt.Errorf("failed to find percentage of infected or non-infected survivors: %w", err)
	}
	if affectedResult.RowsAffected == 0 {
		return 0, nil
	}
	affectedSurvivors := affectedResult.RowsAffected

	//calculate percentage
	percentage := spCRUD.calculatepercentage(float64(totalSurvivors), float64(affectedSurvivors))

	return percentage, nil
}

//calculatepercentage will calculate the percentage of survivors
func (spCRUD *PercentageCRUD) calculatepercentage(totalSurvivors, totalAffectedSurvivors float64) float64 {
	percentage := (totalAffectedSurvivors * 100) / totalSurvivors
	return percentage
}
