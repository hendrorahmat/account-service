package database

import (
	"fmt"

	"github.com/vsouza/go-gin-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (db *gorm.DB, err error) {
	config := config.GetConfig()
	var dbName string = config.GetString("database.connections.default.db")
	var host string = config.GetString("database.connections.default.host")
	var username string = config.GetString("database.connections.default.username")
	var password string = config.GetString("database.connections.default.password")
	var port string = config.GetString("database.connections.default.port")
	var timezone string = config.GetString("timezone")

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=" + timezone
	fmt.Println(config.GetString("database.connections.default.driver"))

	return gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{})
}
