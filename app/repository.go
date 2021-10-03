package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type repo struct {
	db *sql.DB
}

type Repository interface {
	CreateBook(book Book) error
	GetBooks() ([]Book, error)
	GetBookByID(id int) (Book, error)
	UpdateBook(book Book) error
	DeleteBook(book Book) error
}

func New(db *sql.DB) Repository {
	return &repo{db: db}
}

func (r repo) CreateBook(book Book) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	fmt.Println(book)
	_, err = tx.Exec("INSERT INTO books(title,isbn,author) VALUES ($1,$2,$3)", book.Title, book.ISBN, book.Author)
	if err != nil {
		return err
	}

	return nil
}
func (r repo) GetBooks() ([]Book, error) {
	var results []Book
	rows, err := r.db.Query("SELECT id,title,isbn,author,created_at,updated_at FROM books WHERE is_deleted = false;")
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Book
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.ISBN,
			&item.Author,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return results, err
		}

		results = append(results, item)
	}
	return results, nil
}
func (r repo) GetBookByID(id int) (Book, error) {
	var result Book
	if err := r.db.QueryRow("SELECT id,title,isbn,author,created_at,updated_at FROM books WHERE is_deleted = false;").Scan(
		&result.ID,
		&result.Title,
		&result.ISBN,
		&result.Author,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return result, err
	}

	return result, nil
}
func (r repo) UpdateBook(book Book) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("UPDATE books SET title = $1, isbn = $2, author = $3 WHERE id = $4 AND is_deleted = false", book.Title, book.ISBN, book.Author, book.ID); err != nil {
		return err
	}

	return nil
}
func (r repo) DeleteBook(book Book) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("UPDATE books SET is_deleted = true WHERE id = $1", book.ID); err != nil {
		return err
	}

	return nil
}
