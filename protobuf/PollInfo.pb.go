// tbclient.PollInfo

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: PollInfo.proto

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

type PollInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsMulti       int32                  `protobuf:"varint,2,opt,name=is_multi,json=isMulti,proto3" json:"is_multi,omitempty"`
	TotalNum      int64                  `protobuf:"varint,3,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	Options       []*PollInfo_PollOption `protobuf:"bytes,9,rep,name=options,proto3" json:"options,omitempty"`
	TotalPoll     int64                  `protobuf:"varint,11,opt,name=total_poll,json=totalPoll,proto3" json:"total_poll,omitempty"`
	Title         string                 `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PollInfo) Reset() {
	*x = PollInfo{}
	mi := &file_PollInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PollInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollInfo) ProtoMessage() {}

func (x *PollInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PollInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollInfo.ProtoReflect.Descriptor instead.
func (*PollInfo) Descriptor() ([]byte, []int) {
	return file_PollInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PollInfo) GetIsMulti() int32 {
	if x != nil {
		return x.IsMulti
	}
	return 0
}

func (x *PollInfo) GetTotalNum() int64 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *PollInfo) GetOptions() []*PollInfo_PollOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PollInfo) GetTotalPoll() int64 {
	if x != nil {
		return x.TotalPoll
	}
	return 0
}

func (x *PollInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PollInfo_PollOption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int64                  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PollInfo_PollOption) Reset() {
	*x = PollInfo_PollOption{}
	mi := &file_PollInfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PollInfo_PollOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollInfo_PollOption) ProtoMessage() {}

func (x *PollInfo_PollOption) ProtoReflect() protoreflect.Message {
	mi := &file_PollInfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollInfo_PollOption.ProtoReflect.Descriptor instead.
func (*PollInfo_PollOption) Descriptor() ([]byte, []int) {
	return file_PollInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PollInfo_PollOption) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *PollInfo_PollOption) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_PollInfo_proto protoreflect.FileDescriptor

const file_PollInfo_proto_rawDesc = "" +
	"\n" +
	"\x0ePollInfo.proto\"\xdb\x01\n" +
	"\bPollInfo\x12\x19\n" +
	"\bis_multi\x18\x02 \x01(\x05R\aisMulti\x12\x1b\n" +
	"\ttotal_num\x18\x03 \x01(\x03R\btotalNum\x12.\n" +
	"\aoptions\x18\t \x03(\v2\x14.PollInfo.PollOptionR\aoptions\x12\x1d\n" +
	"\n" +
	"total_poll\x18\v \x01(\x03R\ttotalPoll\x12\x14\n" +
	"\x05title\x18\f \x01(\tR\x05title\x1a2\n" +
	"\n" +
	"PollOption\x12\x10\n" +
	"\x03num\x18\x02 \x01(\x03R\x03num\x12\x12\n" +
	"\x04text\x18\x03 \x01(\tR\x04textB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_PollInfo_proto_rawDescOnce sync.Once
	file_PollInfo_proto_rawDescData []byte
)

func file_PollInfo_proto_rawDescGZIP() []byte {
	file_PollInfo_proto_rawDescOnce.Do(func() {
		file_PollInfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_PollInfo_proto_rawDesc), len(file_PollInfo_proto_rawDesc)))
	})
	return file_PollInfo_proto_rawDescData
}

var file_PollInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PollInfo_proto_goTypes = []any{
	(*PollInfo)(nil),            // 0: PollInfo
	(*PollInfo_PollOption)(nil), // 1: PollInfo.PollOption
}
var file_PollInfo_proto_depIdxs = []int32{
	1, // 0: PollInfo.options:type_name -> PollInfo.PollOption
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PollInfo_proto_init() }
func file_PollInfo_proto_init() {
	if File_PollInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_PollInfo_proto_rawDesc), len(file_PollInfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PollInfo_proto_goTypes,
		DependencyIndexes: file_PollInfo_proto_depIdxs,
		MessageInfos:      file_PollInfo_proto_msgTypes,
	}.Build()
	File_PollInfo_proto = out.File
	file_PollInfo_proto_goTypes = nil
	file_PollInfo_proto_depIdxs = nil
}
