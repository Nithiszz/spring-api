package book

import (
	"context"

	"github.com/Nithiszz/sprint-api/pkg/api"
	"github.com/acoshift/ds"
)

// New creates new book service
func New(config Config) api.BookService {
	return &service{
		client: config.Datastore,
	}
}

// Config is the book service config
type Config struct {
	Datastore *ds.Client
}

type service struct {
	client *ds.Client
}

func (s *service) CreateBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {

	model := bookModel{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
	}
	err := s.client.SaveModel(ctx, kindBook, &model)
	if err != nil {
		return nil, err
	}

	return bookToResponse(&model), nil
}

func (s *service) ListBooks(ctx context.Context) ([]*api.BookResponse, error) {

	var books []*bookModel

	err := s.client.Query(ctx, kindBook, &books, ds.Order("-CreatedAt"))

	if err != nil {
		return nil, err
	}

	return booksToResponse(books), nil
}

func (s *service) getBook(ctx context.Context, id int64) (*bookModel, error) {

	var book bookModel

	err := s.client.GetByID(ctx, kindBook, id, &book)

	if ds.NotFound(err) {
		return nil, api.ErrNotFound
	}
	err = ds.IgnoreFieldMismatch(err)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *service) GetBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {

	book, err := s.getBook(ctx, req.ID)

	if err != nil {
		return nil, err
	}

	return bookToResponse(book), nil
}

func (s *service) UpdateBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {

	book, err := s.getBook(ctx, req.ID)

	if err != nil {
		return nil, err
	}

	book.Title = req.Title
	book.Description = req.Description
	book.Author = req.Author

	err = s.client.SaveModel(ctx, "", book)

	if err != nil {
		return nil, err
	}

	return bookToResponse(book), nil

}

func (s *service) DeleteBook(ctx context.Context, req *api.BookRequest) error {

	err := s.client.DeleteByID(ctx, kindBook, req.ID)
	if ds.NotFound(err) {
		return api.ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}
