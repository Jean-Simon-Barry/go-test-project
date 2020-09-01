package book

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Repository : repo for books
type Repository struct {
	DB *gorm.DB
}

// ProvideBookRepostiory : provides ProvideProductRepostiory
func ProvideBookRepostiory(DB *gorm.DB) Repository {
	return Repository{DB: DB}
}

// FindBookByTitle : finds book by title
func (b *Repository) FindBookByTitle(bookTitle string) Book {
	var book Book
	b.DB.Where("title = ?", bookTitle).Find(&book)
	return book
}

// FindBookByID : finds book by ID
func (b *Repository) FindBookByID(bookID uint) Book {
	var book Book
	b.DB.First(&book, bookID)
	return book
}

// GetAllBooks : gets all books from db
func (b *Repository) GetAllBooks() []Book {
	var books []Book
	if err := b.DB.Find(&books).Error; err != nil {
		fmt.Println(err)
	}
	return books
}
