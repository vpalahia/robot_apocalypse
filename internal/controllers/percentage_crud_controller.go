package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/robot_apocalypse/internal/models/modelcrud"
)

//Percentage is to fetch the percentage of infected and non-infected survivors
func Percentage(ctx *gin.Context) {
	infectFlag := ctx.Query(infectFlag)

	flag, err := strconv.ParseBool(infectFlag)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crud := modelcrud.PercentageCRUD{}
	percentage, err := crud.GetAll(flag)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"percentage": percentage})
}
