// tbclient.FrsPage.FrsPageResIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: FrsPageResIdl.proto

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

type FrsPageResIdl struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *Error                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data          *FrsPageResIdl_DataRes `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl) Reset() {
	*x = FrsPageResIdl{}
	mi := &file_FrsPageResIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl) ProtoMessage() {}

func (x *FrsPageResIdl) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0}
}

func (x *FrsPageResIdl) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FrsPageResIdl) GetData() *FrsPageResIdl_DataRes {
	if x != nil {
		return x.Data
	}
	return nil
}

type FrsPageResIdl_DataRes struct {
	state         protoimpl.MessageState                 `protogen:"open.v1"`
	Forum         *FrsPageResIdl_DataRes_ForumInfo       `protobuf:"bytes,2,opt,name=forum,proto3" json:"forum,omitempty"`
	Page          *Page                                  `protobuf:"bytes,4,opt,name=page,proto3" json:"page,omitempty"`
	ThreadList    []*ThreadInfo                          `protobuf:"bytes,7,rep,name=thread_list,json=threadList,proto3" json:"thread_list,omitempty"`
	UserList      []*User                                `protobuf:"bytes,17,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
	NavTabInfo    *FrsPageResIdl_DataRes_NavTabInfo      `protobuf:"bytes,37,opt,name=nav_tab_info,json=navTabInfo,proto3" json:"nav_tab_info,omitempty"`
	ForumRule     *FrsPageResIdl_DataRes_ForumRuleStatus `protobuf:"bytes,105,opt,name=forum_rule,json=forumRule,proto3" json:"forum_rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl_DataRes) Reset() {
	*x = FrsPageResIdl_DataRes{}
	mi := &file_FrsPageResIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl_DataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl_DataRes) ProtoMessage() {}

func (x *FrsPageResIdl_DataRes) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl_DataRes.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl_DataRes) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FrsPageResIdl_DataRes) GetForum() *FrsPageResIdl_DataRes_ForumInfo {
	if x != nil {
		return x.Forum
	}
	return nil
}

func (x *FrsPageResIdl_DataRes) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FrsPageResIdl_DataRes) GetThreadList() []*ThreadInfo {
	if x != nil {
		return x.ThreadList
	}
	return nil
}

func (x *FrsPageResIdl_DataRes) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *FrsPageResIdl_DataRes) GetNavTabInfo() *FrsPageResIdl_DataRes_NavTabInfo {
	if x != nil {
		return x.NavTabInfo
	}
	return nil
}

func (x *FrsPageResIdl_DataRes) GetForumRule() *FrsPageResIdl_DataRes_ForumRuleStatus {
	if x != nil {
		return x.ForumRule
	}
	return nil
}

type FrsPageResIdl_DataRes_ForumInfo struct {
	state         protoimpl.MessageState                     `protogen:"open.v1"`
	Id            int64                                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FirstClass    string                                     `protobuf:"bytes,3,opt,name=first_class,json=firstClass,proto3" json:"first_class,omitempty"`
	SecondClass   string                                     `protobuf:"bytes,4,opt,name=second_class,json=secondClass,proto3" json:"second_class,omitempty"`
	MemberNum     int32                                      `protobuf:"varint,9,opt,name=member_num,json=memberNum,proto3" json:"member_num,omitempty"`
	ThreadNum     int32                                      `protobuf:"varint,10,opt,name=thread_num,json=threadNum,proto3" json:"thread_num,omitempty"`
	PostNum       int32                                      `protobuf:"varint,11,opt,name=post_num,json=postNum,proto3" json:"post_num,omitempty"`
	Managers      []*FrsPageResIdl_DataRes_ForumInfo_Manager `protobuf:"bytes,17,rep,name=managers,proto3" json:"managers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl_DataRes_ForumInfo) Reset() {
	*x = FrsPageResIdl_DataRes_ForumInfo{}
	mi := &file_FrsPageResIdl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl_DataRes_ForumInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl_DataRes_ForumInfo) ProtoMessage() {}

