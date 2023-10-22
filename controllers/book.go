package controllers

import (
	"API/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BookRepo struct {
	Db *gorm.DB
}

type BookCreate struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (repository *BookRepo) FindBooks(c *gin.Context) {
	var bookModel models.Book
	books, err := bookModel.GetBooks(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (repository *BookRepo) FindBooksByAuthor(c *gin.Context) {
	author := c.Param("author")

	var bookModel models.Book
	books, err := bookModel.GetBooksByAuthor(repository.Db, author)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (repository *BookRepo) CreateBook(c *gin.Context) {
	var bookInput BookCreate
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération du JSON")
		return
	}

	newBook := models.Book{
		Title:  bookInput.Title,
		Author: bookInput.Author,
	}
	if repository.Db.Model(&newBook).Where("title = ?", newBook.Title).Updates(&newBook).RowsAffected == 0 && repository.Db.Model(&newBook).Where("author = ?", newBook.Author).Updates(&newBook).RowsAffected == 0 {
		err := newBook.UpdateOrCreateBook(repository.Db)
		if err != nil {
			c.String(http.StatusInternalServerError, "Erreur création du livre")
			return
		}
		c.JSON(http.StatusOK, gin.H{"book": newBook})
	} else {
		c.String(http.StatusInternalServerError, "Book already exist")
	}

}

func (repository *BookRepo) DeleteBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du paramètre")
		return
	}

	var bookFind models.Book
	err = bookFind.GetBookById(repository.Db, uint(bookId))
	if err != nil {
		c.String(http.StatusNotFound, "Le livre n'existe pas")
		return
	}

	err = bookFind.DeleteBook(repository.Db, uint(bookId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur suppression du livre")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
