package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kala-111/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateTable() {
	query := `createtable  login(
		iid int auto increment,
		user_name varchar(255) not null,
		mail_id varchar(255) not null,
		password varchar(255)not null,
				);`
	err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("login table created...")
}

var router *gin.Engine

func main() {
	config.DatabaseConnection()
	//var db *gorm.DB
	//db, _ = DatabaseConnection()
	fmt.Println("hello")
	CreateTable()
	//var err error
	router := gin.New()
	config.DatabaseConnection()
	router.Run(":8080")

}

/*var authinfo Auths
	authinfo.Id = 1
	authinfo.User_name = "chandu"
	authinfo.Mail_id = "chandrakalavathi@gmail.com"
	authinfo.Password = "Kala@123"

	err = db.Create(&authinfo).Error
	if err != nil {
		fmt.Println("error while inserting data : ", err)
		return
	}

	fmt.Println("data inserted successfully...")
}*/
