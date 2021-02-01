package database

import (
	"fmt"
	"strconv"

	"github.com/JoWiel/component-set-generator/config"
	"github.com/jinzhu/gorm"
)

// DB gorm connector
var DB *gorm.DB

// CurrentHostEnv return the current Host Local vs Docker
func CurrentHostEnv() string {
	hostEnv := "DB_DEV_HOST"
	if config.Config("ENV") != "development" {
		hostEnv = "DB_HOST"
	}
	return hostEnv
}

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config(CurrentHostEnv()), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.Product{})
	// fmt.Println("Database Migrated")
}