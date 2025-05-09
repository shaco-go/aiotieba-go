// tbclient.GetDislikeList.GetDislikeListReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: GetDislikeListReqIdl.proto

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

type GetDislikeListReqIdl struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Data          *GetDislikeListReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDislikeListReqIdl) Reset() {
	*x = GetDislikeListReqIdl{}
	mi := &file_GetDislikeListReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDislikeListReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDislikeListReqIdl) ProtoMessage() {}

func (x *GetDislikeListReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_GetDislikeListReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDislikeListReqIdl.ProtoReflect.Descriptor instead.
func (*GetDislikeListReqIdl) Descriptor() ([]byte, []int) {
	return file_GetDislikeListReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *GetDislikeListReqIdl) GetData() *GetDislikeListReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetDislikeListReqIdl_DataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Common        *CommonReq             `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Pn            int32                  `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	Rn            int32                  `protobuf:"varint,4,opt,name=rn,proto3" json:"rn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDislikeListReqIdl_DataReq) Reset() {
	*x = GetDislikeListReqIdl_DataReq{}
	mi := &file_GetDislikeListReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDislikeListReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDislikeListReqIdl_DataReq) ProtoMessage() {}

func (x *GetDislikeListReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetDislikeListReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDislikeListReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*GetDislikeListReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_GetDislikeListReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetDislikeListReqIdl_DataReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *GetDislikeListReqIdl_DataReq) GetPn() int32 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *GetDislikeListReqIdl_DataReq) GetRn() int32 {
	if x != nil {
		return x.Rn
	}
	return 0
}

var File_GetDislikeListReqIdl_proto protoreflect.FileDescriptor

const file_GetDislikeListReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x1aGetDislikeListReqIdl.proto\x1a\x0fCommonReq.proto\"\x98\x01\n" +
	"\x14GetDislikeListReqIdl\x121\n" +
	"\x04data\x18\x01 \x01(\v2\x1d.GetDislikeListReqIdl.DataReqR\x04data\x1aM\n" +
	"\aDataReq\x12\"\n" +
	"\x06common\x18\x01 \x01(\v2\n" +
	".CommonReqR\x06common\x12\x0e\n" +
	"\x02pn\x18\x03 \x01(\x05R\x02pn\x12\x0e\n" +
	"\x02rn\x18\x04 \x01(\x05R\x02rnB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_GetDislikeListReqIdl_proto_rawDescOnce sync.Once
	file_GetDislikeListReqIdl_proto_rawDescData []byte
)

func file_GetDislikeListReqIdl_proto_rawDescGZIP() []byte {
	file_GetDislikeListReqIdl_proto_rawDescOnce.Do(func() {
		file_GetDislikeListReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_GetDislikeListReqIdl_proto_rawDesc), len(file_GetDislikeListReqIdl_proto_rawDesc)))
	})
	return file_GetDislikeListReqIdl_proto_rawDescData
}

var file_GetDislikeListReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetDislikeListReqIdl_proto_goTypes = []any{
	(*GetDislikeListReqIdl)(nil),         // 0: GetDislikeListReqIdl
	(*GetDislikeListReqIdl_DataReq)(nil), // 1: GetDislikeListReqIdl.DataReq
	(*CommonReq)(nil),                    // 2: CommonReq
}
var file_GetDislikeListReqIdl_proto_depIdxs = []int32{
	1, // 0: GetDislikeListReqIdl.data:type_name -> GetDislikeListReqIdl.DataReq
	2, // 1: GetDislikeListReqIdl.DataReq.common:type_name -> CommonReq
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetDislikeListReqIdl_proto_init() }
func file_GetDislikeListReqIdl_proto_init() {
	if File_GetDislikeListReqIdl_proto != nil {
		return
	}
	file_CommonReq_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_GetDislikeListReqIdl_proto_rawDesc), len(file_GetDislikeListReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetDislikeListReqIdl_proto_goTypes,
		DependencyIndexes: file_GetDislikeListReqIdl_proto_depIdxs,
		MessageInfos:      file_GetDislikeListReqIdl_proto_msgTypes,
	}.Build()
	File_GetDislikeListReqIdl_proto = out.File
	file_GetDislikeListReqIdl_proto_goTypes = nil
	file_GetDislikeListReqIdl_proto_depIdxs = nil
}
