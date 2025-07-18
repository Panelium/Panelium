// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: daemon/Backend.proto

package daemon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto_gen_go "panelium/proto_gen_go"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Server struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Sid           string                       `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	OwnerId       string                       `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	UserIds       []string                     `protobuf:"bytes,3,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Allocations   []*proto_gen_go.IPAllocation `protobuf:"bytes,4,rep,name=allocations,proto3" json:"allocations,omitempty"`
	ResourceLimit *proto_gen_go.ResourceLimit  `protobuf:"bytes,5,opt,name=resource_limit,json=resourceLimit,proto3" json:"resource_limit,omitempty"`
	DockerImage   string                       `protobuf:"bytes,6,opt,name=docker_image,json=dockerImage,proto3" json:"docker_image,omitempty"`
	Bid           string                       `protobuf:"bytes,7,opt,name=bid,proto3" json:"bid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_daemon_Backend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_Backend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_daemon_Backend_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Server) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Server) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *Server) GetAllocations() []*proto_gen_go.IPAllocation {
	if x != nil {
		return x.Allocations
	}
	return nil
}

func (x *Server) GetResourceLimit() *proto_gen_go.ResourceLimit {
	if x != nil {
		return x.ResourceLimit
	}
	return nil
}

func (x *Server) GetDockerImage() string {
	if x != nil {
		return x.DockerImage
	}
	return ""
}

func (x *Server) GetBid() string {
	if x != nil {
		return x.Bid
	}
	return ""
}

var File_daemon_Backend_proto protoreflect.FileDescriptor

const file_daemon_Backend_proto_rawDesc = "" +
	"\n" +
	"\x14daemon/Backend.proto\x12\x06daemon\x1a\fcommon.proto\"\xfb\x01\n" +
	"\x06Server\x12\x10\n" +
	"\x03sid\x18\x01 \x01(\tR\x03sid\x12\x19\n" +
	"\bowner_id\x18\x02 \x01(\tR\aownerId\x12\x19\n" +
	"\buser_ids\x18\x03 \x03(\tR\auserIds\x126\n" +
	"\vallocations\x18\x04 \x03(\v2\x14.common.IPAllocationR\vallocations\x12<\n" +
	"\x0eresource_limit\x18\x05 \x01(\v2\x15.common.ResourceLimitR\rresourceLimit\x12!\n" +
	"\fdocker_image\x18\x06 \x01(\tR\vdockerImage\x12\x10\n" +
	"\x03bid\x18\a \x01(\tR\x03bid2\xc1\x01\n" +
	"\x0eBackendService\x126\n" +
	"\fCreateServer\x12\x0e.daemon.Server\x1a\x16.common.SuccessMessage\x126\n" +
	"\fUpdateServer\x12\x0e.daemon.Server\x1a\x16.common.SuccessMessage\x12?\n" +
	"\fDeleteServer\x12\x17.common.SimpleIDMessage\x1a\x16.common.SuccessMessageB\x1eZ\x1cpanelium/proto_gen_go/daemonb\x06proto3"

var (
	file_daemon_Backend_proto_rawDescOnce sync.Once
	file_daemon_Backend_proto_rawDescData []byte
)

func file_daemon_Backend_proto_rawDescGZIP() []byte {
	file_daemon_Backend_proto_rawDescOnce.Do(func() {
		file_daemon_Backend_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_daemon_Backend_proto_rawDesc), len(file_daemon_Backend_proto_rawDesc)))
	})
	return file_daemon_Backend_proto_rawDescData
}

var file_daemon_Backend_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_daemon_Backend_proto_goTypes = []any{
	(*Server)(nil),                       // 0: daemon.Server
	(*proto_gen_go.IPAllocation)(nil),    // 1: common.IPAllocation
	(*proto_gen_go.ResourceLimit)(nil),   // 2: common.ResourceLimit
	(*proto_gen_go.SimpleIDMessage)(nil), // 3: common.SimpleIDMessage
	(*proto_gen_go.SuccessMessage)(nil),  // 4: common.SuccessMessage
}
var file_daemon_Backend_proto_depIdxs = []int32{
	1, // 0: daemon.Server.allocations:type_name -> common.IPAllocation
	2, // 1: daemon.Server.resource_limit:type_name -> common.ResourceLimit
	0, // 2: daemon.BackendService.CreateServer:input_type -> daemon.Server
	0, // 3: daemon.BackendService.UpdateServer:input_type -> daemon.Server
	3, // 4: daemon.BackendService.DeleteServer:input_type -> common.SimpleIDMessage
	4, // 5: daemon.BackendService.CreateServer:output_type -> common.SuccessMessage
	4, // 6: daemon.BackendService.UpdateServer:output_type -> common.SuccessMessage
	4, // 7: daemon.BackendService.DeleteServer:output_type -> common.SuccessMessage
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_daemon_Backend_proto_init() }
func file_daemon_Backend_proto_init() {
	if File_daemon_Backend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_daemon_Backend_proto_rawDesc), len(file_daemon_Backend_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daemon_Backend_proto_goTypes,
		DependencyIndexes: file_daemon_Backend_proto_depIdxs,
		MessageInfos:      file_daemon_Backend_proto_msgTypes,
	}.Build()
	File_daemon_Backend_proto = out.File
	file_daemon_Backend_proto_goTypes = nil
	file_daemon_Backend_proto_depIdxs = nil
}
