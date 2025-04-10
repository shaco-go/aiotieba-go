// tbclient.Profile.ProfileResIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: ProfileResIdl.proto

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

type ProfileResIdl struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *Error                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data          *ProfileResIdl_DataRes `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResIdl) Reset() {
	*x = ProfileResIdl{}
	mi := &file_ProfileResIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResIdl) ProtoMessage() {}

func (x *ProfileResIdl) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileResIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResIdl.ProtoReflect.Descriptor instead.
func (*ProfileResIdl) Descriptor() ([]byte, []int) {
	return file_ProfileResIdl_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileResIdl) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ProfileResIdl) GetData() *ProfileResIdl_DataRes {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProfileResIdl_DataRes struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	User          *User                                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AntiStat      *ProfileResIdl_DataRes_Anti          `protobuf:"bytes,2,opt,name=anti_stat,json=antiStat,proto3" json:"anti_stat,omitempty"`
	PostList      []*PostInfoList                      `protobuf:"bytes,4,rep,name=post_list,json=postList,proto3" json:"post_list,omitempty"`
	UserAgreeInfo *ProfileResIdl_DataRes_UserAgreeInfo `protobuf:"bytes,14,opt,name=user_agree_info,json=userAgreeInfo,proto3" json:"user_agree_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResIdl_DataRes) Reset() {
	*x = ProfileResIdl_DataRes{}
	mi := &file_ProfileResIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResIdl_DataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResIdl_DataRes) ProtoMessage() {}

func (x *ProfileResIdl_DataRes) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileResIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResIdl_DataRes.ProtoReflect.Descriptor instead.
func (*ProfileResIdl_DataRes) Descriptor() ([]byte, []int) {
	return file_ProfileResIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProfileResIdl_DataRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ProfileResIdl_DataRes) GetAntiStat() *ProfileResIdl_DataRes_Anti {
	if x != nil {
		return x.AntiStat
	}
	return nil
}

func (x *ProfileResIdl_DataRes) GetPostList() []*PostInfoList {
	if x != nil {
		return x.PostList
	}
	return nil
}

func (x *ProfileResIdl_DataRes) GetUserAgreeInfo() *ProfileResIdl_DataRes_UserAgreeInfo {
	if x != nil {
		return x.UserAgreeInfo
	}
	return nil
}

type ProfileResIdl_DataRes_Anti struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockStat     int32                  `protobuf:"varint,6,opt,name=block_stat,json=blockStat,proto3" json:"block_stat,omitempty"`
	HideStat      int32                  `protobuf:"varint,7,opt,name=hide_stat,json=hideStat,proto3" json:"hide_stat,omitempty"`
	DaysTofree    int32                  `protobuf:"varint,9,opt,name=days_tofree,json=daysTofree,proto3" json:"days_tofree,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResIdl_DataRes_Anti) Reset() {
	*x = ProfileResIdl_DataRes_Anti{}
	mi := &file_ProfileResIdl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResIdl_DataRes_Anti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResIdl_DataRes_Anti) ProtoMessage() {}

func (x *ProfileResIdl_DataRes_Anti) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileResIdl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResIdl_DataRes_Anti.ProtoReflect.Descriptor instead.
func (*ProfileResIdl_DataRes_Anti) Descriptor() ([]byte, []int) {
	return file_ProfileResIdl_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ProfileResIdl_DataRes_Anti) GetBlockStat() int32 {
	if x != nil {
		return x.BlockStat
	}
	return 0
}

func (x *ProfileResIdl_DataRes_Anti) GetHideStat() int32 {
	if x != nil {
		return x.HideStat
	}
	return 0
}

func (x *ProfileResIdl_DataRes_Anti) GetDaysTofree() int32 {
	if x != nil {
		return x.DaysTofree
	}
	return 0
}

type ProfileResIdl_DataRes_UserAgreeInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalAgreeNum int64                  `protobuf:"varint,1,opt,name=total_agree_num,json=totalAgreeNum,proto3" json:"total_agree_num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResIdl_DataRes_UserAgreeInfo) Reset() {
	*x = ProfileResIdl_DataRes_UserAgreeInfo{}
	mi := &file_ProfileResIdl_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResIdl_DataRes_UserAgreeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResIdl_DataRes_UserAgreeInfo) ProtoMessage() {}

