// tbclient.SetUserBlack.SetUserBlackReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: SetUserBlackReqIdl.proto

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

type SetUserBlackReqIdl struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Data          *SetUserBlackReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetUserBlackReqIdl) Reset() {
	*x = SetUserBlackReqIdl{}
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetUserBlackReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserBlackReqIdl) ProtoMessage() {}

func (x *SetUserBlackReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserBlackReqIdl.ProtoReflect.Descriptor instead.
func (*SetUserBlackReqIdl) Descriptor() ([]byte, []int) {
	return file_SetUserBlackReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *SetUserBlackReqIdl) GetData() *SetUserBlackReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetUserBlackReqIdl_DataReq struct {
	state         protoimpl.MessageState                     `protogen:"open.v1"`
	Common        *CommonReq                                 `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	BlackUid      int64                                      `protobuf:"varint,2,opt,name=black_uid,json=blackUid,proto3" json:"black_uid,omitempty"`
	PermList      *SetUserBlackReqIdl_DataReq_PermissionList `protobuf:"bytes,3,opt,name=perm_list,json=permList,proto3" json:"perm_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetUserBlackReqIdl_DataReq) Reset() {
	*x = SetUserBlackReqIdl_DataReq{}
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetUserBlackReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserBlackReqIdl_DataReq) ProtoMessage() {}

func (x *SetUserBlackReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserBlackReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*SetUserBlackReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_SetUserBlackReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SetUserBlackReqIdl_DataReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *SetUserBlackReqIdl_DataReq) GetBlackUid() int64 {
	if x != nil {
		return x.BlackUid
	}
	return 0
}

func (x *SetUserBlackReqIdl_DataReq) GetPermList() *SetUserBlackReqIdl_DataReq_PermissionList {
	if x != nil {
		return x.PermList
	}
	return nil
}

type SetUserBlackReqIdl_DataReq_PermissionList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Follow        int32                  `protobuf:"varint,1,opt,name=follow,proto3" json:"follow,omitempty"`
	Interact      int32                  `protobuf:"varint,2,opt,name=interact,proto3" json:"interact,omitempty"`
	Chat          int32                  `protobuf:"varint,3,opt,name=chat,proto3" json:"chat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) Reset() {
	*x = SetUserBlackReqIdl_DataReq_PermissionList{}
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserBlackReqIdl_DataReq_PermissionList) ProtoMessage() {}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) ProtoReflect() protoreflect.Message {
	mi := &file_SetUserBlackReqIdl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserBlackReqIdl_DataReq_PermissionList.ProtoReflect.Descriptor instead.
func (*SetUserBlackReqIdl_DataReq_PermissionList) Descriptor() ([]byte, []int) {
	return file_SetUserBlackReqIdl_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) GetFollow() int32 {
	if x != nil {
		return x.Follow
	}
	return 0
}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) GetInteract() int32 {
	if x != nil {
		return x.Interact
	}
	return 0
}

func (x *SetUserBlackReqIdl_DataReq_PermissionList) GetChat() int32 {
	if x != nil {
		return x.Chat
	}
	return 0
}

var File_SetUserBlackReqIdl_proto protoreflect.FileDescriptor

const file_SetUserBlackReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x18SetUserBlackReqIdl.proto\x1a\x0fCommonReq.proto\"\xb5\x02\n" +
	"\x12SetUserBlackReqIdl\x12/\n" +
	"\x04data\x18\x01 \x01(\v2\x1b.SetUserBlackReqIdl.DataReqR\x04data\x1a\xed\x01\n" +
	"\aDataReq\x12\"\n" +
	"\x06common\x18\x01 \x01(\v2\n" +
	".CommonReqR\x06common\x12\x1b\n" +
	"\tblack_uid\x18\x02 \x01(\x03R\bblackUid\x12G\n" +
	"\tperm_list\x18\x03 \x01(\v2*.SetUserBlackReqIdl.DataReq.PermissionListR\bpermList\x1aX\n" +
	"\x0ePermissionList\x12\x16\n" +
	"\x06follow\x18\x01 \x01(\x05R\x06follow\x12\x1a\n" +
	"\binteract\x18\x02 \x01(\x05R\binteract\x12\x12\n" +
	"\x04chat\x18\x03 \x01(\x05R\x04chatB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_SetUserBlackReqIdl_proto_rawDescOnce sync.Once
	file_SetUserBlackReqIdl_proto_rawDescData []byte
)

func file_SetUserBlackReqIdl_proto_rawDescGZIP() []byte {
	file_SetUserBlackReqIdl_proto_rawDescOnce.Do(func() {
		file_SetUserBlackReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_SetUserBlackReqIdl_proto_rawDesc), len(file_SetUserBlackReqIdl_proto_rawDesc)))
	})
	return file_SetUserBlackReqIdl_proto_rawDescData
}

var file_SetUserBlackReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_SetUserBlackReqIdl_proto_goTypes = []any{
	(*SetUserBlackReqIdl)(nil),                        // 0: SetUserBlackReqIdl
	(*SetUserBlackReqIdl_DataReq)(nil),                // 1: SetUserBlackReqIdl.DataReq
	(*SetUserBlackReqIdl_DataReq_PermissionList)(nil), // 2: SetUserBlackReqIdl.DataReq.PermissionList
	(*CommonReq)(nil),                                 // 3: CommonReq
}
var file_SetUserBlackReqIdl_proto_depIdxs = []int32{
	1, // 0: SetUserBlackReqIdl.data:type_name -> SetUserBlackReqIdl.DataReq
	3, // 1: SetUserBlackReqIdl.DataReq.common:type_name -> CommonReq
	2, // 2: SetUserBlackReqIdl.DataReq.perm_list:type_name -> SetUserBlackReqIdl.DataReq.PermissionList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SetUserBlackReqIdl_proto_init() }
func file_SetUserBlackReqIdl_proto_init() {
	if File_SetUserBlackReqIdl_proto != nil {
		return
	}
	file_CommonReq_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_SetUserBlackReqIdl_proto_rawDesc), len(file_SetUserBlackReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetUserBlackReqIdl_proto_goTypes,
		DependencyIndexes: file_SetUserBlackReqIdl_proto_depIdxs,
		MessageInfos:      file_SetUserBlackReqIdl_proto_msgTypes,
	}.Build()
	File_SetUserBlackReqIdl_proto = out.File
	file_SetUserBlackReqIdl_proto_goTypes = nil
	file_SetUserBlackReqIdl_proto_depIdxs = nil
}
