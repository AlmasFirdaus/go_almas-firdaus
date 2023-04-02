package database

import (
	"myApp/config"
	"myApp/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var book []models.Book

	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if e := config.DB.Create(&book).Error; e != nil {
		return nil, e
	}

	return book, nil
}

func UpdateBook(book models.Book) (interface{}, error) {
	var bookUpdate models.Book
	config.DB.First(&bookUpdate, book.Id)

	if e := config.DB.Model(&bookUpdate).Updates(models.Book{Judul: book.Judul, Penulis: book.Penulis, Penerbit: book.Penerbit}).Error; e != nil {
		return nil, e
	}
	return bookUpdate, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.Delete(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}
