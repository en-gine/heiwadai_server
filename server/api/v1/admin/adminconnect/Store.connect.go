// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/admin/Store.proto

package adminconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	admin "server/api/v1/admin"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StoreControllerName is the fully-qualified name of the StoreController service.
	StoreControllerName = "server.admin.StoreController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StoreControllerGetByIDProcedure is the fully-qualified name of the StoreController's GetByID RPC.
	StoreControllerGetByIDProcedure = "/server.admin.StoreController/GetByID"
	// StoreControllerGetAllProcedure is the fully-qualified name of the StoreController's GetAll RPC.
	StoreControllerGetAllProcedure = "/server.admin.StoreController/GetAll"
	// StoreControllerGetActiveAllProcedure is the fully-qualified name of the StoreController's
	// GetActiveAll RPC.
	StoreControllerGetActiveAllProcedure = "/server.admin.StoreController/GetActiveAll"
	// StoreControllerRegisterProcedure is the fully-qualified name of the StoreController's Register
	// RPC.
	StoreControllerRegisterProcedure = "/server.admin.StoreController/Register"
	// StoreControllerUpdateProcedure is the fully-qualified name of the StoreController's Update RPC.
	StoreControllerUpdateProcedure = "/server.admin.StoreController/Update"
	// StoreControllerRegenQRCodeProcedure is the fully-qualified name of the StoreController's
	// RegenQRCode RPC.
	StoreControllerRegenQRCodeProcedure = "/server.admin.StoreController/RegenQRCode"
	// StoreControllerRegenUnlimitQRCodeProcedure is the fully-qualified name of the StoreController's
	// RegenUnlimitQRCode RPC.
	StoreControllerRegenUnlimitQRCodeProcedure = "/server.admin.StoreController/RegenUnlimitQRCode"
)

// StoreControllerClient is a client for the server.admin.StoreController service.
type StoreControllerClient interface {
	GetByID(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.Store], error)
	GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error)
	GetActiveAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error)
	Register(context.Context, *connect_go.Request[admin.StoreRegisterRequest]) (*connect_go.Response[admin.Store], error)
	Update(context.Context, *connect_go.Request[admin.StoreUpdateRequest]) (*connect_go.Response[admin.Store], error)
	RegenQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.QRResponse], error)
	RegenUnlimitQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.UnlimitQRResponse], error)
}

// NewStoreControllerClient constructs a client for the server.admin.StoreController service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStoreControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StoreControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &storeControllerClient{
		getByID: connect_go.NewClient[admin.StoreIDRequest, admin.Store](
			httpClient,
			baseURL+StoreControllerGetByIDProcedure,
			opts...,
		),
		getAll: connect_go.NewClient[emptypb.Empty, admin.Stores](
			httpClient,
			baseURL+StoreControllerGetAllProcedure,
			opts...,
		),
		getActiveAll: connect_go.NewClient[emptypb.Empty, admin.Stores](
			httpClient,
			baseURL+StoreControllerGetActiveAllProcedure,
			opts...,
		),
		register: connect_go.NewClient[admin.StoreRegisterRequest, admin.Store](
			httpClient,
			baseURL+StoreControllerRegisterProcedure,
			opts...,
		),
		update: connect_go.NewClient[admin.StoreUpdateRequest, admin.Store](
			httpClient,
			baseURL+StoreControllerUpdateProcedure,
			opts...,
		),
		regenQRCode: connect_go.NewClient[admin.StoreIDRequest, admin.QRResponse](
			httpClient,
			baseURL+StoreControllerRegenQRCodeProcedure,
			opts...,
		),
		regenUnlimitQRCode: connect_go.NewClient[admin.StoreIDRequest, admin.UnlimitQRResponse](
			httpClient,
			baseURL+StoreControllerRegenUnlimitQRCodeProcedure,
			opts...,
		),
	}
}

// storeControllerClient implements StoreControllerClient.
type storeControllerClient struct {
	getByID            *connect_go.Client[admin.StoreIDRequest, admin.Store]
	getAll             *connect_go.Client[emptypb.Empty, admin.Stores]
	getActiveAll       *connect_go.Client[emptypb.Empty, admin.Stores]
	register           *connect_go.Client[admin.StoreRegisterRequest, admin.Store]
	update             *connect_go.Client[admin.StoreUpdateRequest, admin.Store]
	regenQRCode        *connect_go.Client[admin.StoreIDRequest, admin.QRResponse]
	regenUnlimitQRCode *connect_go.Client[admin.StoreIDRequest, admin.UnlimitQRResponse]
}

// GetByID calls server.admin.StoreController.GetByID.
func (c *storeControllerClient) GetByID(ctx context.Context, req *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.Store], error) {
	return c.getByID.CallUnary(ctx, req)
}

// GetAll calls server.admin.StoreController.GetAll.
func (c *storeControllerClient) GetAll(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error) {
	return c.getAll.CallUnary(ctx, req)
}

// GetActiveAll calls server.admin.StoreController.GetActiveAll.
func (c *storeControllerClient) GetActiveAll(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error) {
	return c.getActiveAll.CallUnary(ctx, req)
}

