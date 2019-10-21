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
	user := web.New()
	gogi.Handle("/user/*", user)

	user.Use(middleware.SubRouter)
	user.Use(SuperSecure)
	user.Get("/index", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)

	gogi.Server()
}

func init {
	db, _ = gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True")
}