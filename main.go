package main

import (
	"github.com/Dhruv501/Zopsmart-mini-Project/datastore"
	"github.com/Dhruv501/Zopsmart-mini-Project/handler"
	"github.com/Dhruv501/Zopsmart-mini-Project/model"
	"gofr.dev/pkg/gofr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root123@tcp(localhost:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Book{})

	app := gofr.New()
	bookStore := datastore.NewBookStore()
	bookHandler := handler.NewBookHandler(bookStore)

	app.GET("/books/{id}", bookHandler.GetByID)
	app.POST("/books", bookHandler.Create)
	app.PUT("/books/{id}", bookHandler.Update)
	app.DELETE("/books/{id}", bookHandler.Delete)
	app.GET("/books", bookHandler.GetAll)

	app.Start()
}
