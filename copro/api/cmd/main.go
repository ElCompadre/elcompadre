package main

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"github.com/elcompadre/copro/config"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/lib/pq"
	"github.com/joho/godotenv"
)

func init(){
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	handleRequests()
}

var db *gorm.DB
var err error

type Resource struct {
    gorm.Model

    Link        string
    Name        string
    Author      string
    Description string
    Tags        pq.StringArray `gorm:"type:varchar(64)[]"`
}

func handleRequests() {
	myRouter := mux.NewRouter()

	configuration := config.Configuration{}
	getConfig := configuration.GetConfig()
	var config = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d", getConfig.User, getConfig.Password, getConfig.Dbname, getConfig.Host, getConfig.Port)
	db, err := gorm.Open("postgres", config)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Resource{})

	myRouter.HandleFunc("/resources", GetResources).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+ strconv.Itoa(getConfig.Port), myRouter))
}

func GetResources(w http.ResponseWriter, r *http.Request) {
    var resources []Resource
    db.Find(&resources)
    json.NewEncoder(w).Encode(&resources)
}