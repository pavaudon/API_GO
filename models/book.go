package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `gorm:"not null" form:"title" json:"title"`
	Author string `gorm:"not null" form:"author" json:"author"`
}

func (book *Book) GetBooks(db *gorm.DB) (*[]Book, error) {
	var books []Book
	err := db.Model(&Book{}).Find(&books).Error
	return &books, err
}

func (book *Book) GetBookById(db *gorm.DB, id uint) error {
	return db.Model(&Book{}).First(book, id).Error
}

func (book *Book) GetBooksByAuthor(db *gorm.DB, author string) (*[]Book, error) {
	var books []Book
	err := db.Model(&Book{}).Where("author = ?", author).Find(&books).Error
	return &books, err
}

func (book *Book) UpdateOrCreateBook(db *gorm.DB) error {
	return db.Save(book).Error
}

func (book *Book) DeleteBook(db *gorm.DB, id uint) error {
	err := db.Delete(book, id).Error
	return err
}
