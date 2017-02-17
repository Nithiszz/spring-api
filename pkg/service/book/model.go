package book

import (
	"time"

	"github.com/Nithiszz/sprint-api/pkg/api"
)

// Book model is the type for save to database
type bookModel struct {
	ID          int
	Title       string
	Description string
	Author      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func bookToResponse(book *bookModel) *api.BookResponse {
	return &api.BookResponse{

		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		CreatedAt:   book.CreatedAt.Format(time.RFC3339),
	}

}

func booksToResponse(books []*bookModel) []*api.BookResponse {
	result := make([]*api.BookResponse, len(books))

	for i, book := range books {
		result[i] = bookToResponse(book)
	}
	return result
}
