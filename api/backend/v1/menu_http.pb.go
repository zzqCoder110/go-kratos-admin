// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.1--rc1
// source: backend/v1/menu.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenuMenuCreate = "/backend.v1.Menu/MenuCreate"
const OperationMenuMenuDelete = "/backend.v1.Menu/MenuDelete"
const OperationMenuMenuList = "/backend.v1.Menu/MenuList"
const OperationMenuMenuUpdate = "/backend.v1.Menu/MenuUpdate"

type MenuHTTPServer interface {
	// MenuCreate Menu
	MenuCreate(context.Context, *MenuCreateReq) (*emptypb.Empty, error)
	MenuDelete(context.Context, *MenuDeleteReq) (*emptypb.Empty, error)
	MenuList(context.Context, *emptypb.Empty) (*MenuGetListRep, error)
	MenuUpdate(context.Context, *MenuUpdateReq) (*emptypb.Empty, error)
}

func RegisterMenuHTTPServer(s *http.Server, srv MenuHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/menu", _Menu_MenuCreate0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/menu", _Menu_MenuDelete0_HTTP_Handler(srv))
	r.PUT("/admin/v1/menu", _Menu_MenuUpdate0_HTTP_Handler(srv))
	r.GET("/admin/v1/menu/list", _Menu_MenuList0_HTTP_Handler(srv))
}

func _Menu_MenuCreate0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MenuCreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuMenuCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MenuCreate(ctx, req.(*MenuCreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Menu_MenuDelete0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MenuDeleteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuMenuDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MenuDelete(ctx, req.(*MenuDeleteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Menu_MenuUpdate0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MenuUpdateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuMenuUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MenuUpdate(ctx, req.(*MenuUpdateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Menu_MenuList0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuMenuList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MenuList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MenuGetListRep)
		return ctx.Result(200, reply)
	}
}

type MenuHTTPClient interface {
	MenuCreate(ctx context.Context, req *MenuCreateReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	MenuDelete(ctx context.Context, req *MenuDeleteReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	MenuList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *MenuGetListRep, err error)
	MenuUpdate(ctx context.Context, req *MenuUpdateReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type MenuHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuHTTPClient(client *http.Client) MenuHTTPClient {
	return &MenuHTTPClientImpl{client}
}

func (c *MenuHTTPClientImpl) MenuCreate(ctx context.Context, in *MenuCreateReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuMenuCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/menu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuMenuDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) MenuList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*MenuGetListRep, error) {
	var out MenuGetListRep
	pattern := "/admin/v1/menu/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuMenuList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuMenuUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
