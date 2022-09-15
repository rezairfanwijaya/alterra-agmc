package controllers

import (
	"altera/Day2/helper"
	"altera/Day2/lib/database"
	"altera/Day2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllBook(e echo.Context) error {
	// panggil database
	books := database.GetAllBook()

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully get all books",
		"data":    books,
	})

}

func GetBookById(e echo.Context) error {
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := database.GetBookById(id)
	if book.Id != 0 {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully get book",
			"data":    book,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Book not found",
		"data":    models.Book{},
	})
}

func UpdateBookById(e echo.Context) error {
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var inputUser models.Book
	if err := e.Bind(&inputUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bookUpdate := database.UpdateBookById(inputUser, id)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update book",
		"data":    bookUpdate,
	})
}

func DeleteBookById(e echo.Context) error {
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := database.DeleteBookById(id)
	if book.Id != 0 {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully delete book",
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Book not found",
	})
}

func AddBook(e echo.Context) error {
	var inputUser models.Book

	if err := e.Bind(&inputUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newBook := database.AddBook(inputUser)
	if newBook.Id != 0 {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully add book",
			"data":    newBook,
		})
	}

	return e.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "failed add book",
	})
}
