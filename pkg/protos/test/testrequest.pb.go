// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: testrequest.proto

package test

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

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testrequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testrequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_testrequest_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_testrequest_proto protoreflect.FileDescriptor

var file_testrequest_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0b, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65, 0x65,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x74, 0x61, 0x79, 0x61, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testrequest_proto_rawDescOnce sync.Once
	file_testrequest_proto_rawDescData = file_testrequest_proto_rawDesc
)

func file_testrequest_proto_rawDescGZIP() []byte {
	file_testrequest_proto_rawDescOnce.Do(func() {
		file_testrequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_testrequest_proto_rawDescData)
	})
	return file_testrequest_proto_rawDescData
}

var file_testrequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testrequest_proto_goTypes = []interface{}{
	(*TestRequest)(nil), // 0: test.TestRequest
}
var file_testrequest_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testrequest_proto_init() }
func file_testrequest_proto_init() {
	if File_testrequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testrequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
			RawDescriptor: file_testrequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testrequest_proto_goTypes,
		DependencyIndexes: file_testrequest_proto_depIdxs,
		MessageInfos:      file_testrequest_proto_msgTypes,
	}.Build()
	File_testrequest_proto = out.File
	file_testrequest_proto_rawDesc = nil
	file_testrequest_proto_goTypes = nil
	file_testrequest_proto_depIdxs = nil
}
