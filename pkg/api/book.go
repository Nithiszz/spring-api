package api

import "context"

type BookService interface {
	CreateBook(context.Context, *BookRequest) (*BookResponse, error)
	ListBooks(context.Context) ([]*BookResponse, error)
	GetBook(context.Context, *BookRequest) (*BookResponse, error)
	UpdateBook(context.Context, *BookRequest) (*BookResponse, error)
	DeleteBook(context.Context, *BookRequest) error
}

// BookRequest is the book type for parse from user's Request
type BookRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

// BookResponse is the book type for return to user
type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	CreatedAt   string `json:"createdAt"`
}
