package grpc

import (
	"context"
	"github.com/csandiego/poc-grpc/go/data"
	"github.com/csandiego/poc-grpc/go/protobuf"
	"github.com/csandiego/poc-grpc/go/service"
)

type BookServiceServer struct {
	service service.BookService
}

func NewBookServiceServer(service service.BookService) *BookServiceServer {
	return &BookServiceServer{service: service}
}

func (s *BookServiceServer) List(ctx context.Context, request *protobuf.BookListRequest) (*protobuf.BookListResponse, error) {
	books, err := s.service.List()
	if err != nil {
		return nil, err
	}
	dto := []*protobuf.Book{}
	for _, v := range books {
		dto = append(dto, &protobuf.Book{Id: int64(v.Id), Title: v.Title, Author: v.Author})
	}
	return &protobuf.BookListResponse{Books: dto}, nil
}

func (s *BookServiceServer) Create(ctx context.Context, request *protobuf.BookCreateRequest) (*protobuf.BookCreateResponse, error) {
	err := s.service.Create(data.Book{Title: request.Book.Title, Author: request.Book.Author})
	if err != nil {
		return nil, err
	}
	return &protobuf.BookCreateResponse{}, nil
}

func (s *BookServiceServer) Get(ctx context.Context, request *protobuf.BookGetRequest) (*protobuf.BookGetResponse, error) {
	book, err := s.service.Get(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &protobuf.BookGetResponse{Book: &protobuf.Book{Id: int64(book.Id), Title: book.Title, Author: book.Author}}, nil
}

func (s *BookServiceServer) Update(ctx context.Context, request *protobuf.BookUpdateRequest) (*protobuf.BookUpdateResponse, error) {
	err := s.service.Update(int(request.Id), data.Book{Title: request.Book.Title, Author: request.Book.Author})
	if err != nil {
		return nil, err
	}
	return &protobuf.BookUpdateResponse{}, nil
}

func (s *BookServiceServer) Delete(ctx context.Context, request *protobuf.BookDeleteRequest) (*protobuf.BookDeleteResponse, error) {
	err := s.service.Delete(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &protobuf.BookDeleteResponse{}, nil
}
