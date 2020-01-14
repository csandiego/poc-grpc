package grpc

import (
	"context"
	"github.com/csandiego/poc-grpc/go/protobuf"
	"google.golang.org/grpc"
)

type testBookServiceClient struct {
	listRequest *protobuf.BookListRequest
	listResponse *protobuf.BookListResponse
	listErr error
	createRequest *protobuf.BookCreateRequest
	createResponse *protobuf.BookCreateResponse
	createErr error
	getRequest *protobuf.BookGetRequest
	getResponse *protobuf.BookGetResponse
	getErr error
	updateRequest *protobuf.BookUpdateRequest
	updateResponse *protobuf.BookUpdateResponse
	updateErr error
	deleteRequest *protobuf.BookDeleteRequest
	deleteResponse *protobuf.BookDeleteResponse
	deleteErr error
}

func (c *testBookServiceClient) List(ctx context.Context, in *protobuf.BookListRequest, opts ...grpc.CallOption) (*protobuf.BookListResponse, error) {
	c.listRequest = in
	return c.listResponse, c.listErr
}

func (c *testBookServiceClient) Create(ctx context.Context, in *protobuf.BookCreateRequest, opts ...grpc.CallOption) (*protobuf.BookCreateResponse, error) {
	c.createRequest = in
	return c.createResponse, c.createErr
}

func (c *testBookServiceClient) Get(ctx context.Context, in *protobuf.BookGetRequest, opts ...grpc.CallOption) (*protobuf.BookGetResponse, error) {
	c.getRequest = in
	return c.getResponse, c.getErr
}

func (c *testBookServiceClient) Update(ctx context.Context, in *protobuf.BookUpdateRequest, opts ...grpc.CallOption) (*protobuf.BookUpdateResponse, error) {
	c.updateRequest = in
	return c.updateResponse, c.updateErr
}

func (c *testBookServiceClient) Delete(ctx context.Context, in *protobuf.BookDeleteRequest, opts ...grpc.CallOption) (*protobuf.BookDeleteResponse, error) {
	c.deleteRequest = in
	return c.deleteResponse, c.deleteErr
}
