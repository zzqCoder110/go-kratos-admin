// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.1--rc1
// source: backend/v1/officer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOfficerCreate = "/backend.v1.Officer/Create"
const OperationOfficerLogin = "/backend.v1.Officer/Login"

type OfficerHTTPServer interface {
	Create(context.Context, *CreateReq) (*CreateRep, error)
	Login(context.Context, *LoginReq) (*LoginRep, error)
}

func RegisterOfficerHTTPServer(s *http.Server, srv OfficerHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/officer", _Officer_Create0_HTTP_Handler(srv))
	r.POST("/admin/v1/login", _Officer_Login0_HTTP_Handler(srv))
}

func _Officer_Create0_HTTP_Handler(srv OfficerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOfficerCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRep)
		return ctx.Result(200, reply)
	}
}

func _Officer_Login0_HTTP_Handler(srv OfficerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOfficerLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginRep)
		return ctx.Result(200, reply)
	}
}

type OfficerHTTPClient interface {
	Create(ctx context.Context, req *CreateReq, opts ...http.CallOption) (rsp *CreateRep, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginRep, err error)
}

type OfficerHTTPClientImpl struct {
	cc *http.Client
}

func NewOfficerHTTPClient(client *http.Client) OfficerHTTPClient {
	return &OfficerHTTPClientImpl{client}
}

func (c *OfficerHTTPClientImpl) Create(ctx context.Context, in *CreateReq, opts ...http.CallOption) (*CreateRep, error) {
	var out CreateRep
	pattern := "/admin/v1/officer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOfficerCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OfficerHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginRep, error) {
	var out LoginRep
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOfficerLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
