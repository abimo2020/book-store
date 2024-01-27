package database

import (
	"github.com/abimo2020/book-store/config"
	"github.com/abimo2020/book-store/models"
)

func CreateBook(book *models.Books) (err error) {
	if err = config.DB.Save(&book).Error; err != nil {
		return
	}
	return
}
