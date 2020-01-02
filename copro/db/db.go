package db

import (
	// "github.com/elcompadre/copro/config"
	"fmt"
	"github.com/jinzhu/gorm"
)

//Database struct
type Database struct {
	db  *gorm.DB
	err error
}

var postgresInfo string

//InitializeConnection method
func (db *Database) InitializeConnection() {
	fmt.Println("Initialize connection endpoint hit")
	// var configuration config.IConfiguration
	// getConfig := configuration.GetConfig()
	// var config = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d", getConfig.User, getConfig.Password, getConfig.Dbname, getConfig.Host, getConfig.Port)
	// db, err := *gorm.Open("postgres", config)
	// defer db.Close()
}
