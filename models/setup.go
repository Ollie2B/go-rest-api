package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	db_name := os.Getenv("MYSQL_DATABASE")
	db_user := "root"
	db_pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	db_address := os.Getenv("MYSQL_ADDRESS")
	db_protocol := "tcp"

	cfg := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_protocol, db_address, db_name)
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Song{})

	DB = db
}
