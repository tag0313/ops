// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/follower/follower.proto

package pbFollower

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

// Api Endpoints for OperateFollow service

func NewOperateFollowEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OperateFollow service

type OperateFollowService interface {
	Follow(ctx context.Context, in *Follower, opts ...client.CallOption) (*FollowNum, error)
	CanalFollow(ctx context.Context, in *Follower, opts ...client.CallOption) (*FollowNum, error)
	QueryFollowListAll(ctx context.Context, in *OopFollowListReq, opts ...client.CallOption) (*OopFollowListResp, error)
	QueryFollowingList(ctx context.Context, in *FollowListReq, opts ...client.CallOption) (*FollowListResp, error)
	QueryFollowedList(ctx context.Context, in *FollowListReq, opts ...client.CallOption) (*FollowListResp, error)
	WhoFollowingMe(ctx context.Context, in *RelationReq, opts ...client.CallOption) (*RelationResp, error)
	WhoFollowedMe(ctx context.Context, in *RelationReq, opts ...client.CallOption) (*RelationResp, error)
}

type operateFollowService struct {
	c    client.Client
	name string
}

func NewOperateFollowService(name string, c client.Client) OperateFollowService {
	return &operateFollowService{
		c:    c,
		name: name,
	}
}

func (c *operateFollowService) Follow(ctx context.Context, in *Follower, opts ...client.CallOption) (*FollowNum, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.Follow", in)
	out := new(FollowNum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) CanalFollow(ctx context.Context, in *Follower, opts ...client.CallOption) (*FollowNum, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.CanalFollow", in)
	out := new(FollowNum)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) QueryFollowListAll(ctx context.Context, in *OopFollowListReq, opts ...client.CallOption) (*OopFollowListResp, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.QueryFollowListAll", in)
	out := new(OopFollowListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) QueryFollowingList(ctx context.Context, in *FollowListReq, opts ...client.CallOption) (*FollowListResp, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.QueryFollowingList", in)
	out := new(FollowListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) QueryFollowedList(ctx context.Context, in *FollowListReq, opts ...client.CallOption) (*FollowListResp, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.QueryFollowedList", in)
	out := new(FollowListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) WhoFollowingMe(ctx context.Context, in *RelationReq, opts ...client.CallOption) (*RelationResp, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.WhoFollowingMe", in)
	out := new(RelationResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operateFollowService) WhoFollowedMe(ctx context.Context, in *RelationReq, opts ...client.CallOption) (*RelationResp, error) {
	req := c.c.NewRequest(c.name, "OperateFollow.WhoFollowedMe", in)
	out := new(RelationResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OperateFollow service

type OperateFollowHandler interface {
	Follow(context.Context, *Follower, *FollowNum) error
	CanalFollow(context.Context, *Follower, *FollowNum) error
	QueryFollowListAll(context.Context, *OopFollowListReq, *OopFollowListResp) error
	QueryFollowingList(context.Context, *FollowListReq, *FollowListResp) error
	QueryFollowedList(context.Context, *FollowListReq, *FollowListResp) error
	WhoFollowingMe(context.Context, *RelationReq, *RelationResp) error
	WhoFollowedMe(context.Context, *RelationReq, *RelationResp) error
}

func RegisterOperateFollowHandler(s server.Server, hdlr OperateFollowHandler, opts ...server.HandlerOption) error {
	type operateFollow interface {
		Follow(ctx context.Context, in *Follower, out *FollowNum) error
		CanalFollow(ctx context.Context, in *Follower, out *FollowNum) error
		QueryFollowListAll(ctx context.Context, in *OopFollowListReq, out *OopFollowListResp) error
		QueryFollowingList(ctx context.Context, in *FollowListReq, out *FollowListResp) error
		QueryFollowedList(ctx context.Context, in *FollowListReq, out *FollowListResp) error
		WhoFollowingMe(ctx context.Context, in *RelationReq, out *RelationResp) error
		WhoFollowedMe(ctx context.Context, in *RelationReq, out *RelationResp) error
	}
	type OperateFollow struct {
		operateFollow
	}
	h := &operateFollowHandler{hdlr}
	return s.Handle(s.NewHandler(&OperateFollow{h}, opts...))
}

type operateFollowHandler struct {
	OperateFollowHandler
}

func (h *operateFollowHandler) Follow(ctx context.Context, in *Follower, out *FollowNum) error {
	return h.OperateFollowHandler.Follow(ctx, in, out)
}

func (h *operateFollowHandler) CanalFollow(ctx context.Context, in *Follower, out *FollowNum) error {
	return h.OperateFollowHandler.CanalFollow(ctx, in, out)
}

func (h *operateFollowHandler) QueryFollowListAll(ctx context.Context, in *OopFollowListReq, out *OopFollowListResp) error {
	return h.OperateFollowHandler.QueryFollowListAll(ctx, in, out)
}

func (h *operateFollowHandler) QueryFollowingList(ctx context.Context, in *FollowListReq, out *FollowListResp) error {
	return h.OperateFollowHandler.QueryFollowingList(ctx, in, out)
}

func (h *operateFollowHandler) QueryFollowedList(ctx context.Context, in *FollowListReq, out *FollowListResp) error {
	return h.OperateFollowHandler.QueryFollowedList(ctx, in, out)
}

func (h *operateFollowHandler) WhoFollowingMe(ctx context.Context, in *RelationReq, out *RelationResp) error {
	return h.OperateFollowHandler.WhoFollowingMe(ctx, in, out)
}

func (h *operateFollowHandler) WhoFollowedMe(ctx context.Context, in *RelationReq, out *RelationResp) error {
	return h.OperateFollowHandler.WhoFollowedMe(ctx, in, out)
}
