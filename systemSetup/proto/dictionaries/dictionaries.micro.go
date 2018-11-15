// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/dictionaries/dictionaries.proto

/*
Package ai is a generated protocol buffer package.

It is generated from these files:
	proto/dictionaries/dictionaries.proto

It has these top-level messages:
	Status
	ResponsePage
	RequestPage
	ListRequest
	ListResponse
	DelRequest
	EditResponse
	DictionaryType
	AddDictionaryType
	Dictionary
	AddDictionary
*/
package ai

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Dictionaries service

type DictionariesService interface {
	// 获取字典列表
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	// 创建字典
	Add(ctx context.Context, in *AddDictionaryType, opts ...client.CallOption) (*EditResponse, error)
	// 修改字典
	Edit(ctx context.Context, in *DictionaryType, opts ...client.CallOption) (*EditResponse, error)
	// 删除字典
	Delete(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*EditResponse, error)
}

type dictionariesService struct {
	c    client.Client
	name string
}

func NewDictionariesService(name string, c client.Client) DictionariesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ai"
	}
	return &dictionariesService{
		c:    c,
		name: name,
	}
}

func (c *dictionariesService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Dictionaries.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionariesService) Add(ctx context.Context, in *AddDictionaryType, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "Dictionaries.Add", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionariesService) Edit(ctx context.Context, in *DictionaryType, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "Dictionaries.Edit", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionariesService) Delete(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "Dictionaries.Delete", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dictionaries service

type DictionariesHandler interface {
	// 获取字典列表
	List(context.Context, *ListRequest, *ListResponse) error
	// 创建字典
	Add(context.Context, *AddDictionaryType, *EditResponse) error
	// 修改字典
	Edit(context.Context, *DictionaryType, *EditResponse) error
	// 删除字典
	Delete(context.Context, *DelRequest, *EditResponse) error
}

func RegisterDictionariesHandler(s server.Server, hdlr DictionariesHandler, opts ...server.HandlerOption) error {
	type dictionaries interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Add(ctx context.Context, in *AddDictionaryType, out *EditResponse) error
		Edit(ctx context.Context, in *DictionaryType, out *EditResponse) error
		Delete(ctx context.Context, in *DelRequest, out *EditResponse) error
	}
	type Dictionaries struct {
		dictionaries
	}
	h := &dictionariesHandler{hdlr}
	return s.Handle(s.NewHandler(&Dictionaries{h}, opts...))
}

type dictionariesHandler struct {
	DictionariesHandler
}

func (h *dictionariesHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.DictionariesHandler.List(ctx, in, out)
}

func (h *dictionariesHandler) Add(ctx context.Context, in *AddDictionaryType, out *EditResponse) error {
	return h.DictionariesHandler.Add(ctx, in, out)
}

func (h *dictionariesHandler) Edit(ctx context.Context, in *DictionaryType, out *EditResponse) error {
	return h.DictionariesHandler.Edit(ctx, in, out)
}

func (h *dictionariesHandler) Delete(ctx context.Context, in *DelRequest, out *EditResponse) error {
	return h.DictionariesHandler.Delete(ctx, in, out)
}
