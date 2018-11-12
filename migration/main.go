package main

import (
	"github.com/LINK/model"

	_ "github.com/jinzhu/gorm"
)

func main() {

	db := model.GetDBConn()

	// migrate
	db.DropTableIfExists(&model.Company{})
	db.DropTableIfExists(&model.Event{})
	db.DropTableIfExists(&model.Content{})
	//db.DropTableIfExists(&model.EventContents{})
	db.DropTableIfExists(&model.ContentsCategories{})
	db.DropTableIfExists(&model.Category{})
	db.DropTableIfExists(&model.Article{})
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.BrowsedHistory{})

	//db.DropTableIfExists(&model.{})

	db.CreateTable(&model.Company{})
	db.CreateTable(&model.Event{})
	db.CreateTable(&model.Content{})
	db.CreateTable(&model.ContentsCategories{})
	db.CreateTable(&model.Category{})
	db.CreateTable(&model.Article{})
	db.CreateTable(&model.User{})
	db.CreateTable(&model.BrowsedHistory{})

	//db.CreateTable(&model.Company{})
}
