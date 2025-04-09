// protobuf.CommitPersonalMsg.CommitPersonalMsgReqIdl

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: CommitPersonalMsgReqIdl.proto

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

type CommitPersonalMsgReqIdl struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Data          *CommitPersonalMsgReqIdl_DataReq `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommitPersonalMsgReqIdl) Reset() {
	*x = CommitPersonalMsgReqIdl{}
	mi := &file_CommitPersonalMsgReqIdl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitPersonalMsgReqIdl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitPersonalMsgReqIdl) ProtoMessage() {}

func (x *CommitPersonalMsgReqIdl) ProtoReflect() protoreflect.Message {
	mi := &file_CommitPersonalMsgReqIdl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitPersonalMsgReqIdl.ProtoReflect.Descriptor instead.
func (*CommitPersonalMsgReqIdl) Descriptor() ([]byte, []int) {
	return file_CommitPersonalMsgReqIdl_proto_rawDescGZIP(), []int{0}
}

func (x *CommitPersonalMsgReqIdl) GetData() *CommitPersonalMsgReqIdl_DataReq {
	if x != nil {
		return x.Data
	}
	return nil
}

type CommitPersonalMsgReqIdl_DataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ToUid         int64                  `protobuf:"varint,2,opt,name=toUid,proto3" json:"toUid,omitempty"`
	MsgType       int32                  `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	RecordId      int64                  `protobuf:"varint,6,opt,name=recordId,proto3" json:"recordId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommitPersonalMsgReqIdl_DataReq) Reset() {
	*x = CommitPersonalMsgReqIdl_DataReq{}
	mi := &file_CommitPersonalMsgReqIdl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitPersonalMsgReqIdl_DataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitPersonalMsgReqIdl_DataReq) ProtoMessage() {}

func (x *CommitPersonalMsgReqIdl_DataReq) ProtoReflect() protoreflect.Message {
	mi := &file_CommitPersonalMsgReqIdl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitPersonalMsgReqIdl_DataReq.ProtoReflect.Descriptor instead.
func (*CommitPersonalMsgReqIdl_DataReq) Descriptor() ([]byte, []int) {
	return file_CommitPersonalMsgReqIdl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommitPersonalMsgReqIdl_DataReq) GetToUid() int64 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *CommitPersonalMsgReqIdl_DataReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *CommitPersonalMsgReqIdl_DataReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommitPersonalMsgReqIdl_DataReq) GetRecordId() int64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

var File_CommitPersonalMsgReqIdl_proto protoreflect.FileDescriptor

const file_CommitPersonalMsgReqIdl_proto_rawDesc = "" +
	"\n" +
	"\x1dCommitPersonalMsgReqIdl.proto\"\xc0\x01\n" +
	"\x17CommitPersonalMsgReqIdl\x124\n" +
	"\x04data\x18\x01 \x01(\v2 .CommitPersonalMsgReqIdl.DataReqR\x04data\x1ao\n" +
	"\aDataReq\x12\x14\n" +
	"\x05toUid\x18\x02 \x01(\x03R\x05toUid\x12\x18\n" +
	"\amsgType\x18\x03 \x01(\x05R\amsgType\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12\x1a\n" +
	"\brecordId\x18\x06 \x01(\x03R\brecordIdB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_CommitPersonalMsgReqIdl_proto_rawDescOnce sync.Once
	file_CommitPersonalMsgReqIdl_proto_rawDescData []byte
)

func file_CommitPersonalMsgReqIdl_proto_rawDescGZIP() []byte {
	file_CommitPersonalMsgReqIdl_proto_rawDescOnce.Do(func() {
		file_CommitPersonalMsgReqIdl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_CommitPersonalMsgReqIdl_proto_rawDesc), len(file_CommitPersonalMsgReqIdl_proto_rawDesc)))
	})
	return file_CommitPersonalMsgReqIdl_proto_rawDescData
}

var file_CommitPersonalMsgReqIdl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CommitPersonalMsgReqIdl_proto_goTypes = []any{
	(*CommitPersonalMsgReqIdl)(nil),         // 0: CommitPersonalMsgReqIdl
	(*CommitPersonalMsgReqIdl_DataReq)(nil), // 1: CommitPersonalMsgReqIdl.DataReq
}
var file_CommitPersonalMsgReqIdl_proto_depIdxs = []int32{
	1, // 0: CommitPersonalMsgReqIdl.data:type_name -> CommitPersonalMsgReqIdl.DataReq
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommitPersonalMsgReqIdl_proto_init() }
func file_CommitPersonalMsgReqIdl_proto_init() {
	if File_CommitPersonalMsgReqIdl_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_CommitPersonalMsgReqIdl_proto_rawDesc), len(file_CommitPersonalMsgReqIdl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommitPersonalMsgReqIdl_proto_goTypes,
		DependencyIndexes: file_CommitPersonalMsgReqIdl_proto_depIdxs,
		MessageInfos:      file_CommitPersonalMsgReqIdl_proto_msgTypes,
	}.Build()
	File_CommitPersonalMsgReqIdl_proto = out.File
	file_CommitPersonalMsgReqIdl_proto_goTypes = nil
	file_CommitPersonalMsgReqIdl_proto_depIdxs = nil
}