func (x *ProfileResIdl_DataRes_UserAgreeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ProfileResIdl_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResIdl_DataRes_UserAgreeInfo.ProtoReflect.Descriptor instead.
func (*ProfileResIdl_DataRes_UserAgreeInfo) Descriptor() ([]byte, []int) {
	return file_ProfileResIdl_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ProfileResIdl_DataRes_UserAgreeInfo) GetTotalAgreeNum() int64 {
	if x != nil {
		return x.TotalAgreeNum
	}
	return 0
}

var File_ProfileResIdl_proto protoreflect.FileDescriptor

const file_ProfileResIdl_proto_rawDesc = "" +
	"\n" +
	"\x13ProfileResIdl.proto\x1a\vError.proto\x1a\n" +
	"User.proto\x1a\x12PostInfoList.proto\"\xd2\x03\n" +
	"\rProfileResIdl\x12\x1c\n" +
	"\x05error\x18\x01 \x01(\v2\x06.ErrorR\x05error\x12*\n" +
	"\x04data\x18\x02 \x01(\v2\x16.ProfileResIdl.DataResR\x04data\x1a\xf6\x02\n" +
	"\aDataRes\x12\x19\n" +
	"\x04user\x18\x01 \x01(\v2\x05.UserR\x04user\x128\n" +
	"\tanti_stat\x18\x02 \x01(\v2\x1b.ProfileResIdl.DataRes.AntiR\bantiStat\x12*\n" +
	"\tpost_list\x18\x04 \x03(\v2\r.PostInfoListR\bpostList\x12L\n" +
	"\x0fuser_agree_info\x18\x0e \x01(\v2$.ProfileResIdl.DataRes.UserAgreeInfoR\ruserAgreeInfo\x1ac\n" +
	"\x04Anti\x12\x1d\n" +
	"\n" +
	"block_stat\x18\x06 \x01(\x05R\tblockStat\x12\x1b\n" +
	"\thide_stat\x18\a \x01(\x05R\bhideStat\x12\x1f\n" +
	"\vdays_tofree\x18\t \x01(\x05R\n" +
	"daysTofree\x1a7\n" +
	"\rUserAgreeInfo\x12&\n" +
	"\x0ftotal_agree_num\x18\x01 \x01(\x03R\rtotalAgreeNumB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_ProfileResIdl_proto_rawDescOnce sync.Once
	file_ProfileResIdl_proto_rawDescData []byte
)

func file_ProfileResIdl_proto_rawDescGZIP() []byte {
	file_ProfileResIdl_proto_rawDescOnce.Do(func() {
		file_ProfileResIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ProfileResIdl_proto_rawDesc), len(file_ProfileResIdl_proto_rawDesc)))
	})
	return file_ProfileResIdl_proto_rawDescData
}

var file_ProfileResIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ProfileResIdl_proto_goTypes = []any{
	(*ProfileResIdl)(nil),                       // 0: ProfileResIdl
	(*ProfileResIdl_DataRes)(nil),               // 1: ProfileResIdl.DataRes
	(*ProfileResIdl_DataRes_Anti)(nil),          // 2: ProfileResIdl.DataRes.Anti
	(*ProfileResIdl_DataRes_UserAgreeInfo)(nil), // 3: ProfileResIdl.DataRes.UserAgreeInfo
	(*Error)(nil),                               // 4: Error
	(*User)(nil),                                // 5: User
	(*PostInfoList)(nil),                        // 6: PostInfoList
}
var file_ProfileResIdl_proto_depIdxs = []int32{
	4, // 0: ProfileResIdl.error:type_name -> Error
	1, // 1: ProfileResIdl.data:type_name -> ProfileResIdl.DataRes
	5, // 2: ProfileResIdl.DataRes.user:type_name -> User
	2, // 3: ProfileResIdl.DataRes.anti_stat:type_name -> ProfileResIdl.DataRes.Anti
	6, // 4: ProfileResIdl.DataRes.post_list:type_name -> PostInfoList
	3, // 5: ProfileResIdl.DataRes.user_agree_info:type_name -> ProfileResIdl.DataRes.UserAgreeInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ProfileResIdl_proto_init() }
func file_ProfileResIdl_proto_init() {
	if File_ProfileResIdl_proto != nil {
		return
	}
	file_Error_proto_init()
	file_User_proto_init()
	file_PostInfoList_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ProfileResIdl_proto_rawDesc), len(file_ProfileResIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProfileResIdl_proto_goTypes,
		DependencyIndexes: file_ProfileResIdl_proto_depIdxs,
		MessageInfos:      file_ProfileResIdl_proto_msgTypes,
	}.Build()
	File_ProfileResIdl_proto = out.File
	file_ProfileResIdl_proto_goTypes = nil
	file_ProfileResIdl_proto_depIdxs = nil
}
