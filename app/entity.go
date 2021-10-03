package app

import "time"

type Book struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	ISBN      string `json:"isbn"`
	Author    string `json:"author"`
	IsDeleted bool
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBook struct {
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
}

type UpdateBook struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
}

type DeleteBook struct {
	ID int64 `json:"id"`
}