func (x *FrsPageResIdl_DataRes_ForumInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl_DataRes_ForumInfo.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl_DataRes_ForumInfo) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetFirstClass() string {
	if x != nil {
		return x.FirstClass
	}
	return ""
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetSecondClass() string {
	if x != nil {
		return x.SecondClass
	}
	return ""
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetMemberNum() int32 {
	if x != nil {
		return x.MemberNum
	}
	return 0
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetThreadNum() int32 {
	if x != nil {
		return x.ThreadNum
	}
	return 0
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetPostNum() int32 {
	if x != nil {
		return x.PostNum
	}
	return 0
}

func (x *FrsPageResIdl_DataRes_ForumInfo) GetManagers() []*FrsPageResIdl_DataRes_ForumInfo_Manager {
	if x != nil {
		return x.Managers
	}
	return nil
}

type FrsPageResIdl_DataRes_NavTabInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tab           []*FrsTabInfo          `protobuf:"bytes,1,rep,name=tab,proto3" json:"tab,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl_DataRes_NavTabInfo) Reset() {
	*x = FrsPageResIdl_DataRes_NavTabInfo{}
	mi := &file_FrsPageResIdl_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl_DataRes_NavTabInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl_DataRes_NavTabInfo) ProtoMessage() {}

func (x *FrsPageResIdl_DataRes_NavTabInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl_DataRes_NavTabInfo.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl_DataRes_NavTabInfo) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *FrsPageResIdl_DataRes_NavTabInfo) GetTab() []*FrsTabInfo {
	if x != nil {
		return x.Tab
	}
	return nil
}

type FrsPageResIdl_DataRes_ForumRuleStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasForumRule  int32                  `protobuf:"varint,4,opt,name=has_forum_rule,json=hasForumRule,proto3" json:"has_forum_rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl_DataRes_ForumRuleStatus) Reset() {
	*x = FrsPageResIdl_DataRes_ForumRuleStatus{}
	mi := &file_FrsPageResIdl_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl_DataRes_ForumRuleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl_DataRes_ForumRuleStatus) ProtoMessage() {}

func (x *FrsPageResIdl_DataRes_ForumRuleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl_DataRes_ForumRuleStatus.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl_DataRes_ForumRuleStatus) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *FrsPageResIdl_DataRes_ForumRuleStatus) GetHasForumRule() int32 {
	if x != nil {
		return x.HasForumRule
	}
	return 0
}

type FrsPageResIdl_DataRes_ForumInfo_Manager struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrsPageResIdl_DataRes_ForumInfo_Manager) Reset() {
	*x = FrsPageResIdl_DataRes_ForumInfo_Manager{}
	mi := &file_FrsPageResIdl_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrsPageResIdl_DataRes_ForumInfo_Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrsPageResIdl_DataRes_ForumInfo_Manager) ProtoMessage() {}

func (x *FrsPageResIdl_DataRes_ForumInfo_Manager) ProtoReflect() protoreflect.Message {
	mi := &file_FrsPageResIdl_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrsPageResIdl_DataRes_ForumInfo_Manager.ProtoReflect.Descriptor instead.
func (*FrsPageResIdl_DataRes_ForumInfo_Manager) Descriptor() ([]byte, []int) {
	return file_FrsPageResIdl_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

var File_FrsPageResIdl_proto protoreflect.FileDescriptor

const file_FrsPageResIdl_proto_rawDesc = "" +
	"\n" +
	"\x13FrsPageResIdl.proto\x1a\vError.proto\x1a\n" +
	"Page.proto\x1a\x10ThreadInfo.proto\x1a\n" +
	"User.proto\x1a\x10FrsTabInfo.proto\"\x9c\x06\n" +
	"\rFrsPageResIdl\x12\x1c\n" +
	"\x05error\x18\x01 \x01(\v2\x06.ErrorR\x05error\x12*\n" +
	"\x04data\x18\x02 \x01(\v2\x16.FrsPageResIdl.DataResR\x04data\x1a\xc0\x05\n" +
	"\aDataRes\x126\n" +
	"\x05forum\x18\x02 \x01(\v2 .FrsPageResIdl.DataRes.ForumInfoR\x05forum\x12\x19\n" +
	"\x04page\x18\x04 \x01(\v2\x05.PageR\x04page\x12,\n" +
	"\vthread_list\x18\a \x03(\v2\v.ThreadInfoR\n" +
	"threadList\x12\"\n" +
	"\tuser_list\x18\x11 \x03(\v2\x05.UserR\buserList\x12C\n" +
	"\fnav_tab_info\x18% \x01(\v2!.FrsPageResIdl.DataRes.NavTabInfoR\n" +
	"navTabInfo\x12E\n" +
	"\n" +
	"forum_rule\x18i \x01(\v2&.FrsPageResIdl.DataRes.ForumRuleStatusR\tforumRule\x1a\x9d\x02\n" +
	"\tForumInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1f\n" +
	"\vfirst_class\x18\x03 \x01(\tR\n" +
	"firstClass\x12!\n" +
	"\fsecond_class\x18\x04 \x01(\tR\vsecondClass\x12\x1d\n" +
	"\n" +
	"member_num\x18\t \x01(\x05R\tmemberNum\x12\x1d\n" +
	"\n" +
	"thread_num\x18\n" +
	" \x01(\x05R\tthreadNum\x12\x19\n" +
	"\bpost_num\x18\v \x01(\x05R\apostNum\x12D\n" +
	"\bmanagers\x18\x11 \x03(\v2(.FrsPageResIdl.DataRes.ForumInfo.ManagerR\bmanagers\x1a\t\n" +
	"\aManager\x1a+\n" +
	"\n" +
	"NavTabInfo\x12\x1d\n" +
	"\x03tab\x18\x01 \x03(\v2\v.FrsTabInfoR\x03tab\x1a7\n" +
	"\x0fForumRuleStatus\x12$\n" +
	"\x0ehas_forum_rule\x18\x04 \x01(\x05R\fhasForumRuleB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_FrsPageResIdl_proto_rawDescOnce sync.Once
	file_FrsPageResIdl_proto_rawDescData []byte
)

func file_FrsPageResIdl_proto_rawDescGZIP() []byte {
	file_FrsPageResIdl_proto_rawDescOnce.Do(func() {
		file_FrsPageResIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_FrsPageResIdl_proto_rawDesc), len(file_FrsPageResIdl_proto_rawDesc)))
	})
	return file_FrsPageResIdl_proto_rawDescData
}

