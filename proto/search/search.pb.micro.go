// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/search/search.proto

package pbSearch

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OperateSearch service

func NewOperateSearchEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OperateSearch service

type OperateSearchService interface {
	SearchID(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchContentResults, error)
	SearchContent(ctx context.Context, in *Content, opts ...client.CallOption) (*SearchContentResults, error)
	SearchUser(ctx context.Context, in *Content, opts ...client.CallOption) (*SearchUserResults, error)
}

type operateSearchService struct {
	c    client.Client
	name string
}

func NewOperateSearchService(name string, c client.Client) OperateSearchService {
	return &operateSearchService{
		c:    c,
		name: name,
	}
}

func (c *operateSearchService) SearchID(ctx context.Context, in *UserID, opts ...client.CallOption) (*SearchContentResults, error) {
	req := c.c.NewRequest(c.name, "OperateSearch.SearchID", in)
	out := new(SearchContentResults)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateSearchService) SearchContent(ctx context.Context, in *Content, opts ...client.CallOption) (*SearchContentResults, error) {
	req := c.c.NewRequest(c.name, "OperateSearch.SearchContent", in)
	out := new(SearchContentResults)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateSearchService) SearchUser(ctx context.Context, in *Content, opts ...client.CallOption) (*SearchUserResults, error) {
	req := c.c.NewRequest(c.name, "OperateSearch.SearchUser", in)
	out := new(SearchUserResults)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OperateSearch service

type OperateSearchHandler interface {
	SearchID(context.Context, *UserID, *SearchContentResults) error
	SearchContent(context.Context, *Content, *SearchContentResults) error
	SearchUser(context.Context, *Content, *SearchUserResults) error
}

func RegisterOperateSearchHandler(s server.Server, hdlr OperateSearchHandler, opts ...server.HandlerOption) error {
	type operateSearch interface {
		SearchID(ctx context.Context, in *UserID, out *SearchContentResults) error
		SearchContent(ctx context.Context, in *Content, out *SearchContentResults) error
		SearchUser(ctx context.Context, in *Content, out *SearchUserResults) error
	}
	type OperateSearch struct {
		operateSearch
	}
	h := &operateSearchHandler{hdlr}
	return s.Handle(s.NewHandler(&OperateSearch{h}, opts...))
}

type operateSearchHandler struct {
	OperateSearchHandler
}

func (h *operateSearchHandler) SearchID(ctx context.Context, in *UserID, out *SearchContentResults) error {
	return h.OperateSearchHandler.SearchID(ctx, in, out)
}

func (h *operateSearchHandler) SearchContent(ctx context.Context, in *Content, out *SearchContentResults) error {
	return h.OperateSearchHandler.SearchContent(ctx, in, out)
}

func (h *operateSearchHandler) SearchUser(ctx context.Context, in *Content, out *SearchUserResults) error {
	return h.OperateSearchHandler.SearchUser(ctx, in, out)
}
