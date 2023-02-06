package models

import (
	"database/sql"
	// "strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Book struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Author		string	`json:"author"`
	PublishYear	int		`json:"publish_year"`
}

func GetBooks(count int) ([]Book, error) {
	rows, err := DB.Query("SELECT id, title, author, publish_year from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishYear)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return books, err
}