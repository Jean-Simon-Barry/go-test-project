package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"unity-go-project/database"
)

func main() {
	os.Setenv("POSTGRES_USER", "unity_go_project_user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_DB", "unity_go_project")
	os.Setenv("POSTGRES_HOST", "host.docker.internal")
	os.Setenv("POSTGRES_PORT", "5432")
	db := database.Init()
	bookController := initBookAPI(db)
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/books", bookController.FindBooks)
	router.GET("/book/:id", bookController.FindBookByID)
	router.GET("books/title/:title", bookController.FindBookByTitle)
	router.Run(":8080")
}
