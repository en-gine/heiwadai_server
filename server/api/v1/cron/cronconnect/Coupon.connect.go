// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/cron/Coupon.proto

package cronconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	cron "server/api/v1/cron"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CronCouponControllerName is the fully-qualified name of the CronCouponController service.
	CronCouponControllerName = "server.cron.CronCouponController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CronCouponControllerBulkIssueBirthdayCouponProcedure is the fully-qualified name of the
	// CronCouponController's BulkIssueBirthdayCoupon RPC.
	CronCouponControllerBulkIssueBirthdayCouponProcedure = "/server.cron.CronCouponController/BulkIssueBirthdayCoupon"
)

// CronCouponControllerClient is a client for the server.cron.CronCouponController service.
type CronCouponControllerClient interface {
	BulkIssueBirthdayCoupon(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[cron.AffectedCountResponse], error)
}

// NewCronCouponControllerClient constructs a client for the server.cron.CronCouponController
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCronCouponControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CronCouponControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cronCouponControllerClient{
		bulkIssueBirthdayCoupon: connect_go.NewClient[emptypb.Empty, cron.AffectedCountResponse](
			httpClient,
			baseURL+CronCouponControllerBulkIssueBirthdayCouponProcedure,
			opts...,
		),
	}
}

// cronCouponControllerClient implements CronCouponControllerClient.
type cronCouponControllerClient struct {
	bulkIssueBirthdayCoupon *connect_go.Client[emptypb.Empty, cron.AffectedCountResponse]
}

// BulkIssueBirthdayCoupon calls server.cron.CronCouponController.BulkIssueBirthdayCoupon.
func (c *cronCouponControllerClient) BulkIssueBirthdayCoupon(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[cron.AffectedCountResponse], error) {
	return c.bulkIssueBirthdayCoupon.CallUnary(ctx, req)
}

// CronCouponControllerHandler is an implementation of the server.cron.CronCouponController service.
type CronCouponControllerHandler interface {
	BulkIssueBirthdayCoupon(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[cron.AffectedCountResponse], error)
}

// NewCronCouponControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCronCouponControllerHandler(svc CronCouponControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	cronCouponControllerBulkIssueBirthdayCouponHandler := connect_go.NewUnaryHandler(
		CronCouponControllerBulkIssueBirthdayCouponProcedure,
		svc.BulkIssueBirthdayCoupon,
		opts...,
	)
	return "/server.cron.CronCouponController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CronCouponControllerBulkIssueBirthdayCouponProcedure:
			cronCouponControllerBulkIssueBirthdayCouponHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCronCouponControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedCronCouponControllerHandler struct{}

func (UnimplementedCronCouponControllerHandler) BulkIssueBirthdayCoupon(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[cron.AffectedCountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.cron.CronCouponController.BulkIssueBirthdayCoupon is not implemented"))
}