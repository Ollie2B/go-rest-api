package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	
	db_name := "recordings"
	db_user := os.Getenv("DBUSER")
	db_pass := os.Getenv("DBPASS")
	db_address := "127.0.0.1:3306"
	db_protocol  := "tcp"

	cfg := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_protocol, db_address, db_name)
  db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Song{})

	DB = db
}