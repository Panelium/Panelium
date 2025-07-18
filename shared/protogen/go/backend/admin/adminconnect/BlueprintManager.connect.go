// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: backend/admin/BlueprintManager.proto

package adminconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	admin "panelium/proto_gen_go/backend/admin"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BlueprintManagerServiceName is the fully-qualified name of the BlueprintManagerService service.
	BlueprintManagerServiceName = "backend_admin.BlueprintManagerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BlueprintManagerServiceGetBlueprintsProcedure is the fully-qualified name of the
	// BlueprintManagerService's GetBlueprints RPC.
	BlueprintManagerServiceGetBlueprintsProcedure = "/backend_admin.BlueprintManagerService/GetBlueprints"
	// BlueprintManagerServiceGetBlueprintProcedure is the fully-qualified name of the
	// BlueprintManagerService's GetBlueprint RPC.
	BlueprintManagerServiceGetBlueprintProcedure = "/backend_admin.BlueprintManagerService/GetBlueprint"
	// BlueprintManagerServiceCreateBlueprintProcedure is the fully-qualified name of the
	// BlueprintManagerService's CreateBlueprint RPC.
	BlueprintManagerServiceCreateBlueprintProcedure = "/backend_admin.BlueprintManagerService/CreateBlueprint"
	// BlueprintManagerServiceUpdateBlueprintProcedure is the fully-qualified name of the
	// BlueprintManagerService's UpdateBlueprint RPC.
	BlueprintManagerServiceUpdateBlueprintProcedure = "/backend_admin.BlueprintManagerService/UpdateBlueprint"
	// BlueprintManagerServiceDeleteBlueprintProcedure is the fully-qualified name of the
	// BlueprintManagerService's DeleteBlueprint RPC.
	BlueprintManagerServiceDeleteBlueprintProcedure = "/backend_admin.BlueprintManagerService/DeleteBlueprint"
)

// BlueprintManagerServiceClient is a client for the backend_admin.BlueprintManagerService service.
type BlueprintManagerServiceClient interface {
	GetBlueprints(context.Context, *connect.Request[admin.GetBlueprintsRequest]) (*connect.Response[admin.GetBlueprintsResponse], error)
	GetBlueprint(context.Context, *connect.Request[admin.GetBlueprintRequest]) (*connect.Response[admin.GetBlueprintResponse], error)
	CreateBlueprint(context.Context, *connect.Request[admin.CreateBlueprintRequest]) (*connect.Response[admin.CreateBlueprintResponse], error)
	UpdateBlueprint(context.Context, *connect.Request[admin.UpdateBlueprintRequest]) (*connect.Response[admin.UpdateBlueprintResponse], error)
	DeleteBlueprint(context.Context, *connect.Request[admin.DeleteBlueprintRequest]) (*connect.Response[admin.DeleteBlueprintResponse], error)
}

// NewBlueprintManagerServiceClient constructs a client for the
// backend_admin.BlueprintManagerService service. By default, it uses the Connect protocol with the
// binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use the
// gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBlueprintManagerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BlueprintManagerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	blueprintManagerServiceMethods := admin.File_backend_admin_BlueprintManager_proto.Services().ByName("BlueprintManagerService").Methods()
	return &blueprintManagerServiceClient{
		getBlueprints: connect.NewClient[admin.GetBlueprintsRequest, admin.GetBlueprintsResponse](
			httpClient,
			baseURL+BlueprintManagerServiceGetBlueprintsProcedure,
			connect.WithSchema(blueprintManagerServiceMethods.ByName("GetBlueprints")),
			connect.WithClientOptions(opts...),
		),
		getBlueprint: connect.NewClient[admin.GetBlueprintRequest, admin.GetBlueprintResponse](
			httpClient,
			baseURL+BlueprintManagerServiceGetBlueprintProcedure,
			connect.WithSchema(blueprintManagerServiceMethods.ByName("GetBlueprint")),
			connect.WithClientOptions(opts...),
		),
		createBlueprint: connect.NewClient[admin.CreateBlueprintRequest, admin.CreateBlueprintResponse](
			httpClient,
			baseURL+BlueprintManagerServiceCreateBlueprintProcedure,
			connect.WithSchema(blueprintManagerServiceMethods.ByName("CreateBlueprint")),
			connect.WithClientOptions(opts...),
		),
		updateBlueprint: connect.NewClient[admin.UpdateBlueprintRequest, admin.UpdateBlueprintResponse](
			httpClient,
			baseURL+BlueprintManagerServiceUpdateBlueprintProcedure,
			connect.WithSchema(blueprintManagerServiceMethods.ByName("UpdateBlueprint")),
			connect.WithClientOptions(opts...),
		),
		deleteBlueprint: connect.NewClient[admin.DeleteBlueprintRequest, admin.DeleteBlueprintResponse](
			httpClient,
			baseURL+BlueprintManagerServiceDeleteBlueprintProcedure,
			connect.WithSchema(blueprintManagerServiceMethods.ByName("DeleteBlueprint")),
			connect.WithClientOptions(opts...),
		),
	}
}

