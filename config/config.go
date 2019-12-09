package config

import (
	"fmt"

	"sejuta_pragrammer_restful/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	
)

var (
	DB *gorm.DB
)

func init() {
func initDB
InitialMigration()
}

// func init() {
// 	initDB()
// }

// type Config struct {
// 	DB_Username string
// 	DB_Password string
// 	DB_Port     string
// 	DB_Host     string
// 	DB_Name     string
// }

// func InitDB() {

// 	config := Config{
// 		DB_Username: "root",
// 		DB_Password: "",
// 		DB_Port:     "3306",
// 		DB_Host:     "localhost",
// 		DB_Name:     "crud_go",
// 	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Host,
// 		config.DB_Port,
// 		config.DB_Name,
// 	)

	var err error
	DB, err = gorm.Open("sqlite3", "./lib/crud_go.db")
	if err != nil {
		panic(err)
	}
	DB = db
}
