package main

import (
	"API/config"
	"API/controllers"
	"API/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db := config.Connecttodb()
	config.CreateModel(db)
	database.InitDatabase(db)

	bookRepo := controllers.BookRepo{
		Db: db,
	}

	r := gin.Default()

	r.GET("/books", bookRepo.FindBooks)
	r.GET("/books/author/:author", bookRepo.FindBooksByAuthor)
	r.POST("/book", bookRepo.CreateBook)
	r.DELETE("/book/:id", bookRepo.DeleteBook)

	r.Run("localhost:8080")
}
