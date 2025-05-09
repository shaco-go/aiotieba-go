// tbclient.UserMuteQuery.UserMuteQueryResIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: UserMuteQueryResIdl.proto

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

type UserMuteQueryResIdl struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Data          *UserMuteQueryResIdl_DataRes `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error         *Error                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserMuteQueryResIdl) Reset() {
	*x = UserMuteQueryResIdl{}
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMuteQueryResIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMuteQueryResIdl) ProtoMessage() {}

func (x *UserMuteQueryResIdl) ProtoReflect() protoreflect.Message {
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMuteQueryResIdl.ProtoReflect.Descriptor instead.
func (*UserMuteQueryResIdl) Descriptor() ([]byte, []int) {
	return file_UserMuteQueryResIdl_proto_rawDescGZIP(), []int{0}
}

func (x *UserMuteQueryResIdl) GetData() *UserMuteQueryResIdl_DataRes {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UserMuteQueryResIdl) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UserMuteQueryResIdl_DataRes struct {
	state         protoimpl.MessageState                  `protogen:"open.v1"`
	MuteUser      []*UserMuteQueryResIdl_DataRes_MuteUser `protobuf:"bytes,1,rep,name=mute_user,json=muteUser,proto3" json:"mute_user,omitempty"`
	Page          *Page                                   `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserMuteQueryResIdl_DataRes) Reset() {
	*x = UserMuteQueryResIdl_DataRes{}
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMuteQueryResIdl_DataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMuteQueryResIdl_DataRes) ProtoMessage() {}

func (x *UserMuteQueryResIdl_DataRes) ProtoReflect() protoreflect.Message {
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMuteQueryResIdl_DataRes.ProtoReflect.Descriptor instead.
func (*UserMuteQueryResIdl_DataRes) Descriptor() ([]byte, []int) {
	return file_UserMuteQueryResIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UserMuteQueryResIdl_DataRes) GetMuteUser() []*UserMuteQueryResIdl_DataRes_MuteUser {
	if x != nil {
		return x.MuteUser
	}
	return nil
}

func (x *UserMuteQueryResIdl_DataRes) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type UserMuteQueryResIdl_DataRes_MuteUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	MuteTime      int32                  `protobuf:"varint,3,opt,name=mute_time,json=muteTime,proto3" json:"mute_time,omitempty"`
	Portrait      string                 `protobuf:"bytes,4,opt,name=portrait,proto3" json:"portrait,omitempty"`
	NameShow      string                 `protobuf:"bytes,5,opt,name=name_show,json=nameShow,proto3" json:"name_show,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) Reset() {
	*x = UserMuteQueryResIdl_DataRes_MuteUser{}
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMuteQueryResIdl_DataRes_MuteUser) ProtoMessage() {}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) ProtoReflect() protoreflect.Message {
	mi := &file_UserMuteQueryResIdl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMuteQueryResIdl_DataRes_MuteUser.ProtoReflect.Descriptor instead.
func (*UserMuteQueryResIdl_DataRes_MuteUser) Descriptor() ([]byte, []int) {
	return file_UserMuteQueryResIdl_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) GetMuteTime() int32 {
	if x != nil {
		return x.MuteTime
	}
	return 0
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) GetPortrait() string {
	if x != nil {
		return x.Portrait
	}
	return ""
}

func (x *UserMuteQueryResIdl_DataRes_MuteUser) GetNameShow() string {
	if x != nil {
		return x.NameShow
	}
	return ""
}

var File_UserMuteQueryResIdl_proto protoreflect.FileDescriptor

const file_UserMuteQueryResIdl_proto_rawDesc = "" +
	"\n" +
	"\x19UserMuteQueryResIdl.proto\x1a\vError.proto\x1a\n" +
	"Page.proto\"\xe9\x02\n" +
	"\x13UserMuteQueryResIdl\x120\n" +
	"\x04data\x18\x01 \x01(\v2\x1c.UserMuteQueryResIdl.DataResR\x04data\x12\x1c\n" +
	"\x05error\x18\x02 \x01(\v2\x06.ErrorR\x05error\x1a\x81\x02\n" +
	"\aDataRes\x12B\n" +
	"\tmute_user\x18\x01 \x03(\v2%.UserMuteQueryResIdl.DataRes.MuteUserR\bmuteUser\x12\x19\n" +
	"\x04page\x18\x02 \x01(\v2\x05.PageR\x04page\x1a\x96\x01\n" +
	"\bMuteUser\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x1b\n" +
	"\tmute_time\x18\x03 \x01(\x05R\bmuteTime\x12\x1a\n" +
	"\bportrait\x18\x04 \x01(\tR\bportrait\x12\x1b\n" +
	"\tname_show\x18\x05 \x01(\tR\bnameShowB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_UserMuteQueryResIdl_proto_rawDescOnce sync.Once
	file_UserMuteQueryResIdl_proto_rawDescData []byte
)

func file_UserMuteQueryResIdl_proto_rawDescGZIP() []byte {
	file_UserMuteQueryResIdl_proto_rawDescOnce.Do(func() {
		file_UserMuteQueryResIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_UserMuteQueryResIdl_proto_rawDesc), len(file_UserMuteQueryResIdl_proto_rawDesc)))
	})
	return file_UserMuteQueryResIdl_proto_rawDescData
}

var file_UserMuteQueryResIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_UserMuteQueryResIdl_proto_goTypes = []any{
	(*UserMuteQueryResIdl)(nil),                  // 0: UserMuteQueryResIdl
	(*UserMuteQueryResIdl_DataRes)(nil),          // 1: UserMuteQueryResIdl.DataRes
	(*UserMuteQueryResIdl_DataRes_MuteUser)(nil), // 2: UserMuteQueryResIdl.DataRes.MuteUser
	(*Error)(nil), // 3: Error
	(*Page)(nil),  // 4: Page
}
var file_UserMuteQueryResIdl_proto_depIdxs = []int32{
	1, // 0: UserMuteQueryResIdl.data:type_name -> UserMuteQueryResIdl.DataRes
	3, // 1: UserMuteQueryResIdl.error:type_name -> Error
	2, // 2: UserMuteQueryResIdl.DataRes.mute_user:type_name -> UserMuteQueryResIdl.DataRes.MuteUser
	4, // 3: UserMuteQueryResIdl.DataRes.page:type_name -> Page
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_UserMuteQueryResIdl_proto_init() }
func file_UserMuteQueryResIdl_proto_init() {
	if File_UserMuteQueryResIdl_proto != nil {
		return
	}
	file_Error_proto_init()
	file_Page_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_UserMuteQueryResIdl_proto_rawDesc), len(file_UserMuteQueryResIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UserMuteQueryResIdl_proto_goTypes,
		DependencyIndexes: file_UserMuteQueryResIdl_proto_depIdxs,
		MessageInfos:      file_UserMuteQueryResIdl_proto_msgTypes,
	}.Build()
	File_UserMuteQueryResIdl_proto = out.File
	file_UserMuteQueryResIdl_proto_goTypes = nil
	file_UserMuteQueryResIdl_proto_depIdxs = nil
}
