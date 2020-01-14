package grpc

import (
	"github.com/csandiego/poc-grpc/go/data"
	"github.com/csandiego/poc-grpc/go/protobuf"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBookServiceList(t *testing.T) {
	dto := []*protobuf.Book{
		&protobuf.Book{Id: 1, Title: "abc", Author: "xyz"},
	}
	client := testBookServiceClient{listResponse: &protobuf.BookListResponse{Books: dto}, listErr: nil}
	service := NewBookService(&client)
	books, err := service.List()
	require.Nil(t, err)
	require.Equal(t, dto[0].Id, int64(books[0].Id))
	require.Equal(t, dto[0].Title, books[0].Title)
	require.Equal(t, dto[0].Author, books[0].Author)
}

func TestBookServiceGet(t *testing.T) {
	id := 1
	dto := protobuf.Book{Id: int64(id), Title: "abc", Author: "xyz"}
	client := testBookServiceClient{getResponse: &protobuf.BookGetResponse{Book: &dto}, getErr: nil}
	service := NewBookService(&client)
	book, err := service.Get(id)
	require.Nil(t, err)
	require.Equal(t, int64(id), client.getRequest.Id)
	require.Equal(t, dto.Id, int64(book.Id))
	require.Equal(t, dto.Title, book.Title)
	require.Equal(t, dto.Author, book.Author)
}

func TestBookServiceCreate(t *testing.T) {
	book := data.Book{Title: "abc", Author: "xyz"}
	client := testBookServiceClient{createResponse: &protobuf.BookCreateResponse{}, createErr: nil}
	service := NewBookService(&client)
	err := service.Create(book)
	require.Nil(t, err)
	require.Equal(t, book.Title, client.createRequest.Book.Title)
	require.Equal(t, book.Author, client.createRequest.Book.Author)
}

func TestBookServiceUpdate(t *testing.T) {
	id := 1
	book := data.Book{Title: "abc", Author: "xyz"}
	client := testBookServiceClient{updateResponse: &protobuf.BookUpdateResponse{}, updateErr: nil}
	service := NewBookService(&client)
	err := service.Update(id, book)
	require.Nil(t, err)
	require.Equal(t, int64(id), client.updateRequest.Id)
	require.Equal(t, book.Title, client.updateRequest.Book.Title)
	require.Equal(t, book.Author, client.updateRequest.Book.Author)
}

func TestBookServiceDelete(t *testing.T) {
	id := 1
	client := testBookServiceClient{deleteResponse: &protobuf.BookDeleteResponse{}, deleteErr: nil}
	service := NewBookService(&client)
	err := service.Delete(id)
	require.Nil(t, err)
	require.Equal(t, int64(id), client.deleteRequest.Id)
}
