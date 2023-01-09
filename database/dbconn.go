package database

import (
	"log"

	"github.com/dchf12/jwt/models"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(sqlite.Open("jwt.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	db.AutoMigrate(&models.User{})
}
