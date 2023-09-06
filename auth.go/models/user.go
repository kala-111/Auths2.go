package model

import (
	"gorm.io/gorm"
)

type Auths struct {
	gorm.Model
	Id        int    `gorm:"id"`
	User_name string `gorm:"user_name"`
	Mail_id   string `gorm:"mail_id"`
	Password  string `gorm:"password"`
}

