package book

import (
	"context"
	"time"

	"github.com/Nithiszz/sprint-api/pkg/api"
)

// New creates new book service
func New() api.BookService {
	return &service{}
}

type service struct {
	store  []*bookModel
	lastID int
}

func (s *service) getBook(bookID int) (*bookModel, error) {
	return nil, nil
}

func (s *service) CreateBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {
	s.lastID++
	now := time.Now()
	model := bookModel{
		ID:          s.lastID,
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	s.store = append(s.store, &model)
	return bookToResponse(&model), nil

}

func (s *service) ListBooks(ctx context.Context) ([]*api.BookResponse, error) {

	return booksToResponse(s.store), nil
}

func (s *service) GetBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {
	for _, book := range s.store {
		if book.ID == req.ID {
			return bookToResponse(book), nil
		}
	}
	return nil, api.ErrNotFound
}

func (s *service) UpdateBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {
	for _, book := range s.store {
		if book.ID == req.ID {
			book.Title = req.Title
			book.Description = req.Description
			book.Author = req.Author
			book.UpdatedAt = time.Now()
			return bookToResponse(book), nil
		}
	}

	return nil, api.ErrNotFound
}

func (s *service) DeleteBook(ctx context.Context, req *api.BookRequest) error {
	for i, book := range s.store {
		if book.ID == req.ID {
			s.store = append(s.store[:i], s.store[i+1:]...)
			return nil
		}
	}
	return api.ErrNotFound
}