// Register calls server.admin.StoreController.Register.
func (c *storeControllerClient) Register(ctx context.Context, req *connect_go.Request[admin.StoreRegisterRequest]) (*connect_go.Response[admin.Store], error) {
	return c.register.CallUnary(ctx, req)
}

// Update calls server.admin.StoreController.Update.
func (c *storeControllerClient) Update(ctx context.Context, req *connect_go.Request[admin.StoreUpdateRequest]) (*connect_go.Response[admin.Store], error) {
	return c.update.CallUnary(ctx, req)
}

// RegenQRCode calls server.admin.StoreController.RegenQRCode.
func (c *storeControllerClient) RegenQRCode(ctx context.Context, req *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.QRResponse], error) {
	return c.regenQRCode.CallUnary(ctx, req)
}

// RegenUnlimitQRCode calls server.admin.StoreController.RegenUnlimitQRCode.
func (c *storeControllerClient) RegenUnlimitQRCode(ctx context.Context, req *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.UnlimitQRResponse], error) {
	return c.regenUnlimitQRCode.CallUnary(ctx, req)
}

// StoreControllerHandler is an implementation of the server.admin.StoreController service.
type StoreControllerHandler interface {
	GetByID(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.Store], error)
	GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error)
	GetActiveAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error)
	Register(context.Context, *connect_go.Request[admin.StoreRegisterRequest]) (*connect_go.Response[admin.Store], error)
	Update(context.Context, *connect_go.Request[admin.StoreUpdateRequest]) (*connect_go.Response[admin.Store], error)
	RegenQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.QRResponse], error)
	RegenUnlimitQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.UnlimitQRResponse], error)
}

// NewStoreControllerHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStoreControllerHandler(svc StoreControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	storeControllerGetByIDHandler := connect_go.NewUnaryHandler(
		StoreControllerGetByIDProcedure,
		svc.GetByID,
		opts...,
	)
	storeControllerGetAllHandler := connect_go.NewUnaryHandler(
		StoreControllerGetAllProcedure,
		svc.GetAll,
		opts...,
	)
	storeControllerGetActiveAllHandler := connect_go.NewUnaryHandler(
		StoreControllerGetActiveAllProcedure,
		svc.GetActiveAll,
		opts...,
	)
	storeControllerRegisterHandler := connect_go.NewUnaryHandler(
		StoreControllerRegisterProcedure,
		svc.Register,
		opts...,
	)
	storeControllerUpdateHandler := connect_go.NewUnaryHandler(
		StoreControllerUpdateProcedure,
		svc.Update,
		opts...,
	)
	storeControllerRegenQRCodeHandler := connect_go.NewUnaryHandler(
		StoreControllerRegenQRCodeProcedure,
		svc.RegenQRCode,
		opts...,
	)
	storeControllerRegenUnlimitQRCodeHandler := connect_go.NewUnaryHandler(
		StoreControllerRegenUnlimitQRCodeProcedure,
		svc.RegenUnlimitQRCode,
		opts...,
	)
	return "/server.admin.StoreController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StoreControllerGetByIDProcedure:
			storeControllerGetByIDHandler.ServeHTTP(w, r)
		case StoreControllerGetAllProcedure:
			storeControllerGetAllHandler.ServeHTTP(w, r)
		case StoreControllerGetActiveAllProcedure:
			storeControllerGetActiveAllHandler.ServeHTTP(w, r)
		case StoreControllerRegisterProcedure:
			storeControllerRegisterHandler.ServeHTTP(w, r)
		case StoreControllerUpdateProcedure:
			storeControllerUpdateHandler.ServeHTTP(w, r)
		case StoreControllerRegenQRCodeProcedure:
			storeControllerRegenQRCodeHandler.ServeHTTP(w, r)
		case StoreControllerRegenUnlimitQRCodeProcedure:
			storeControllerRegenUnlimitQRCodeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStoreControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedStoreControllerHandler struct{}

func (UnimplementedStoreControllerHandler) GetByID(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.Store], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.GetByID is not implemented"))
}

func (UnimplementedStoreControllerHandler) GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.GetAll is not implemented"))
}

func (UnimplementedStoreControllerHandler) GetActiveAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[admin.Stores], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.GetActiveAll is not implemented"))
}

func (UnimplementedStoreControllerHandler) Register(context.Context, *connect_go.Request[admin.StoreRegisterRequest]) (*connect_go.Response[admin.Store], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.Register is not implemented"))
}

func (UnimplementedStoreControllerHandler) Update(context.Context, *connect_go.Request[admin.StoreUpdateRequest]) (*connect_go.Response[admin.Store], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.Update is not implemented"))
}

func (UnimplementedStoreControllerHandler) RegenQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.QRResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.RegenQRCode is not implemented"))
}

func (UnimplementedStoreControllerHandler) RegenUnlimitQRCode(context.Context, *connect_go.Request[admin.StoreIDRequest]) (*connect_go.Response[admin.UnlimitQRResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.StoreController.RegenUnlimitQRCode is not implemented"))
}
