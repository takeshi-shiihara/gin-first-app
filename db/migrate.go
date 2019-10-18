package main

import (
	"models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, _ := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