var file_FrsPageResIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_FrsPageResIdl_proto_goTypes = []any{
	(*FrsPageResIdl)(nil),                           // 0: FrsPageResIdl
	(*FrsPageResIdl_DataRes)(nil),                   // 1: FrsPageResIdl.DataRes
	(*FrsPageResIdl_DataRes_ForumInfo)(nil),         // 2: FrsPageResIdl.DataRes.ForumInfo
	(*FrsPageResIdl_DataRes_NavTabInfo)(nil),        // 3: FrsPageResIdl.DataRes.NavTabInfo
	(*FrsPageResIdl_DataRes_ForumRuleStatus)(nil),   // 4: FrsPageResIdl.DataRes.ForumRuleStatus
	(*FrsPageResIdl_DataRes_ForumInfo_Manager)(nil), // 5: FrsPageResIdl.DataRes.ForumInfo.Manager
	(*Error)(nil),      // 6: Error
	(*Page)(nil),       // 7: Page
	(*ThreadInfo)(nil), // 8: ThreadInfo
	(*User)(nil),       // 9: User
	(*FrsTabInfo)(nil), // 10: FrsTabInfo
}
var file_FrsPageResIdl_proto_depIdxs = []int32{
	6,  // 0: FrsPageResIdl.error:type_name -> Error
	1,  // 1: FrsPageResIdl.data:type_name -> FrsPageResIdl.DataRes
	2,  // 2: FrsPageResIdl.DataRes.forum:type_name -> FrsPageResIdl.DataRes.ForumInfo
	7,  // 3: FrsPageResIdl.DataRes.page:type_name -> Page
	8,  // 4: FrsPageResIdl.DataRes.thread_list:type_name -> ThreadInfo
	9,  // 5: FrsPageResIdl.DataRes.user_list:type_name -> User
	3,  // 6: FrsPageResIdl.DataRes.nav_tab_info:type_name -> FrsPageResIdl.DataRes.NavTabInfo
	4,  // 7: FrsPageResIdl.DataRes.forum_rule:type_name -> FrsPageResIdl.DataRes.ForumRuleStatus
	5,  // 8: FrsPageResIdl.DataRes.ForumInfo.managers:type_name -> FrsPageResIdl.DataRes.ForumInfo.Manager
	10, // 9: FrsPageResIdl.DataRes.NavTabInfo.tab:type_name -> FrsTabInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_FrsPageResIdl_proto_init() }
func file_FrsPageResIdl_proto_init() {
	if File_FrsPageResIdl_proto != nil {
		return
	}
	file_Error_proto_init()
	file_Page_proto_init()
	file_ThreadInfo_proto_init()
	file_User_proto_init()
	file_FrsTabInfo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_FrsPageResIdl_proto_rawDesc), len(file_FrsPageResIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FrsPageResIdl_proto_goTypes,
		DependencyIndexes: file_FrsPageResIdl_proto_depIdxs,
		MessageInfos:      file_FrsPageResIdl_proto_msgTypes,
	}.Build()
	File_FrsPageResIdl_proto = out.File
	file_FrsPageResIdl_proto_goTypes = nil
	file_FrsPageResIdl_proto_depIdxs = nil
}
