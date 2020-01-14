package grpc

import (
	"context"
	"github.com/csandiego/poc-grpc/go/data"
	"github.com/csandiego/poc-grpc/go/protobuf"
)

type BookService struct {
	client protobuf.BookServiceClient
}

func NewBookService(client protobuf.BookServiceClient) *BookService {
	return &BookService{client: client}
}

func (s *BookService) List() ([]data.Book, error) {
	response, err := s.client.List(context.Background(), &protobuf.BookListRequest{})
	if err != nil {
		return nil, err
	}
	books := make([]data.Book, 0, len(response.Books))
	for _, v := range response.Books {
		books = append(books, data.Book{Id: int(v.Id), Title: v.Title, Author: v.Author})
	}
	return books, nil
}

func (s *BookService) Get(id int) (data.Book, error) {
	book := data.Book{}
	response, err := s.client.Get(context.Background(), &protobuf.BookGetRequest{Id: int64(id)})
	if err == nil {
		book.Id = int(response.Book.Id)
		book.Title = response.Book.Title
		book.Author = response.Book.Author
	}
	return book, err
}

func (s *BookService) Create(book data.Book) error {
	_, err := s.client.Create(context.Background(), &protobuf.BookCreateRequest{Book: &protobuf.Book{Title: book.Title, Author: book.Author}})
	return err
}

func (s *BookService) Update(id int, book data.Book) error {
	_, err := s.client.Update(context.Background(), &protobuf.BookUpdateRequest{Id: int64(id), Book: &protobuf.Book{Title: book.Title, Author: book.Author}})
	return err
}

func (s *BookService) Delete(id int) error {
	_, err := s.client.Delete(context.Background(), &protobuf.BookDeleteRequest{Id: int64(id)})
	return err
}
