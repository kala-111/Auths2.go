package config

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {

	var db *gorm.DB
	var err error
	username := "root"
	password := "root"
	host := "localhost"
	database := "jwt"
	dsn := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("database connected successfully...")

	return db, nil

}
