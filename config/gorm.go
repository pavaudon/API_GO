package config

import (
	"API/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connecttodb() *gorm.DB {
	databaseConfig := "root:root@tcp(192.168.1.219:3306)/bookdb?parseTime=true"
	db, err := gorm.Open(mysql.Open(databaseConfig))

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}
	return db
}

func CreateModel(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}
