package book

import (
	"time"

	"github.com/Nithiszz/sprint-api/pkg/api"
)

const kindBook = "Book"

// Book model is the type for save to database
type bookModel struct {
	ID          int64 `datastore:"-"`
	Title       string
	Description string `datastore:",noindex"`
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
