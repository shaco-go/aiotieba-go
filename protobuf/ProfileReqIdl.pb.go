// tbclient.Profile.ProfileReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: ProfileReqIdl.proto

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

type ProfileReqIdl struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *ProfileReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileReqIdl) Reset() {
	*x = ProfileReqIdl{}
	mi := &file_ProfileReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileReqIdl) ProtoMessage() {}

func (x *ProfileReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileReqIdl.ProtoReflect.Descriptor instead.
func (*ProfileReqIdl) Descriptor() ([]byte, []int) {
	return file_ProfileReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileReqIdl) GetData() *ProfileReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProfileReqIdl_DataReq struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Uid               int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	NeedPostCount     uint32                 `protobuf:"varint,2,opt,name=need_post_count,json=needPostCount,proto3" json:"need_post_count,omitempty"`
	Pn                uint32                 `protobuf:"varint,6,opt,name=pn,proto3" json:"pn,omitempty"`
	Common            *CommonReq             `protobuf:"bytes,9,opt,name=common,proto3" json:"common,omitempty"`
	Page              int32                  `protobuf:"varint,15,opt,name=page,proto3" json:"page,omitempty"`
	FriendUidPortrait string                 `protobuf:"bytes,16,opt,name=friend_uid_portrait,json=friendUidPortrait,proto3" json:"friend_uid_portrait,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ProfileReqIdl_DataReq) Reset() {
	*x = ProfileReqIdl_DataReq{}
	mi := &file_ProfileReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileReqIdl_DataReq) ProtoMessage() {}

func (x *ProfileReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*ProfileReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_ProfileReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProfileReqIdl_DataReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ProfileReqIdl_DataReq) GetNeedPostCount() uint32 {
	if x != nil {
		return x.NeedPostCount
	}
	return 0
}

func (x *ProfileReqIdl_DataReq) GetPn() uint32 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *ProfileReqIdl_DataReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ProfileReqIdl_DataReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProfileReqIdl_DataReq) GetFriendUidPortrait() string {
	if x != nil {
		return x.FriendUidPortrait
	}
	return ""
}

var File_ProfileReqIdl_proto protoreflect.FileDescriptor

const file_ProfileReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x13ProfileReqIdl.proto\x1a\x0fCommonReq.proto\"\xf9\x01\n" +
	"\rProfileReqIdl\x12*\n" +
	"\x04data\x18\x01 \x01(\v2\x16.ProfileReqIdl.DataReqR\x04data\x1a\xbb\x01\n" +
	"\aDataReq\x12\x10\n" +
	"\x03uid\x18\x01 \x01(\x03R\x03uid\x12&\n" +
	"\x0fneed_post_count\x18\x02 \x01(\rR\rneedPostCount\x12\x0e\n" +
	"\x02pn\x18\x06 \x01(\rR\x02pn\x12\"\n" +
	"\x06common\x18\t \x01(\v2\n" +
	".CommonReqR\x06common\x12\x12\n" +
	"\x04page\x18\x0f \x01(\x05R\x04page\x12.\n" +
	"\x13friend_uid_portrait\x18\x10 \x01(\tR\x11friendUidPortraitB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_ProfileReqIdl_proto_rawDescOnce sync.Once
	file_ProfileReqIdl_proto_rawDescData []byte
)

func file_ProfileReqIdl_proto_rawDescGZIP() []byte {
	file_ProfileReqIdl_proto_rawDescOnce.Do(func() {
		file_ProfileReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ProfileReqIdl_proto_rawDesc), len(file_ProfileReqIdl_proto_rawDesc)))
	})
	return file_ProfileReqIdl_proto_rawDescData
}

var file_ProfileReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProfileReqIdl_proto_goTypes = []any{
	(*ProfileReqIdl)(nil),         // 0: ProfileReqIdl
	(*ProfileReqIdl_DataReq)(nil), // 1: ProfileReqIdl.DataReq
	(*CommonReq)(nil),             // 2: CommonReq
}
var file_ProfileReqIdl_proto_depIdxs = []int32{
	1, // 0: ProfileReqIdl.data:type_name -> ProfileReqIdl.DataReq
	2, // 1: ProfileReqIdl.DataReq.common:type_name -> CommonReq
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ProfileReqIdl_proto_init() }
func file_ProfileReqIdl_proto_init() {
	if File_ProfileReqIdl_proto != nil {
		return
	}
	file_CommonReq_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ProfileReqIdl_proto_rawDesc), len(file_ProfileReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProfileReqIdl_proto_goTypes,
		DependencyIndexes: file_ProfileReqIdl_proto_depIdxs,
		MessageInfos:      file_ProfileReqIdl_proto_msgTypes,
	}.Build()
	File_ProfileReqIdl_proto = out.File
	file_ProfileReqIdl_proto_goTypes = nil
	file_ProfileReqIdl_proto_depIdxs = nil
}
