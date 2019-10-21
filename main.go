package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func main() {
	gogi.Server()
}

func init {
	db, _ = gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True")
}