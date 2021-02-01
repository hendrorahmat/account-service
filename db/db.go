package db

import (
	"fmt"
	"log"

	"github.com/vsouza/go-gin-boilerplate/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	_, errorPing := db.DB()
	// fmt.Println(statusPing)
	fmt.Println(errorPing)
	// err = db.DB().Ping()

	// error ping to database
	if errorPing != nil {
		log.Fatalln(errorPing)
	}

	DB = db
}

func getDatabase() *gorm.DB {
	return DB
}
