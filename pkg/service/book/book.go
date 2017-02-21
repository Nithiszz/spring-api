package book

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Nithiszz/sprint-api/pkg/api"
)

// New creates new book service
func New(config Config) api.BookService {
	return &service{
		client: config.Datastore,
	}
}

// Config is the book service config
type Config struct {
	Datastore *datastore.Client
}

type service struct {
	client *datastore.Client
}

func (s *service) CreateBook(ctx context.Context, req *api.BookRequest) (*api.BookResponse, error) {

	now := time.Now()
	model := bookModel{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	key, err := s.client.Put(ctx, datastore.IncompleteKey(kindBook, nil), &model)
	if err != nil {
		return nil, err
	}

	model.ID = key.ID

	return bookToResponse(&model), nil

}

func (s *service) ListBooks(ctx context.Context) ([]*api.BookResponse, error) {

	var books []*bookModel
	q := datastore.NewQuery(kindBook)

	keys, err := s.client.GetAll(ctx, q, &books)
	if err != nil {
		return nil, err
	}

	for i := range keys {
		books[i].ID = keys[i].ID
	}

	return booksToResponse(books), nil
}

func (s *service) getBook(ctx context.Context, id int64) (*bookModel, error) {

	key := datastore.IDKey(kindBook, id, nil)
	var book bookModel
	err := s.client.Get(ctx, key, &book)
	if err == datastore.ErrNoSuchEntity {
		return nil, api.ErrNotFound
	}

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
	book.UpdatedAt = time.Now()

	key := datastore.IDKey(kindBook, req.ID, nil)
	_, err = s.client.Put(ctx, key, book)

	if err != nil {
		return nil, err
	}

	return bookToResponse(book), nil

}

func (s *service) DeleteBook(ctx context.Context, req *api.BookRequest) error {
	key := datastore.IDKey(kindBook, req.ID, nil)
	err := s.client.Delete(ctx, key)
	if err == datastore.ErrNoSuchEntity {
		return api.ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}
