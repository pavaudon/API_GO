package database

import (
	"API/models"
	"gorm.io/gorm"
)

func InitDatabase(db *gorm.DB) {
	var book1 = models.Book{
		Title:  "1984",
		Author: "Georges Orwell",
	}
	if db.Model(&book1).Where("title = ?", book1.Title).Updates(&book1).RowsAffected == 0 && db.Model(&book1).Where("author = ?", book1.Author).Updates(&book1).RowsAffected == 0 {
		book1.UpdateOrCreateBook(db)
	}
	var book2 = models.Book{
		Title:  "Percy Jackson",
		Author: "Rick Riordan",
	}
	if db.Model(&book2).Where("title = ?", book2.Title).Updates(&book2).RowsAffected == 0 && db.Model(&book2).Where("author = ?", book2.Author).Updates(&book2).RowsAffected == 0 {
		book2.UpdateOrCreateBook(db)
	}
	var book3 = models.Book{
		Title:  "Les Misérables",
		Author: "Victor Hugo",
	}
	if db.Model(&book3).Where("title = ?", book3.Title).Updates(&book3).RowsAffected == 0 && db.Model(&book3).Where("author = ?", book3.Author).Updates(&book3).RowsAffected == 0 {
		book3.UpdateOrCreateBook(db)
	}
}
