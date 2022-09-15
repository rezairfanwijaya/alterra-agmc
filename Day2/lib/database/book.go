package database

import "altera/Day2/models"

func GetAllBook() []models.Book {
	books := []models.Book{
		{
			Id:     1,
			Title:  "Book 1",
			Page:   100,
			Writer: "Writer 1",
		}, {
			Id:     2,
			Title:  "Book 2",
			Page:   150,
			Writer: "Writer 2",
		}, {
			Id:     3,
			Title:  "Book 3",
			Page:   200,
			Writer: "Writer 3",
		},
	}

	return books
}

func GetBookById(id int) models.Book {
	books := GetAllBook()

	for _, book := range books {
		if id == book.Id {
			return models.Book{
				Id:     book.Id,
				Title:  book.Title,
				Page:   book.Page,
				Writer: book.Writer,
			}
		}
	}

	return models.Book{}
}

func UpdateBookById(inputUser models.Book, id int) models.Book {
	book := GetBookById(id)
	if inputUser.Page != 0 {
		book.Page = inputUser.Page
	}

	if inputUser.Title != "" {
		book.Title = inputUser.Title
	}

	if inputUser.Writer != "" {
		book.Writer = inputUser.Writer
	}

	return book
}

func DeleteBookById(id int) models.Book {
	book := GetBookById(id)

	return book
}

func AddBook(input models.Book) models.Book {
	var book models.Book
	book.Id = 4
	book.Page = input.Page
	book.Title = input.Title
	book.Writer = input.Writer

	return book
}
