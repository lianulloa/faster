// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: faster/fasterpb/faster.proto

package fasterpb

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

type InferNetworkSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InferNetworkSpeedRequest) Reset() {
	*x = InferNetworkSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faster_fasterpb_faster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferNetworkSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferNetworkSpeedRequest) ProtoMessage() {}

func (x *InferNetworkSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faster_fasterpb_faster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferNetworkSpeedRequest.ProtoReflect.Descriptor instead.
func (*InferNetworkSpeedRequest) Descriptor() ([]byte, []int) {
	return file_faster_fasterpb_faster_proto_rawDescGZIP(), []int{0}
}

func (x *InferNetworkSpeedRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InferNetworkSpeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed string `protobuf:"bytes,1,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *InferNetworkSpeedResponse) Reset() {
	*x = InferNetworkSpeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faster_fasterpb_faster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferNetworkSpeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferNetworkSpeedResponse) ProtoMessage() {}

func (x *InferNetworkSpeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faster_fasterpb_faster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferNetworkSpeedResponse.ProtoReflect.Descriptor instead.
func (*InferNetworkSpeedResponse) Descriptor() ([]byte, []int) {
	return file_faster_fasterpb_faster_proto_rawDescGZIP(), []int{1}
}

func (x *InferNetworkSpeedResponse) GetSpeed() string {
	if x != nil {
		return x.Speed
	}
	return ""
}

var File_faster_fasterpb_faster_proto protoreflect.FileDescriptor

var file_faster_fasterpb_faster_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x62, 0x2f, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x18, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x31, 0x0a, 0x19,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x32,
	0x69, 0x0a, 0x0d, 0x46, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x11, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x66, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x66,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_faster_fasterpb_faster_proto_rawDescOnce sync.Once
	file_faster_fasterpb_faster_proto_rawDescData = file_faster_fasterpb_faster_proto_rawDesc
)

func file_faster_fasterpb_faster_proto_rawDescGZIP() []byte {
	file_faster_fasterpb_faster_proto_rawDescOnce.Do(func() {
		file_faster_fasterpb_faster_proto_rawDescData = protoimpl.X.CompressGZIP(file_faster_fasterpb_faster_proto_rawDescData)
	})
	return file_faster_fasterpb_faster_proto_rawDescData
}

var file_faster_fasterpb_faster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_faster_fasterpb_faster_proto_goTypes = []interface{}{
	(*InferNetworkSpeedRequest)(nil),  // 0: faster.InferNetworkSpeedRequest
	(*InferNetworkSpeedResponse)(nil), // 1: faster.InferNetworkSpeedResponse
}
var file_faster_fasterpb_faster_proto_depIdxs = []int32{
	0, // 0: faster.FasterService.InferNetworkSpeed:input_type -> faster.InferNetworkSpeedRequest
	1, // 1: faster.FasterService.InferNetworkSpeed:output_type -> faster.InferNetworkSpeedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_faster_fasterpb_faster_proto_init() }
func file_faster_fasterpb_faster_proto_init() {
	if File_faster_fasterpb_faster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faster_fasterpb_faster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferNetworkSpeedRequest); i {
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
		file_faster_fasterpb_faster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferNetworkSpeedResponse); i {
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
			RawDescriptor: file_faster_fasterpb_faster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_faster_fasterpb_faster_proto_goTypes,
		DependencyIndexes: file_faster_fasterpb_faster_proto_depIdxs,
		MessageInfos:      file_faster_fasterpb_faster_proto_msgTypes,
	}.Build()
	File_faster_fasterpb_faster_proto = out.File
	file_faster_fasterpb_faster_proto_rawDesc = nil
	file_faster_fasterpb_faster_proto_goTypes = nil
	file_faster_fasterpb_faster_proto_depIdxs = nil
}
