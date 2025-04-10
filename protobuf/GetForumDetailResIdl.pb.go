// tbclient.GetForumDetail.GetForumDetailResIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: GetForumDetailResIdl.proto

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

type GetForumDetailResIdl struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Error         *Error                        `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data          *GetForumDetailResIdl_DataRes `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetForumDetailResIdl) Reset() {
	*x = GetForumDetailResIdl{}
	mi := &file_GetForumDetailResIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailResIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailResIdl) ProtoMessage() {}

func (x *GetForumDetailResIdl) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailResIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailResIdl.ProtoReflect.Descriptor instead.
func (*GetForumDetailResIdl) Descriptor() ([]byte, []int) {
	return file_GetForumDetailResIdl_proto_rawDescGZIP(), []int{0}
}

func (x *GetForumDetailResIdl) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetForumDetailResIdl) GetData() *GetForumDetailResIdl_DataRes {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetForumDetailResIdl_DataRes struct {
	state         protoimpl.MessageState                           `protogen:"open.v1"`
	ForumInfo     *GetForumDetailResIdl_DataRes_RecommendForumInfo `protobuf:"bytes,1,opt,name=forum_info,json=forumInfo,proto3" json:"forum_info,omitempty"`
	ElectionTab   *GetForumDetailResIdl_DataRes_ManagerElectionTab `protobuf:"bytes,8,opt,name=election_tab,json=electionTab,proto3" json:"election_tab,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetForumDetailResIdl_DataRes) Reset() {
	*x = GetForumDetailResIdl_DataRes{}
	mi := &file_GetForumDetailResIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailResIdl_DataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailResIdl_DataRes) ProtoMessage() {}

func (x *GetForumDetailResIdl_DataRes) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailResIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailResIdl_DataRes.ProtoReflect.Descriptor instead.
func (*GetForumDetailResIdl_DataRes) Descriptor() ([]byte, []int) {
	return file_GetForumDetailResIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetForumDetailResIdl_DataRes) GetForumInfo() *GetForumDetailResIdl_DataRes_RecommendForumInfo {
	if x != nil {
		return x.ForumInfo
	}
	return nil
}

func (x *GetForumDetailResIdl_DataRes) GetElectionTab() *GetForumDetailResIdl_DataRes_ManagerElectionTab {
	if x != nil {
		return x.ElectionTab
	}
	return nil
}

type GetForumDetailResIdl_DataRes_RecommendForumInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Avatar        string                 `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	ForumId       uint64                 `protobuf:"varint,2,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	ForumName     string                 `protobuf:"bytes,3,opt,name=forum_name,json=forumName,proto3" json:"forum_name,omitempty"`
	MemberCount   uint32                 `protobuf:"varint,5,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	ThreadCount   uint32                 `protobuf:"varint,6,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
	Slogan        string                 `protobuf:"bytes,7,opt,name=slogan,proto3" json:"slogan,omitempty"`
	Lv1Name       string                 `protobuf:"bytes,18,opt,name=lv1_name,json=lv1Name,proto3" json:"lv1_name,omitempty"`
	AvatarOrigin  string                 `protobuf:"bytes,20,opt,name=avatar_origin,json=avatarOrigin,proto3" json:"avatar_origin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) Reset() {
	*x = GetForumDetailResIdl_DataRes_RecommendForumInfo{}
	mi := &file_GetForumDetailResIdl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailResIdl_DataRes_RecommendForumInfo) ProtoMessage() {}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailResIdl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailResIdl_DataRes_RecommendForumInfo.ProtoReflect.Descriptor instead.
func (*GetForumDetailResIdl_DataRes_RecommendForumInfo) Descriptor() ([]byte, []int) {
	return file_GetForumDetailResIdl_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetForumId() uint64 {
	if x != nil {
		return x.ForumId
	}
	return 0
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetForumName() string {
	if x != nil {
		return x.ForumName
	}
	return ""
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetMemberCount() uint32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetThreadCount() uint32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetLv1Name() string {
	if x != nil {
		return x.Lv1Name
	}
	return ""
}

func (x *GetForumDetailResIdl_DataRes_RecommendForumInfo) GetAvatarOrigin() string {
	if x != nil {
		return x.AvatarOrigin
	}
	return ""
}

type GetForumDetailResIdl_DataRes_ManagerElectionTab struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	NewStrategyText string                 `protobuf:"bytes,5,opt,name=new_strategy_text,json=newStrategyText,proto3" json:"new_strategy_text,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetForumDetailResIdl_DataRes_ManagerElectionTab) Reset() {
	*x = GetForumDetailResIdl_DataRes_ManagerElectionTab{}
	mi := &file_GetForumDetailResIdl_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetForumDetailResIdl_DataRes_ManagerElectionTab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForumDetailResIdl_DataRes_ManagerElectionTab) ProtoMessage() {}

func (x *GetForumDetailResIdl_DataRes_ManagerElectionTab) ProtoReflect() protoreflect.Message {
	mi := &file_GetForumDetailResIdl_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForumDetailResIdl_DataRes_ManagerElectionTab.ProtoReflect.Descriptor instead.
func (*GetForumDetailResIdl_DataRes_ManagerElectionTab) Descriptor() ([]byte, []int) {
	return file_GetForumDetailResIdl_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *GetForumDetailResIdl_DataRes_ManagerElectionTab) GetNewStrategyText() string {
	if x != nil {
		return x.NewStrategyText
	}
	return ""
}

var File_GetForumDetailResIdl_proto protoreflect.FileDescriptor

const file_GetForumDetailResIdl_proto_rawDesc = "" +
	"\n" +
	"\x1aGetForumDetailResIdl.proto\x1a\vError.proto\"\xe2\x04\n" +
	"\x14GetForumDetailResIdl\x12\x1c\n" +
	"\x05error\x18\x01 \x01(\v2\x06.ErrorR\x05error\x121\n" +
	"\x04data\x18\x02 \x01(\v2\x1d.GetForumDetailResIdl.DataResR\x04data\x1a\xf8\x03\n" +
	"\aDataRes\x12O\n" +
	"\n" +
	"forum_info\x18\x01 \x01(\v20.GetForumDetailResIdl.DataRes.RecommendForumInfoR\tforumInfo\x12S\n" +
	"\felection_tab\x18\b \x01(\v20.GetForumDetailResIdl.DataRes.ManagerElectionTabR\velectionTab\x1a\x84\x02\n" +
	"\x12RecommendForumInfo\x12\x16\n" +
	"\x06avatar\x18\x01 \x01(\tR\x06avatar\x12\x19\n" +
	"\bforum_id\x18\x02 \x01(\x04R\aforumId\x12\x1d\n" +
	"\n" +
	"forum_name\x18\x03 \x01(\tR\tforumName\x12!\n" +
	"\fmember_count\x18\x05 \x01(\rR\vmemberCount\x12!\n" +
	"\fthread_count\x18\x06 \x01(\rR\vthreadCount\x12\x16\n" +
	"\x06slogan\x18\a \x01(\tR\x06slogan\x12\x19\n" +
	"\blv1_name\x18\x12 \x01(\tR\alv1Name\x12#\n" +
	"\ravatar_origin\x18\x14 \x01(\tR\favatarOrigin\x1a@\n" +
	"\x12ManagerElectionTab\x12*\n" +
	"\x11new_strategy_text\x18\x05 \x01(\tR\x0fnewStrategyTextB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_GetForumDetailResIdl_proto_rawDescOnce sync.Once
	file_GetForumDetailResIdl_proto_rawDescData []byte
)

func file_GetForumDetailResIdl_proto_rawDescGZIP() []byte {
	file_GetForumDetailResIdl_proto_rawDescOnce.Do(func() {
		file_GetForumDetailResIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_GetForumDetailResIdl_proto_rawDesc), len(file_GetForumDetailResIdl_proto_rawDesc)))
	})
	return file_GetForumDetailResIdl_proto_rawDescData
}

