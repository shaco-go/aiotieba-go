// tbclient.VideoInfo

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: VideoInfo.proto

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

type VideoInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VideoUrl      string                 `protobuf:"bytes,2,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	VideoDuration uint32                 `protobuf:"varint,3,opt,name=video_duration,json=videoDuration,proto3" json:"video_duration,omitempty"`
	VideoWidth    uint32                 `protobuf:"varint,4,opt,name=video_width,json=videoWidth,proto3" json:"video_width,omitempty"`
	VideoHeight   uint32                 `protobuf:"varint,5,opt,name=video_height,json=videoHeight,proto3" json:"video_height,omitempty"`
	ThumbnailUrl  string                 `protobuf:"bytes,6,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	PlayCount     int32                  `protobuf:"varint,10,opt,name=play_count,json=playCount,proto3" json:"play_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoInfo) Reset() {
	*x = VideoInfo{}
	mi := &file_VideoInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfo) ProtoMessage() {}

func (x *VideoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VideoInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfo.ProtoReflect.Descriptor instead.
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return file_VideoInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VideoInfo) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoInfo) GetVideoDuration() uint32 {
	if x != nil {
		return x.VideoDuration
	}
	return 0
}

func (x *VideoInfo) GetVideoWidth() uint32 {
	if x != nil {
		return x.VideoWidth
	}
	return 0
}

func (x *VideoInfo) GetVideoHeight() uint32 {
	if x != nil {
		return x.VideoHeight
	}
	return 0
}

func (x *VideoInfo) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *VideoInfo) GetPlayCount() int32 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

var File_VideoInfo_proto protoreflect.FileDescriptor

const file_VideoInfo_proto_rawDesc = "" +
	"\n" +
	"\x0fVideoInfo.proto\"\xd7\x01\n" +
	"\tVideoInfo\x12\x1b\n" +
	"\tvideo_url\x18\x02 \x01(\tR\bvideoUrl\x12%\n" +
	"\x0evideo_duration\x18\x03 \x01(\rR\rvideoDuration\x12\x1f\n" +
	"\vvideo_width\x18\x04 \x01(\rR\n" +
	"videoWidth\x12!\n" +
	"\fvideo_height\x18\x05 \x01(\rR\vvideoHeight\x12#\n" +
	"\rthumbnail_url\x18\x06 \x01(\tR\fthumbnailUrl\x12\x1d\n" +
	"\n" +
	"play_count\x18\n" +
	" \x01(\x05R\tplayCountB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_VideoInfo_proto_rawDescOnce sync.Once
	file_VideoInfo_proto_rawDescData []byte
)

func file_VideoInfo_proto_rawDescGZIP() []byte {
	file_VideoInfo_proto_rawDescOnce.Do(func() {
		file_VideoInfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_VideoInfo_proto_rawDesc), len(file_VideoInfo_proto_rawDesc)))
	})
	return file_VideoInfo_proto_rawDescData
}

var file_VideoInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_VideoInfo_proto_goTypes = []any{
	(*VideoInfo)(nil), // 0: VideoInfo
}
var file_VideoInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VideoInfo_proto_init() }
func file_VideoInfo_proto_init() {
	if File_VideoInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_VideoInfo_proto_rawDesc), len(file_VideoInfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VideoInfo_proto_goTypes,
		DependencyIndexes: file_VideoInfo_proto_depIdxs,
		MessageInfos:      file_VideoInfo_proto_msgTypes,
	}.Build()
	File_VideoInfo_proto = out.File
	file_VideoInfo_proto_goTypes = nil
	file_VideoInfo_proto_depIdxs = nil
}
