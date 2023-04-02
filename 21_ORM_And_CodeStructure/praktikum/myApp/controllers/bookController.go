package controllers

import (
	"myApp/lib/database"
	"myApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	books, e := database.GetBook(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"books":   books,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{
		Judul:    c.FormValue("judul"),
		Penulis:  c.FormValue("penulis"),
		Penerbit: c.FormValue("penerbit"),
	}
	newBook, e := database.CreateBook(book)
	c.Bind(&book)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"books":   newBook,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{
		Id:       id,
		Judul:    c.FormValue("judul"),
		Penulis:  c.FormValue("penulis"),
		Penerbit: c.FormValue("penerbit"),
	}
	UpdateBook, e := database.UpdateBook(book)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data book",
		"books":   UpdateBook,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	deleteBook, e := database.DeleteBook(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data book",
		"books":   deleteBook,
	})

}