// blueprintManagerServiceClient implements BlueprintManagerServiceClient.
type blueprintManagerServiceClient struct {
	getBlueprints   *connect.Client[admin.GetBlueprintsRequest, admin.GetBlueprintsResponse]
	getBlueprint    *connect.Client[admin.GetBlueprintRequest, admin.GetBlueprintResponse]
	createBlueprint *connect.Client[admin.CreateBlueprintRequest, admin.CreateBlueprintResponse]
	updateBlueprint *connect.Client[admin.UpdateBlueprintRequest, admin.UpdateBlueprintResponse]
	deleteBlueprint *connect.Client[admin.DeleteBlueprintRequest, admin.DeleteBlueprintResponse]
}

// GetBlueprints calls backend_admin.BlueprintManagerService.GetBlueprints.
func (c *blueprintManagerServiceClient) GetBlueprints(ctx context.Context, req *connect.Request[admin.GetBlueprintsRequest]) (*connect.Response[admin.GetBlueprintsResponse], error) {
	return c.getBlueprints.CallUnary(ctx, req)
}

// GetBlueprint calls backend_admin.BlueprintManagerService.GetBlueprint.
func (c *blueprintManagerServiceClient) GetBlueprint(ctx context.Context, req *connect.Request[admin.GetBlueprintRequest]) (*connect.Response[admin.GetBlueprintResponse], error) {
	return c.getBlueprint.CallUnary(ctx, req)
}

// CreateBlueprint calls backend_admin.BlueprintManagerService.CreateBlueprint.
func (c *blueprintManagerServiceClient) CreateBlueprint(ctx context.Context, req *connect.Request[admin.CreateBlueprintRequest]) (*connect.Response[admin.CreateBlueprintResponse], error) {
	return c.createBlueprint.CallUnary(ctx, req)
}

// UpdateBlueprint calls backend_admin.BlueprintManagerService.UpdateBlueprint.
func (c *blueprintManagerServiceClient) UpdateBlueprint(ctx context.Context, req *connect.Request[admin.UpdateBlueprintRequest]) (*connect.Response[admin.UpdateBlueprintResponse], error) {
	return c.updateBlueprint.CallUnary(ctx, req)
}

// DeleteBlueprint calls backend_admin.BlueprintManagerService.DeleteBlueprint.
func (c *blueprintManagerServiceClient) DeleteBlueprint(ctx context.Context, req *connect.Request[admin.DeleteBlueprintRequest]) (*connect.Response[admin.DeleteBlueprintResponse], error) {
	return c.deleteBlueprint.CallUnary(ctx, req)
}

// BlueprintManagerServiceHandler is an implementation of the backend_admin.BlueprintManagerService
// service.
type BlueprintManagerServiceHandler interface {
	GetBlueprints(context.Context, *connect.Request[admin.GetBlueprintsRequest]) (*connect.Response[admin.GetBlueprintsResponse], error)
	GetBlueprint(context.Context, *connect.Request[admin.GetBlueprintRequest]) (*connect.Response[admin.GetBlueprintResponse], error)
	CreateBlueprint(context.Context, *connect.Request[admin.CreateBlueprintRequest]) (*connect.Response[admin.CreateBlueprintResponse], error)
	UpdateBlueprint(context.Context, *connect.Request[admin.UpdateBlueprintRequest]) (*connect.Response[admin.UpdateBlueprintResponse], error)
	DeleteBlueprint(context.Context, *connect.Request[admin.DeleteBlueprintRequest]) (*connect.Response[admin.DeleteBlueprintResponse], error)
}

// NewBlueprintManagerServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBlueprintManagerServiceHandler(svc BlueprintManagerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	blueprintManagerServiceMethods := admin.File_backend_admin_BlueprintManager_proto.Services().ByName("BlueprintManagerService").Methods()
	blueprintManagerServiceGetBlueprintsHandler := connect.NewUnaryHandler(
		BlueprintManagerServiceGetBlueprintsProcedure,
		svc.GetBlueprints,
		connect.WithSchema(blueprintManagerServiceMethods.ByName("GetBlueprints")),
		connect.WithHandlerOptions(opts...),
	)
	blueprintManagerServiceGetBlueprintHandler := connect.NewUnaryHandler(
		BlueprintManagerServiceGetBlueprintProcedure,
		svc.GetBlueprint,
		connect.WithSchema(blueprintManagerServiceMethods.ByName("GetBlueprint")),
		connect.WithHandlerOptions(opts...),
	)
	blueprintManagerServiceCreateBlueprintHandler := connect.NewUnaryHandler(
		BlueprintManagerServiceCreateBlueprintProcedure,
		svc.CreateBlueprint,
		connect.WithSchema(blueprintManagerServiceMethods.ByName("CreateBlueprint")),
		connect.WithHandlerOptions(opts...),
	)
	blueprintManagerServiceUpdateBlueprintHandler := connect.NewUnaryHandler(
		BlueprintManagerServiceUpdateBlueprintProcedure,
		svc.UpdateBlueprint,
		connect.WithSchema(blueprintManagerServiceMethods.ByName("UpdateBlueprint")),
		connect.WithHandlerOptions(opts...),
	)
	blueprintManagerServiceDeleteBlueprintHandler := connect.NewUnaryHandler(
		BlueprintManagerServiceDeleteBlueprintProcedure,
		svc.DeleteBlueprint,
		connect.WithSchema(blueprintManagerServiceMethods.ByName("DeleteBlueprint")),
		connect.WithHandlerOptions(opts...),
	)
	return "/backend_admin.BlueprintManagerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BlueprintManagerServiceGetBlueprintsProcedure:
			blueprintManagerServiceGetBlueprintsHandler.ServeHTTP(w, r)
		case BlueprintManagerServiceGetBlueprintProcedure:
			blueprintManagerServiceGetBlueprintHandler.ServeHTTP(w, r)
		case BlueprintManagerServiceCreateBlueprintProcedure:
			blueprintManagerServiceCreateBlueprintHandler.ServeHTTP(w, r)
		case BlueprintManagerServiceUpdateBlueprintProcedure:
			blueprintManagerServiceUpdateBlueprintHandler.ServeHTTP(w, r)
		case BlueprintManagerServiceDeleteBlueprintProcedure:
			blueprintManagerServiceDeleteBlueprintHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBlueprintManagerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBlueprintManagerServiceHandler struct{}

func (UnimplementedBlueprintManagerServiceHandler) GetBlueprints(context.Context, *connect.Request[admin.GetBlueprintsRequest]) (*connect.Response[admin.GetBlueprintsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.BlueprintManagerService.GetBlueprints is not implemented"))
}

func (UnimplementedBlueprintManagerServiceHandler) GetBlueprint(context.Context, *connect.Request[admin.GetBlueprintRequest]) (*connect.Response[admin.GetBlueprintResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.BlueprintManagerService.GetBlueprint is not implemented"))
}

func (UnimplementedBlueprintManagerServiceHandler) CreateBlueprint(context.Context, *connect.Request[admin.CreateBlueprintRequest]) (*connect.Response[admin.CreateBlueprintResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.BlueprintManagerService.CreateBlueprint is not implemented"))
}

func (UnimplementedBlueprintManagerServiceHandler) UpdateBlueprint(context.Context, *connect.Request[admin.UpdateBlueprintRequest]) (*connect.Response[admin.UpdateBlueprintResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.BlueprintManagerService.UpdateBlueprint is not implemented"))
}

func (UnimplementedBlueprintManagerServiceHandler) DeleteBlueprint(context.Context, *connect.Request[admin.DeleteBlueprintRequest]) (*connect.Response[admin.DeleteBlueprintResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.BlueprintManagerService.DeleteBlueprint is not implemented"))
}
