package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tyasrush/golang-simple-api/app"
)

func main() {
	mDb, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=testing port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	repo := app.New(mDb)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("testing aja"))
	})
	http.HandleFunc("/books", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var param app.CreateBook
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&param); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			if err := repo.CreateBook(app.Book{Title: param.Title, ISBN: param.ISBN, Author: param.Author}); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Create book berhasil!"))
			return
		case http.MethodGet:
			books, err := repo.GetBooks()
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			byteResponse, err := json.Marshal(books)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write(byteResponse)
			return
		case http.MethodPut:
			var param app.UpdateBook
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&param); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			if err := repo.UpdateBook(app.Book{ID: param.ID, Title: param.Title, Author: param.Author, ISBN: param.ISBN}); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Update book berhasil dong"))
			return
		case http.MethodDelete:
			var param app.DeleteBook
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&param); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			if err := repo.DeleteBook(app.Book{ID: param.ID}); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("Internal server error aja"))
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Hapus book berhasil dong"))
			return
		}
	})
	http.ListenAndServe(":8000", nil)
}
