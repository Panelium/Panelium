// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: backend/admin/NodeAllocationManager.proto

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
	// NodeAllocationManagerServiceName is the fully-qualified name of the NodeAllocationManagerService
	// service.
	NodeAllocationManagerServiceName = "backend_admin.NodeAllocationManagerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NodeAllocationManagerServiceGetNodeAllocationsProcedure is the fully-qualified name of the
	// NodeAllocationManagerService's GetNodeAllocations RPC.
	NodeAllocationManagerServiceGetNodeAllocationsProcedure = "/backend_admin.NodeAllocationManagerService/GetNodeAllocations"
	// NodeAllocationManagerServiceGetNodeAllocationProcedure is the fully-qualified name of the
	// NodeAllocationManagerService's GetNodeAllocation RPC.
	NodeAllocationManagerServiceGetNodeAllocationProcedure = "/backend_admin.NodeAllocationManagerService/GetNodeAllocation"
	// NodeAllocationManagerServiceCreateNodeAllocationProcedure is the fully-qualified name of the
	// NodeAllocationManagerService's CreateNodeAllocation RPC.
	NodeAllocationManagerServiceCreateNodeAllocationProcedure = "/backend_admin.NodeAllocationManagerService/CreateNodeAllocation"
	// NodeAllocationManagerServiceUpdateNodeAllocationProcedure is the fully-qualified name of the
	// NodeAllocationManagerService's UpdateNodeAllocation RPC.
	NodeAllocationManagerServiceUpdateNodeAllocationProcedure = "/backend_admin.NodeAllocationManagerService/UpdateNodeAllocation"
	// NodeAllocationManagerServiceDeleteNodeAllocationProcedure is the fully-qualified name of the
	// NodeAllocationManagerService's DeleteNodeAllocation RPC.
	NodeAllocationManagerServiceDeleteNodeAllocationProcedure = "/backend_admin.NodeAllocationManagerService/DeleteNodeAllocation"
)

// NodeAllocationManagerServiceClient is a client for the backend_admin.NodeAllocationManagerService
// service.
type NodeAllocationManagerServiceClient interface {
	GetNodeAllocations(context.Context, *connect.Request[admin.GetNodeAllocationsRequest]) (*connect.Response[admin.GetNodeAllocationsResponse], error)
	GetNodeAllocation(context.Context, *connect.Request[admin.GetNodeAllocationRequest]) (*connect.Response[admin.GetNodeAllocationResponse], error)
	CreateNodeAllocation(context.Context, *connect.Request[admin.CreateNodeAllocationRequest]) (*connect.Response[admin.CreateNodeAllocationResponse], error)
	UpdateNodeAllocation(context.Context, *connect.Request[admin.UpdateNodeAllocationRequest]) (*connect.Response[admin.UpdateNodeAllocationResponse], error)
	DeleteNodeAllocation(context.Context, *connect.Request[admin.DeleteNodeAllocationRequest]) (*connect.Response[admin.DeleteNodeAllocationResponse], error)
}

// NewNodeAllocationManagerServiceClient constructs a client for the
// backend_admin.NodeAllocationManagerService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNodeAllocationManagerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NodeAllocationManagerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	nodeAllocationManagerServiceMethods := admin.File_backend_admin_NodeAllocationManager_proto.Services().ByName("NodeAllocationManagerService").Methods()
	return &nodeAllocationManagerServiceClient{
		getNodeAllocations: connect.NewClient[admin.GetNodeAllocationsRequest, admin.GetNodeAllocationsResponse](
			httpClient,
			baseURL+NodeAllocationManagerServiceGetNodeAllocationsProcedure,
			connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("GetNodeAllocations")),
			connect.WithClientOptions(opts...),
		),
		getNodeAllocation: connect.NewClient[admin.GetNodeAllocationRequest, admin.GetNodeAllocationResponse](
			httpClient,
			baseURL+NodeAllocationManagerServiceGetNodeAllocationProcedure,
			connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("GetNodeAllocation")),
			connect.WithClientOptions(opts...),
		),
		createNodeAllocation: connect.NewClient[admin.CreateNodeAllocationRequest, admin.CreateNodeAllocationResponse](
			httpClient,
			baseURL+NodeAllocationManagerServiceCreateNodeAllocationProcedure,
			connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("CreateNodeAllocation")),
			connect.WithClientOptions(opts...),
		),
		updateNodeAllocation: connect.NewClient[admin.UpdateNodeAllocationRequest, admin.UpdateNodeAllocationResponse](
			httpClient,
			baseURL+NodeAllocationManagerServiceUpdateNodeAllocationProcedure,
			connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("UpdateNodeAllocation")),
			connect.WithClientOptions(opts...),
		),
		deleteNodeAllocation: connect.NewClient[admin.DeleteNodeAllocationRequest, admin.DeleteNodeAllocationResponse](
			httpClient,
			baseURL+NodeAllocationManagerServiceDeleteNodeAllocationProcedure,
			connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("DeleteNodeAllocation")),
			connect.WithClientOptions(opts...),
		),
	}
}

