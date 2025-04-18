// tbclient.GetForumDetail.GetForumDetailReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: GetForumDetailReqIdl.proto

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

type GetForumDetailReqIdl struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Data          *GetForumDetailReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetForumDetailReqIdl) Reset() {
	*x = GetForumDetailReqIdl{}
	mi := &file_GetForumDetailReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailReqIdl) ProtoMessage() {}

func (x *GetForumDetailReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailReqIdl.ProtoReflect.Descriptor instead.
func (*GetForumDetailReqIdl) Descriptor() ([]byte, []int) {
	return file_GetForumDetailReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *GetForumDetailReqIdl) GetData() *GetForumDetailReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetForumDetailReqIdl_DataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ForumId       int64                  `protobuf:"varint,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	Common        *CommonReq             `protobuf:"bytes,2,opt,name=common,proto3" json:"common,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetForumDetailReqIdl_DataReq) Reset() {
	*x = GetForumDetailReqIdl_DataReq{}
	mi := &file_GetForumDetailReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailReqIdl_DataReq) ProtoMessage() {}

func (x *GetForumDetailReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*GetForumDetailReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_GetForumDetailReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetForumDetailReqIdl_DataReq) GetForumId() int64 {
	if x != nil {
		return x.ForumId
	}
	return 0
}

func (x *GetForumDetailReqIdl_DataReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

var File_GetForumDetailReqIdl_proto protoreflect.FileDescriptor

const file_GetForumDetailReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x1aGetForumDetailReqIdl.proto\x1a\x0fCommonReq.proto\"\x93\x01\n" +
	"\x14GetForumDetailReqIdl\x121\n" +
	"\x04data\x18\x01 \x01(\v2\x1d.GetForumDetailReqIdl.DataReqR\x04data\x1aH\n" +
	"\aDataReq\x12\x19\n" +
	"\bforum_id\x18\x01 \x01(\x03R\aforumId\x12\"\n" +
	"\x06common\x18\x02 \x01(\v2\n" +
	".CommonReqR\x06commonB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_GetForumDetailReqIdl_proto_rawDescOnce sync.Once
	file_GetForumDetailReqIdl_proto_rawDescData []byte
)

func file_GetForumDetailReqIdl_proto_rawDescGZIP() []byte {
	file_GetForumDetailReqIdl_proto_rawDescOnce.Do(func() {
		file_GetForumDetailReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_GetForumDetailReqIdl_proto_rawDesc), len(file_GetForumDetailReqIdl_proto_rawDesc)))
	})
	return file_GetForumDetailReqIdl_proto_rawDescData
}

var file_GetForumDetailReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetForumDetailReqIdl_proto_goTypes = []any{
	(*GetForumDetailReqIdl)(nil),         // 0: GetForumDetailReqIdl
	(*GetForumDetailReqIdl_DataReq)(nil), // 1: GetForumDetailReqIdl.DataReq
	(*CommonReq)(nil),                    // 2: CommonReq
}
var file_GetForumDetailReqIdl_proto_depIdxs = []int32{
	1, // 0: GetForumDetailReqIdl.data:type_name -> GetForumDetailReqIdl.DataReq
	2, // 1: GetForumDetailReqIdl.DataReq.common:type_name -> CommonReq
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetForumDetailReqIdl_proto_init() }
func file_GetForumDetailReqIdl_proto_init() {
	if File_GetForumDetailReqIdl_proto != nil {
		return
	}
	file_CommonReq_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_GetForumDetailReqIdl_proto_rawDesc), len(file_GetForumDetailReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetForumDetailReqIdl_proto_goTypes,
		DependencyIndexes: file_GetForumDetailReqIdl_proto_depIdxs,
		MessageInfos:      file_GetForumDetailReqIdl_proto_msgTypes,
	}.Build()
	File_GetForumDetailReqIdl_proto = out.File
	file_GetForumDetailReqIdl_proto_goTypes = nil
	file_GetForumDetailReqIdl_proto_depIdxs = nil
}
