package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vpalahia/robot_apocalypse/internal/controllers"
	"github.com/vpalahia/robot_apocalypse/pkg/dbg"
)

func main() {
	r := gin.Default()

	//database connection
	dbg.DbConnection()

	//routers
	r.GET("/survivors", controllers.GetSurvivors)
	r.POST("/survivors", controllers.AddSurvivors)
	r.PUT("/survivors", controllers.UpdateSurvivor)
	r.GET("/survivor_percentage", controllers.Percentage)
	r.GET("/robo", controllers.Robot)

	//run server on 8080 port
	r.Run(":8080")

}
