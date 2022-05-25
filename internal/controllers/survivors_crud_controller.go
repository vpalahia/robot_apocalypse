package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/robot_apocalypse/internal/models"
	"github.com/vpalahia/robot_apocalypse/internal/models/modelcrud"
)

//GetSurvivors to fetch all the survivors based on infect flag
func GetSurvivors(ctx *gin.Context) {
	infectFlag := ctx.Query(infectFlag)
	flag, err := strconv.ParseBool(infectFlag)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crud := modelcrud.SurvivorCRUD{}
	survivors, err := crud.GetAll(flag)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, survivors)
}

//AddSurvivors to add the survivors
func AddSurvivors(ctx *gin.Context) {
	var survivor []models.Survivor
	if err := ctx.ShouldBindJSON(&survivor); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crud := modelcrud.SurvivorCRUD{}
	survivors, err := crud.Create(survivor)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, survivors)
}

//UpdateSurvivor to update last location and infect flag of a survivor
func UpdateSurvivor(ctx *gin.Context) {
	var survivor models.Survivor
	if err := ctx.ShouldBindJSON(&survivor); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if survivor.Reporter != nil && survivor.Reporter.Valid {
		//Mark survivor as infected when atleast three other survivors report their contamination
		if survivor.Reporter.Int64 >= 3 {
			survivor.Infected = true
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot update survivor to infected as reporter count is less"})
			return
		}
	}

	crud := modelcrud.SurvivorCRUD{}
	err := crud.Update(&survivor)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, survivor)
}
