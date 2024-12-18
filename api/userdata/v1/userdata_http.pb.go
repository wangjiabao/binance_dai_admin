// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.21.7
// source: api/userdata/v1/userdata.proto

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

const OperationUserdataGetNum = "/api.userdata.v1.Userdata/GetNum"
const OperationUserdataGetUserOrders = "/api.userdata.v1.Userdata/GetUserOrders"
const OperationUserdataGetUsers = "/api.userdata.v1.Userdata/GetUsers"
const OperationUserdataGetUsersIncome = "/api.userdata.v1.Userdata/GetUsersIncome"
const OperationUserdataPullUserIncome = "/api.userdata.v1.Userdata/PullUserIncome"

type UserdataHTTPServer interface {
	GetNum(context.Context, *GetNumRequest) (*GetNumReply, error)
	GetUserOrders(context.Context, *GetUserOrdersRequest) (*GetUserOrdersReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	GetUsersIncome(context.Context, *GetUsersIncomeRequest) (*GetUsersIncomeReply, error)
	PullUserIncome(context.Context, *PullUserIncomeRequest) (*PullUserIncomeReply, error)
}

func RegisterUserdataHTTPServer(s *http.Server, srv UserdataHTTPServer) {
	r := s.Route("/")
	r.GET("/api/binance_dai_admin/get_user_orders", _Userdata_GetUserOrders0_HTTP_Handler(srv))
	r.GET("/api/binance_dai_admin/pull_user_incomes", _Userdata_PullUserIncome0_HTTP_Handler(srv))
	r.GET("/api/binance_dai_admin/get_users", _Userdata_GetUsers0_HTTP_Handler(srv))
	r.GET("/api/binance_dai_admin/get_users_income", _Userdata_GetUsersIncome0_HTTP_Handler(srv))
	r.GET("/api/binance_dai_admin/get_num", _Userdata_GetNum0_HTTP_Handler(srv))
}

func _Userdata_GetUserOrders0_HTTP_Handler(srv UserdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserOrdersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserdataGetUserOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserOrders(ctx, req.(*GetUserOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserOrdersReply)
		return ctx.Result(200, reply)
	}
}

func _Userdata_PullUserIncome0_HTTP_Handler(srv UserdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PullUserIncomeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserdataPullUserIncome)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PullUserIncome(ctx, req.(*PullUserIncomeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PullUserIncomeReply)
		return ctx.Result(200, reply)
	}
}

func _Userdata_GetUsers0_HTTP_Handler(srv UserdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUsersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserdataGetUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUsers(ctx, req.(*GetUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUsersReply)
		return ctx.Result(200, reply)
	}
}

func _Userdata_GetUsersIncome0_HTTP_Handler(srv UserdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUsersIncomeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserdataGetUsersIncome)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUsersIncome(ctx, req.(*GetUsersIncomeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUsersIncomeReply)
		return ctx.Result(200, reply)
	}
}

func _Userdata_GetNum0_HTTP_Handler(srv UserdataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNumRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserdataGetNum)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNum(ctx, req.(*GetNumRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNumReply)
		return ctx.Result(200, reply)
	}
}

type UserdataHTTPClient interface {
	GetNum(ctx context.Context, req *GetNumRequest, opts ...http.CallOption) (rsp *GetNumReply, err error)
	GetUserOrders(ctx context.Context, req *GetUserOrdersRequest, opts ...http.CallOption) (rsp *GetUserOrdersReply, err error)
	GetUsers(ctx context.Context, req *GetUsersRequest, opts ...http.CallOption) (rsp *GetUsersReply, err error)
	GetUsersIncome(ctx context.Context, req *GetUsersIncomeRequest, opts ...http.CallOption) (rsp *GetUsersIncomeReply, err error)
	PullUserIncome(ctx context.Context, req *PullUserIncomeRequest, opts ...http.CallOption) (rsp *PullUserIncomeReply, err error)
}

type UserdataHTTPClientImpl struct {
	cc *http.Client
}

func NewUserdataHTTPClient(client *http.Client) UserdataHTTPClient {
	return &UserdataHTTPClientImpl{client}
}

func (c *UserdataHTTPClientImpl) GetNum(ctx context.Context, in *GetNumRequest, opts ...http.CallOption) (*GetNumReply, error) {
	var out GetNumReply
	pattern := "/api/binance_dai_admin/get_num"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserdataGetNum))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserdataHTTPClientImpl) GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...http.CallOption) (*GetUserOrdersReply, error) {
	var out GetUserOrdersReply
	pattern := "/api/binance_dai_admin/get_user_orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserdataGetUserOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserdataHTTPClientImpl) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...http.CallOption) (*GetUsersReply, error) {
	var out GetUsersReply
	pattern := "/api/binance_dai_admin/get_users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserdataGetUsers))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserdataHTTPClientImpl) GetUsersIncome(ctx context.Context, in *GetUsersIncomeRequest, opts ...http.CallOption) (*GetUsersIncomeReply, error) {
	var out GetUsersIncomeReply
	pattern := "/api/binance_dai_admin/get_users_income"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserdataGetUsersIncome))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserdataHTTPClientImpl) PullUserIncome(ctx context.Context, in *PullUserIncomeRequest, opts ...http.CallOption) (*PullUserIncomeReply, error) {
	var out PullUserIncomeReply
	pattern := "/api/binance_dai_admin/pull_user_incomes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserdataPullUserIncome))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
