// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: carving/session.proto

package carving

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogSessionStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogSessionStartRequest) Reset() {
	*x = LogSessionStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carving_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSessionStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSessionStartRequest) ProtoMessage() {}

func (x *LogSessionStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carving_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSessionStartRequest.ProtoReflect.Descriptor instead.
func (*LogSessionStartRequest) Descriptor() ([]byte, []int) {
	return file_carving_session_proto_rawDescGZIP(), []int{0}
}

type LogSessionStartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *LogSessionStartResponse) Reset() {
	*x = LogSessionStartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carving_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSessionStartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSessionStartResponse) ProtoMessage() {}

func (x *LogSessionStartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carving_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSessionStartResponse.ProtoReflect.Descriptor instead.
func (*LogSessionStartResponse) Descriptor() ([]byte, []int) {
	return file_carving_session_proto_rawDescGZIP(), []int{1}
}

func (x *LogSessionStartResponse) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

var File_carving_session_proto protoreflect.FileDescriptor

var file_carving_session_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x22, 0x18, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x17, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x68, 0x0a, 0x0e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0f,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x1f, 0x2e, 0x63, 0x61, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x75, 0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x61, 0x70, 0x70, 0x2f, 0x64,
	0x61, 0x79, 0x74, 0x72, 0x69, 0x70, 0x2d, 0x61, 0x6c, 0x61, 0x6e, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carving_session_proto_rawDescOnce sync.Once
	file_carving_session_proto_rawDescData = file_carving_session_proto_rawDesc
)

func file_carving_session_proto_rawDescGZIP() []byte {
	file_carving_session_proto_rawDescOnce.Do(func() {
		file_carving_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_carving_session_proto_rawDescData)
	})
	return file_carving_session_proto_rawDescData
}

var file_carving_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_carving_session_proto_goTypes = []interface{}{
	(*LogSessionStartRequest)(nil),  // 0: carving.LogSessionStartRequest
	(*LogSessionStartResponse)(nil), // 1: carving.LogSessionStartResponse
}
var file_carving_session_proto_depIdxs = []int32{
	0, // 0: carving.SessionService.LogSessionStart:input_type -> carving.LogSessionStartRequest
	1, // 1: carving.SessionService.LogSessionStart:output_type -> carving.LogSessionStartResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carving_session_proto_init() }
func file_carving_session_proto_init() {
	if File_carving_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carving_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSessionStartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carving_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSessionStartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_carving_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carving_session_proto_goTypes,
		DependencyIndexes: file_carving_session_proto_depIdxs,
		MessageInfos:      file_carving_session_proto_msgTypes,
	}.Build()
	File_carving_session_proto = out.File
	file_carving_session_proto_rawDesc = nil
	file_carving_session_proto_goTypes = nil
	file_carving_session_proto_depIdxs = nil
}
