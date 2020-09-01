package main

import (
	"github.com/jinzhu/gorm"
	"unity-go-project/book"
)

// Injectors from wire.go:

func initBookAPI(db *gorm.DB) book.Controller {
	repository := book.ProvideBookRepostiory(db)
	service := book.ProvideBookService(repository)
	controller := book.ProvideBookController(service)
	return controller
}
