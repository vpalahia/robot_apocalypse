package dbg

import (
	"fmt"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vpalahia/robot_apocalypse/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	Name     = "robo"
	User     = "postgres"
	Password = "Welcome@123"
	Host     = "localhost"
	Port     = "5432"
)

//DialString returns a dial string for Postgres
func DialString() string {

	var dialBuilder strings.Builder
	dialBuilder.WriteString("host=" + Host)
	dialBuilder.WriteString(" port=" + Port)
	dialBuilder.WriteString(" user=" + User)
	dialBuilder.WriteString(" password=" + Password)
	dialBuilder.WriteString(" dbname=" + Name)
	dialBuilder.WriteString(" sslmode=disable")

	return dialBuilder.String()
}

//DbConnection connects to database and migrates the models
func DbConnection() {
	//database connection
	dbConnect, err := gorm.Open(postgres.Open(DialString()), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	//migration of models
	err = dbConnect.AutoMigrate(&models.Survivor{}, &models.LastLocation{}, &models.Inventory{}, &models.Robot{})
	if err != nil {
		fmt.Println("migrate models failed", err.Error())
	}
	DB = dbConnect
}
