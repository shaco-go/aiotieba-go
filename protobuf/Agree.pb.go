// tbclient.Agree

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: Agree.proto

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

type Agree struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgreeNum      int64                  `protobuf:"varint,1,opt,name=agree_num,json=agreeNum,proto3" json:"agree_num,omitempty"`
	DisagreeNum   int64                  `protobuf:"varint,4,opt,name=disagree_num,json=disagreeNum,proto3" json:"disagree_num,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Agree) Reset() {
	*x = Agree{}
	mi := &file_Agree_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Agree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agree) ProtoMessage() {}

func (x *Agree) ProtoReflect() protoreflect.Message {
	mi := &file_Agree_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agree.ProtoReflect.Descriptor instead.
func (*Agree) Descriptor() ([]byte, []int) {
	return file_Agree_proto_rawDescGZIP(), []int{0}
}

func (x *Agree) GetAgreeNum() int64 {
	if x != nil {
		return x.AgreeNum
	}
	return 0
}

func (x *Agree) GetDisagreeNum() int64 {
	if x != nil {
		return x.DisagreeNum
	}
	return 0
}

var File_Agree_proto protoreflect.FileDescriptor

const file_Agree_proto_rawDesc = "" +
	"\n" +
	"\vAgree.proto\"G\n" +
	"\x05Agree\x12\x1b\n" +
	"\tagree_num\x18\x01 \x01(\x03R\bagreeNum\x12!\n" +
	"\fdisagree_num\x18\x04 \x01(\x03R\vdisagreeNumB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_Agree_proto_rawDescOnce sync.Once
	file_Agree_proto_rawDescData []byte
)

func file_Agree_proto_rawDescGZIP() []byte {
	file_Agree_proto_rawDescOnce.Do(func() {
		file_Agree_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Agree_proto_rawDesc), len(file_Agree_proto_rawDesc)))
	})
	return file_Agree_proto_rawDescData
}

var file_Agree_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Agree_proto_goTypes = []any{
	(*Agree)(nil), // 0: Agree
}
var file_Agree_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Agree_proto_init() }
func file_Agree_proto_init() {
	if File_Agree_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Agree_proto_rawDesc), len(file_Agree_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Agree_proto_goTypes,
		DependencyIndexes: file_Agree_proto_depIdxs,
		MessageInfos:      file_Agree_proto_msgTypes,
	}.Build()
	File_Agree_proto = out.File
	file_Agree_proto_goTypes = nil
	file_Agree_proto_depIdxs = nil
}