var file_GetForumDetailResIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_GetForumDetailResIdl_proto_goTypes = []any{
	(*GetForumDetailResIdl)(nil),                            // 0: GetForumDetailResIdl
	(*GetForumDetailResIdl_DataRes)(nil),                    // 1: GetForumDetailResIdl.DataRes
	(*GetForumDetailResIdl_DataRes_RecommendForumInfo)(nil), // 2: GetForumDetailResIdl.DataRes.RecommendForumInfo
	(*GetForumDetailResIdl_DataRes_ManagerElectionTab)(nil), // 3: GetForumDetailResIdl.DataRes.ManagerElectionTab
	(*Error)(nil), // 4: Error
}
var file_GetForumDetailResIdl_proto_depIdxs = []int32{
	4, // 0: GetForumDetailResIdl.error:type_name -> Error
	1, // 1: GetForumDetailResIdl.data:type_name -> GetForumDetailResIdl.DataRes
	2, // 2: GetForumDetailResIdl.DataRes.forum_info:type_name -> GetForumDetailResIdl.DataRes.RecommendForumInfo
	3, // 3: GetForumDetailResIdl.DataRes.election_tab:type_name -> GetForumDetailResIdl.DataRes.ManagerElectionTab
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GetForumDetailResIdl_proto_init() }
func file_GetForumDetailResIdl_proto_init() {
	if File_GetForumDetailResIdl_proto != nil {
		return
	}
	file_Error_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_GetForumDetailResIdl_proto_rawDesc), len(file_GetForumDetailResIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetForumDetailResIdl_proto_goTypes,
		DependencyIndexes: file_GetForumDetailResIdl_proto_depIdxs,
		MessageInfos:      file_GetForumDetailResIdl_proto_msgTypes,
	}.Build()
	File_GetForumDetailResIdl_proto = out.File
	file_GetForumDetailResIdl_proto_goTypes = nil
	file_GetForumDetailResIdl_proto_depIdxs = nil
}
