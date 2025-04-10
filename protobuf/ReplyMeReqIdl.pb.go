// tbclient.ReplyMe.ReplyMeReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: ReplyMeReqIdl.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ReplyMeReqIdl struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *ReplyMeReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyMeReqIdl) Reset() {
	*x = ReplyMeReqIdl{}
	mi := &file_ReplyMeReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyMeReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMeReqIdl) ProtoMessage() {}

func (x *ReplyMeReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_ReplyMeReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMeReqIdl.ProtoReflect.Descriptor instead.
func (*ReplyMeReqIdl) Descriptor() ([]byte, []int) {
	return file_ReplyMeReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *ReplyMeReqIdl) GetData() *ReplyMeReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReplyMeReqIdl_DataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pn            string                 `protobuf:"bytes,1,opt,name=pn,proto3" json:"pn,omitempty"`
	Common        *CommonReq             `protobuf:"bytes,3,opt,name=common,proto3" json:"common,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyMeReqIdl_DataReq) Reset() {
	*x = ReplyMeReqIdl_DataReq{}
	mi := &file_ReplyMeReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyMeReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMeReqIdl_DataReq) ProtoMessage() {}

func (x *ReplyMeReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_ReplyMeReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMeReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*ReplyMeReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_ReplyMeReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReplyMeReqIdl_DataReq) GetPn() string {
	if x != nil {
		return x.Pn
	}
	return ""
}

func (x *ReplyMeReqIdl_DataReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

var File_ReplyMeReqIdl_proto protoreflect.FileDescriptor

const file_ReplyMeReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x13ReplyMeReqIdl.proto\x1a\x0fCommonReq.proto\"z\n" +
	"\rReplyMeReqIdl\x12*\n" +
	"\x04data\x18\x01 \x01(\v2\x16.ReplyMeReqIdl.DataReqR\x04data\x1a=\n" +
	"\aDataReq\x12\x0e\n" +
	"\x02pn\x18\x01 \x01(\tR\x02pn\x12\"\n" +
	"\x06common\x18\x03 \x01(\v2\n" +
	".CommonReqR\x06commonB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_ReplyMeReqIdl_proto_rawDescOnce sync.Once
	file_ReplyMeReqIdl_proto_rawDescData []byte
)

func file_ReplyMeReqIdl_proto_rawDescGZIP() []byte {
	file_ReplyMeReqIdl_proto_rawDescOnce.Do(func() {
		file_ReplyMeReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ReplyMeReqIdl_proto_rawDesc), len(file_ReplyMeReqIdl_proto_rawDesc)))
	})
	return file_ReplyMeReqIdl_proto_rawDescData
}

var file_ReplyMeReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ReplyMeReqIdl_proto_goTypes = []any{
	(*ReplyMeReqIdl)(nil),         // 0: ReplyMeReqIdl
	(*ReplyMeReqIdl_DataReq)(nil), // 1: ReplyMeReqIdl.DataReq
	(*CommonReq)(nil),             // 2: CommonReq
}
var file_ReplyMeReqIdl_proto_depIdxs = []int32{
	1, // 0: ReplyMeReqIdl.data:type_name -> ReplyMeReqIdl.DataReq
	2, // 1: ReplyMeReqIdl.DataReq.common:type_name -> CommonReq
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ReplyMeReqIdl_proto_init() }
func file_ReplyMeReqIdl_proto_init() {
	if File_ReplyMeReqIdl_proto != nil {
		return
	}
	file_CommonReq_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ReplyMeReqIdl_proto_rawDesc), len(file_ReplyMeReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReplyMeReqIdl_proto_goTypes,
		DependencyIndexes: file_ReplyMeReqIdl_proto_depIdxs,
		MessageInfos:      file_ReplyMeReqIdl_proto_msgTypes,
	}.Build()
	File_ReplyMeReqIdl_proto = out.File
	file_ReplyMeReqIdl_proto_goTypes = nil
	file_ReplyMeReqIdl_proto_depIdxs = nil
}