// nodeAllocationManagerServiceClient implements NodeAllocationManagerServiceClient.
type nodeAllocationManagerServiceClient struct {
	getNodeAllocations   *connect.Client[admin.GetNodeAllocationsRequest, admin.GetNodeAllocationsResponse]
	getNodeAllocation    *connect.Client[admin.GetNodeAllocationRequest, admin.GetNodeAllocationResponse]
	createNodeAllocation *connect.Client[admin.CreateNodeAllocationRequest, admin.CreateNodeAllocationResponse]
	updateNodeAllocation *connect.Client[admin.UpdateNodeAllocationRequest, admin.UpdateNodeAllocationResponse]
	deleteNodeAllocation *connect.Client[admin.DeleteNodeAllocationRequest, admin.DeleteNodeAllocationResponse]
}

// GetNodeAllocations calls backend_admin.NodeAllocationManagerService.GetNodeAllocations.
func (c *nodeAllocationManagerServiceClient) GetNodeAllocations(ctx context.Context, req *connect.Request[admin.GetNodeAllocationsRequest]) (*connect.Response[admin.GetNodeAllocationsResponse], error) {
	return c.getNodeAllocations.CallUnary(ctx, req)
}

// GetNodeAllocation calls backend_admin.NodeAllocationManagerService.GetNodeAllocation.
func (c *nodeAllocationManagerServiceClient) GetNodeAllocation(ctx context.Context, req *connect.Request[admin.GetNodeAllocationRequest]) (*connect.Response[admin.GetNodeAllocationResponse], error) {
	return c.getNodeAllocation.CallUnary(ctx, req)
}

// CreateNodeAllocation calls backend_admin.NodeAllocationManagerService.CreateNodeAllocation.
func (c *nodeAllocationManagerServiceClient) CreateNodeAllocation(ctx context.Context, req *connect.Request[admin.CreateNodeAllocationRequest]) (*connect.Response[admin.CreateNodeAllocationResponse], error) {
	return c.createNodeAllocation.CallUnary(ctx, req)
}

// UpdateNodeAllocation calls backend_admin.NodeAllocationManagerService.UpdateNodeAllocation.
func (c *nodeAllocationManagerServiceClient) UpdateNodeAllocation(ctx context.Context, req *connect.Request[admin.UpdateNodeAllocationRequest]) (*connect.Response[admin.UpdateNodeAllocationResponse], error) {
	return c.updateNodeAllocation.CallUnary(ctx, req)
}

// DeleteNodeAllocation calls backend_admin.NodeAllocationManagerService.DeleteNodeAllocation.
func (c *nodeAllocationManagerServiceClient) DeleteNodeAllocation(ctx context.Context, req *connect.Request[admin.DeleteNodeAllocationRequest]) (*connect.Response[admin.DeleteNodeAllocationResponse], error) {
	return c.deleteNodeAllocation.CallUnary(ctx, req)
}

// NodeAllocationManagerServiceHandler is an implementation of the
// backend_admin.NodeAllocationManagerService service.
type NodeAllocationManagerServiceHandler interface {
	GetNodeAllocations(context.Context, *connect.Request[admin.GetNodeAllocationsRequest]) (*connect.Response[admin.GetNodeAllocationsResponse], error)
	GetNodeAllocation(context.Context, *connect.Request[admin.GetNodeAllocationRequest]) (*connect.Response[admin.GetNodeAllocationResponse], error)
	CreateNodeAllocation(context.Context, *connect.Request[admin.CreateNodeAllocationRequest]) (*connect.Response[admin.CreateNodeAllocationResponse], error)
	UpdateNodeAllocation(context.Context, *connect.Request[admin.UpdateNodeAllocationRequest]) (*connect.Response[admin.UpdateNodeAllocationResponse], error)
	DeleteNodeAllocation(context.Context, *connect.Request[admin.DeleteNodeAllocationRequest]) (*connect.Response[admin.DeleteNodeAllocationResponse], error)
}

// NewNodeAllocationManagerServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNodeAllocationManagerServiceHandler(svc NodeAllocationManagerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	nodeAllocationManagerServiceMethods := admin.File_backend_admin_NodeAllocationManager_proto.Services().ByName("NodeAllocationManagerService").Methods()
	nodeAllocationManagerServiceGetNodeAllocationsHandler := connect.NewUnaryHandler(
		NodeAllocationManagerServiceGetNodeAllocationsProcedure,
		svc.GetNodeAllocations,
		connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("GetNodeAllocations")),
		connect.WithHandlerOptions(opts...),
	)
	nodeAllocationManagerServiceGetNodeAllocationHandler := connect.NewUnaryHandler(
		NodeAllocationManagerServiceGetNodeAllocationProcedure,
		svc.GetNodeAllocation,
		connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("GetNodeAllocation")),
		connect.WithHandlerOptions(opts...),
	)
	nodeAllocationManagerServiceCreateNodeAllocationHandler := connect.NewUnaryHandler(
		NodeAllocationManagerServiceCreateNodeAllocationProcedure,
		svc.CreateNodeAllocation,
		connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("CreateNodeAllocation")),
		connect.WithHandlerOptions(opts...),
	)
	nodeAllocationManagerServiceUpdateNodeAllocationHandler := connect.NewUnaryHandler(
		NodeAllocationManagerServiceUpdateNodeAllocationProcedure,
		svc.UpdateNodeAllocation,
		connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("UpdateNodeAllocation")),
		connect.WithHandlerOptions(opts...),
	)
	nodeAllocationManagerServiceDeleteNodeAllocationHandler := connect.NewUnaryHandler(
		NodeAllocationManagerServiceDeleteNodeAllocationProcedure,
		svc.DeleteNodeAllocation,
		connect.WithSchema(nodeAllocationManagerServiceMethods.ByName("DeleteNodeAllocation")),
		connect.WithHandlerOptions(opts...),
	)
	return "/backend_admin.NodeAllocationManagerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NodeAllocationManagerServiceGetNodeAllocationsProcedure:
			nodeAllocationManagerServiceGetNodeAllocationsHandler.ServeHTTP(w, r)
		case NodeAllocationManagerServiceGetNodeAllocationProcedure:
			nodeAllocationManagerServiceGetNodeAllocationHandler.ServeHTTP(w, r)
		case NodeAllocationManagerServiceCreateNodeAllocationProcedure:
			nodeAllocationManagerServiceCreateNodeAllocationHandler.ServeHTTP(w, r)
		case NodeAllocationManagerServiceUpdateNodeAllocationProcedure:
			nodeAllocationManagerServiceUpdateNodeAllocationHandler.ServeHTTP(w, r)
		case NodeAllocationManagerServiceDeleteNodeAllocationProcedure:
			nodeAllocationManagerServiceDeleteNodeAllocationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNodeAllocationManagerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNodeAllocationManagerServiceHandler struct{}

func (UnimplementedNodeAllocationManagerServiceHandler) GetNodeAllocations(context.Context, *connect.Request[admin.GetNodeAllocationsRequest]) (*connect.Response[admin.GetNodeAllocationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.NodeAllocationManagerService.GetNodeAllocations is not implemented"))
}

func (UnimplementedNodeAllocationManagerServiceHandler) GetNodeAllocation(context.Context, *connect.Request[admin.GetNodeAllocationRequest]) (*connect.Response[admin.GetNodeAllocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.NodeAllocationManagerService.GetNodeAllocation is not implemented"))
}

func (UnimplementedNodeAllocationManagerServiceHandler) CreateNodeAllocation(context.Context, *connect.Request[admin.CreateNodeAllocationRequest]) (*connect.Response[admin.CreateNodeAllocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.NodeAllocationManagerService.CreateNodeAllocation is not implemented"))
}

func (UnimplementedNodeAllocationManagerServiceHandler) UpdateNodeAllocation(context.Context, *connect.Request[admin.UpdateNodeAllocationRequest]) (*connect.Response[admin.UpdateNodeAllocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.NodeAllocationManagerService.UpdateNodeAllocation is not implemented"))
}

func (UnimplementedNodeAllocationManagerServiceHandler) DeleteNodeAllocation(context.Context, *connect.Request[admin.DeleteNodeAllocationRequest]) (*connect.Response[admin.DeleteNodeAllocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend_admin.NodeAllocationManagerService.DeleteNodeAllocation is not implemented"))
}
