// tbclient.ThreadInfo

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: ThreadInfo.proto

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

type ThreadInfo struct {
	state                protoimpl.MessageState       `protogen:"open.v1"`
	Id                   int64                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string                       `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ReplyNum             int32                        `protobuf:"varint,4,opt,name=reply_num,json=replyNum,proto3" json:"reply_num,omitempty"`
	ViewNum              int32                        `protobuf:"varint,5,opt,name=view_num,json=viewNum,proto3" json:"view_num,omitempty"`
	LastTimeInt          int32                        `protobuf:"varint,7,opt,name=last_time_int,json=lastTimeInt,proto3" json:"last_time_int,omitempty"`
	IsTop                int32                        `protobuf:"varint,9,opt,name=is_top,json=isTop,proto3" json:"is_top,omitempty"`
	IsGood               int32                        `protobuf:"varint,10,opt,name=is_good,json=isGood,proto3" json:"is_good,omitempty"`
	IsVoiceThread        int32                        `protobuf:"varint,15,opt,name=is_voice_thread,json=isVoiceThread,proto3" json:"is_voice_thread,omitempty"`
	Author               *User                        `protobuf:"bytes,18,opt,name=author,proto3" json:"author,omitempty"`
	VoiceInfo            []*Voice                     `protobuf:"bytes,23,rep,name=voice_info,json=voiceInfo,proto3" json:"voice_info,omitempty"`
	ThreadType           int32                        `protobuf:"varint,26,opt,name=thread_type,json=threadType,proto3" json:"thread_type,omitempty"`
	Fid                  int64                        `protobuf:"varint,27,opt,name=fid,proto3" json:"fid,omitempty"`
	Fname                string                       `protobuf:"bytes,28,opt,name=fname,proto3" json:"fname,omitempty"`
	IsLivepost           int32                        `protobuf:"varint,30,opt,name=is_livepost,json=isLivepost,proto3" json:"is_livepost,omitempty"`
	FirstPostId          int64                        `protobuf:"varint,40,opt,name=first_post_id,json=firstPostId,proto3" json:"first_post_id,omitempty"`
	CreateTime           int32                        `protobuf:"varint,45,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	PostId               int64                        `protobuf:"varint,52,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	AuthorId             int64                        `protobuf:"varint,56,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	IsAd                 uint32                       `protobuf:"varint,59,opt,name=is_ad,json=isAd,proto3" json:"is_ad,omitempty"`
	PollInfo             *PollInfo                    `protobuf:"bytes,74,opt,name=poll_info,json=pollInfo,proto3" json:"poll_info,omitempty"`
	VideoInfo            *VideoInfo                   `protobuf:"bytes,79,opt,name=video_info,json=videoInfo,proto3" json:"video_info,omitempty"`
	IsGodthreadRecommend int32                        `protobuf:"varint,85,opt,name=is_godthread_recommend,json=isGodthreadRecommend,proto3" json:"is_godthread_recommend,omitempty"`
	Agree                *Agree                       `protobuf:"bytes,126,opt,name=agree,proto3" json:"agree,omitempty"`
	ShareNum             int32                        `protobuf:"varint,135,opt,name=share_num,json=shareNum,proto3" json:"share_num,omitempty"`
	OriginThreadInfo     *ThreadInfo_OriginThreadInfo `protobuf:"bytes,141,opt,name=origin_thread_info,json=originThreadInfo,proto3" json:"origin_thread_info,omitempty"`
	FirstPostContent     []*PbContent                 `protobuf:"bytes,142,rep,name=first_post_content,json=firstPostContent,proto3" json:"first_post_content,omitempty"`
	IsShareThread        int32                        `protobuf:"varint,143,opt,name=is_share_thread,json=isShareThread,proto3" json:"is_share_thread,omitempty"`
	TabId                int32                        `protobuf:"varint,175,opt,name=tab_id,json=tabId,proto3" json:"tab_id,omitempty"`
	IsDeleted            int32                        `protobuf:"varint,181,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	IsFrsMask            int32                        `protobuf:"varint,198,opt,name=is_frs_mask,json=isFrsMask,proto3" json:"is_frs_mask,omitempty"`
	CustomFigure         *ThreadInfo_CustomFigure     `protobuf:"bytes,211,opt,name=custom_figure,json=customFigure,proto3" json:"custom_figure,omitempty"`
	CustomState          *ThreadInfo_CustomState      `protobuf:"bytes,212,opt,name=custom_state,json=customState,proto3" json:"custom_state,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ThreadInfo) Reset() {
	*x = ThreadInfo{}
	mi := &file_ThreadInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfo) ProtoMessage() {}

func (x *ThreadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfo.ProtoReflect.Descriptor instead.
func (*ThreadInfo) Descriptor() ([]byte, []int) {
	return file_ThreadInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ThreadInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ThreadInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ThreadInfo) GetReplyNum() int32 {
	if x != nil {
		return x.ReplyNum
	}
	return 0
}

func (x *ThreadInfo) GetViewNum() int32 {
	if x != nil {
		return x.ViewNum
	}
	return 0
}

func (x *ThreadInfo) GetLastTimeInt() int32 {
	if x != nil {
		return x.LastTimeInt
	}
	return 0
}

func (x *ThreadInfo) GetIsTop() int32 {
	if x != nil {
		return x.IsTop
	}
	return 0
}

func (x *ThreadInfo) GetIsGood() int32 {
	if x != nil {
		return x.IsGood
	}
	return 0
}

func (x *ThreadInfo) GetIsVoiceThread() int32 {
	if x != nil {
		return x.IsVoiceThread
	}
	return 0
}

func (x *ThreadInfo) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *ThreadInfo) GetVoiceInfo() []*Voice {
	if x != nil {
		return x.VoiceInfo
	}
	return nil
}

func (x *ThreadInfo) GetThreadType() int32 {
	if x != nil {
		return x.ThreadType
	}
	return 0
}

func (x *ThreadInfo) GetFid() int64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *ThreadInfo) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *ThreadInfo) GetIsLivepost() int32 {
	if x != nil {
		return x.IsLivepost
	}
	return 0
}

func (x *ThreadInfo) GetFirstPostId() int64 {
	if x != nil {
		return x.FirstPostId
	}
	return 0
}

func (x *ThreadInfo) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ThreadInfo) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *ThreadInfo) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *ThreadInfo) GetIsAd() uint32 {
	if x != nil {
		return x.IsAd
	}
	return 0
}

func (x *ThreadInfo) GetPollInfo() *PollInfo {
	if x != nil {
		return x.PollInfo
	}
	return nil
}

func (x *ThreadInfo) GetVideoInfo() *VideoInfo {
	if x != nil {
		return x.VideoInfo
	}
	return nil
}

func (x *ThreadInfo) GetIsGodthreadRecommend() int32 {
	if x != nil {
		return x.IsGodthreadRecommend
	}
	return 0
}

func (x *ThreadInfo) GetAgree() *Agree {
	if x != nil {
		return x.Agree
	}
	return nil
}

func (x *ThreadInfo) GetShareNum() int32 {
	if x != nil {
		return x.ShareNum
	}
	return 0
}

func (x *ThreadInfo) GetOriginThreadInfo() *ThreadInfo_OriginThreadInfo {
	if x != nil {
		return x.OriginThreadInfo
	}
	return nil
}

func (x *ThreadInfo) GetFirstPostContent() []*PbContent {
	if x != nil {
		return x.FirstPostContent
	}
	return nil
}

func (x *ThreadInfo) GetIsShareThread() int32 {
	if x != nil {
		return x.IsShareThread
	}
	return 0
}

func (x *ThreadInfo) GetTabId() int32 {
	if x != nil {
		return x.TabId
	}
	return 0
}

func (x *ThreadInfo) GetIsDeleted() int32 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

func (x *ThreadInfo) GetIsFrsMask() int32 {
	if x != nil {
		return x.IsFrsMask
	}
	return 0
}

func (x *ThreadInfo) GetCustomFigure() *ThreadInfo_CustomFigure {
	if x != nil {
		return x.CustomFigure
	}
	return nil
}

func (x *ThreadInfo) GetCustomState() *ThreadInfo_CustomState {
	if x != nil {
		return x.CustomState
	}
	return nil
}

type ThreadInfo_OriginThreadInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Media         []*Media               `protobuf:"bytes,2,rep,name=media,proto3" json:"media,omitempty"`
	Fname         string                 `protobuf:"bytes,4,opt,name=fname,proto3" json:"fname,omitempty"`
	Tid           string                 `protobuf:"bytes,5,opt,name=tid,proto3" json:"tid,omitempty"`
	Fid           int64                  `protobuf:"varint,7,opt,name=fid,proto3" json:"fid,omitempty"`
	VoiceInfo     []*Voice               `protobuf:"bytes,12,rep,name=voice_info,json=voiceInfo,proto3" json:"voice_info,omitempty"`
	VideoInfo     *VideoInfo             `protobuf:"bytes,13,opt,name=video_info,json=videoInfo,proto3" json:"video_info,omitempty"`
	Content       []*PbContent           `protobuf:"bytes,14,rep,name=content,proto3" json:"content,omitempty"`
	PollInfo      *PollInfo              `protobuf:"bytes,21,opt,name=poll_info,json=pollInfo,proto3" json:"poll_info,omitempty"`
	Pid           int64                  `protobuf:"varint,25,opt,name=pid,proto3" json:"pid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ThreadInfo_OriginThreadInfo) Reset() {
	*x = ThreadInfo_OriginThreadInfo{}
	mi := &file_ThreadInfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadInfo_OriginThreadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfo_OriginThreadInfo) ProtoMessage() {}

func (x *ThreadInfo_OriginThreadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadInfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfo_OriginThreadInfo.ProtoReflect.Descriptor instead.
func (*ThreadInfo_OriginThreadInfo) Descriptor() ([]byte, []int) {
	return file_ThreadInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ThreadInfo_OriginThreadInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ThreadInfo_OriginThreadInfo) GetMedia() []*Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *ThreadInfo_OriginThreadInfo) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *ThreadInfo_OriginThreadInfo) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *ThreadInfo_OriginThreadInfo) GetFid() int64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *ThreadInfo_OriginThreadInfo) GetVoiceInfo() []*Voice {
	if x != nil {
		return x.VoiceInfo
	}
	return nil
}

func (x *ThreadInfo_OriginThreadInfo) GetVideoInfo() *VideoInfo {
	if x != nil {
		return x.VideoInfo
	}
	return nil
}

func (x *ThreadInfo_OriginThreadInfo) GetContent() []*PbContent {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ThreadInfo_OriginThreadInfo) GetPollInfo() *PollInfo {
	if x != nil {
		return x.PollInfo
	}
	return nil
}

func (x *ThreadInfo_OriginThreadInfo) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type ThreadInfo_CustomFigure struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	BackgroundValue string                 `protobuf:"bytes,3,opt,name=background_value,json=backgroundValue,proto3" json:"background_value,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ThreadInfo_CustomFigure) Reset() {
	*x = ThreadInfo_CustomFigure{}
	mi := &file_ThreadInfo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadInfo_CustomFigure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfo_CustomFigure) ProtoMessage() {}

func (x *ThreadInfo_CustomFigure) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadInfo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfo_CustomFigure.ProtoReflect.Descriptor instead.
func (*ThreadInfo_CustomFigure) Descriptor() ([]byte, []int) {
	return file_ThreadInfo_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ThreadInfo_CustomFigure) GetBackgroundValue() string {
	if x != nil {
		return x.BackgroundValue
	}
	return ""
}

type ThreadInfo_CustomState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ThreadInfo_CustomState) Reset() {
	*x = ThreadInfo_CustomState{}
	mi := &file_ThreadInfo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadInfo_CustomState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfo_CustomState) ProtoMessage() {}

func (x *ThreadInfo_CustomState) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadInfo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfo_CustomState.ProtoReflect.Descriptor instead.
func (*ThreadInfo_CustomState) Descriptor() ([]byte, []int) {
	return file_ThreadInfo_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ThreadInfo_CustomState) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_ThreadInfo_proto protoreflect.FileDescriptor

const file_ThreadInfo_proto_rawDesc = "" +
	"\n" +
	"\x10ThreadInfo.proto\x1a\n" +
	"User.proto\x1a\vVoice.proto\x1a\x0ePollInfo.proto\x1a\x0fVideoInfo.proto\x1a\x0fPbContent.proto\x1a\vAgree.proto\x1a\vMedia.proto\"\x8b\f\n" +
	"\n" +
	"ThreadInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x1b\n" +
	"\treply_num\x18\x04 \x01(\x05R\breplyNum\x12\x19\n" +
	"\bview_num\x18\x05 \x01(\x05R\aviewNum\x12\"\n" +
	"\rlast_time_int\x18\a \x01(\x05R\vlastTimeInt\x12\x15\n" +
	"\x06is_top\x18\t \x01(\x05R\x05isTop\x12\x17\n" +
	"\ais_good\x18\n" +
	" \x01(\x05R\x06isGood\x12&\n" +
	"\x0fis_voice_thread\x18\x0f \x01(\x05R\risVoiceThread\x12\x1d\n" +
	"\x06author\x18\x12 \x01(\v2\x05.UserR\x06author\x12%\n" +
	"\n" +
	"voice_info\x18\x17 \x03(\v2\x06.VoiceR\tvoiceInfo\x12\x1f\n" +
	"\vthread_type\x18\x1a \x01(\x05R\n" +
	"threadType\x12\x10\n" +
	"\x03fid\x18\x1b \x01(\x03R\x03fid\x12\x14\n" +
	"\x05fname\x18\x1c \x01(\tR\x05fname\x12\x1f\n" +
	"\vis_livepost\x18\x1e \x01(\x05R\n" +
	"isLivepost\x12\"\n" +
	"\rfirst_post_id\x18( \x01(\x03R\vfirstPostId\x12\x1f\n" +
	"\vcreate_time\x18- \x01(\x05R\n" +
	"createTime\x12\x17\n" +
	"\apost_id\x184 \x01(\x03R\x06postId\x12\x1b\n" +
	"\tauthor_id\x188 \x01(\x03R\bauthorId\x12\x13\n" +
	"\x05is_ad\x18; \x01(\rR\x04isAd\x12&\n" +
	"\tpoll_info\x18J \x01(\v2\t.PollInfoR\bpollInfo\x12)\n" +
	"\n" +
	"video_info\x18O \x01(\v2\n" +
	".VideoInfoR\tvideoInfo\x124\n" +
	"\x16is_godthread_recommend\x18U \x01(\x05R\x14isGodthreadRecommend\x12\x1c\n" +
	"\x05agree\x18~ \x01(\v2\x06.AgreeR\x05agree\x12\x1c\n" +
	"\tshare_num\x18\x87\x01 \x01(\x05R\bshareNum\x12K\n" +
	"\x12origin_thread_info\x18\x8d\x01 \x01(\v2\x1c.ThreadInfo.OriginThreadInfoR\x10originThreadInfo\x129\n" +
	"\x12first_post_content\x18\x8e\x01 \x03(\v2\n" +
	".PbContentR\x10firstPostContent\x12'\n" +
	"\x0fis_share_thread\x18\x8f\x01 \x01(\x05R\risShareThread\x12\x16\n" +
	"\x06tab_id\x18\xaf\x01 \x01(\x05R\x05tabId\x12\x1e\n" +
	"\n" +
	"is_deleted\x18\xb5\x01 \x01(\x05R\tisDeleted\x12\x1f\n" +
	"\vis_frs_mask\x18\xc6\x01 \x01(\x05R\tisFrsMask\x12>\n" +
	"\rcustom_figure\x18\xd3\x01 \x01(\v2\x18.ThreadInfo.CustomFigureR\fcustomFigure\x12;\n" +
	"\fcustom_state\x18\xd4\x01 \x01(\v2\x17.ThreadInfo.CustomStateR\vcustomState\x1a\xb2\x02\n" +
	"\x10OriginThreadInfo\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x1c\n" +
	"\x05media\x18\x02 \x03(\v2\x06.MediaR\x05media\x12\x14\n" +
	"\x05fname\x18\x04 \x01(\tR\x05fname\x12\x10\n" +
	"\x03tid\x18\x05 \x01(\tR\x03tid\x12\x10\n" +
	"\x03fid\x18\a \x01(\x03R\x03fid\x12%\n" +
	"\n" +
	"voice_info\x18\f \x03(\v2\x06.VoiceR\tvoiceInfo\x12)\n" +
	"\n" +
	"video_info\x18\r \x01(\v2\n" +
	".VideoInfoR\tvideoInfo\x12$\n" +
	"\acontent\x18\x0e \x03(\v2\n" +
	".PbContentR\acontent\x12&\n" +
	"\tpoll_info\x18\x15 \x01(\v2\t.PollInfoR\bpollInfo\x12\x10\n" +
	"\x03pid\x18\x19 \x01(\x03R\x03pid\x1a9\n" +
	"\fCustomFigure\x12)\n" +
	"\x10background_value\x18\x03 \x01(\tR\x0fbackgroundValue\x1a'\n" +
	"\vCustomState\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontentB\fZ\n" +
	".;protobufb\x06proto3"

var (
	file_ThreadInfo_proto_rawDescOnce sync.Once
	file_ThreadInfo_proto_rawDescData []byte
)

func file_ThreadInfo_proto_rawDescGZIP() []byte {
	file_ThreadInfo_proto_rawDescOnce.Do(func() {
		file_ThreadInfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ThreadInfo_proto_rawDesc), len(file_ThreadInfo_proto_rawDesc)))
	})
	return file_ThreadInfo_proto_rawDescData
}

var file_ThreadInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ThreadInfo_proto_goTypes = []any{
	(*ThreadInfo)(nil),                  // 0: ThreadInfo
	(*ThreadInfo_OriginThreadInfo)(nil), // 1: ThreadInfo.OriginThreadInfo
	(*ThreadInfo_CustomFigure)(nil),     // 2: ThreadInfo.CustomFigure
	(*ThreadInfo_CustomState)(nil),      // 3: ThreadInfo.CustomState
	(*User)(nil),                        // 4: User
	(*Voice)(nil),                       // 5: Voice
	(*PollInfo)(nil),                    // 6: PollInfo
	(*VideoInfo)(nil),                   // 7: VideoInfo
	(*Agree)(nil),                       // 8: Agree
	(*PbContent)(nil),                   // 9: PbContent
	(*Media)(nil),                       // 10: Media
}
var file_ThreadInfo_proto_depIdxs = []int32{
	4,  // 0: ThreadInfo.author:type_name -> User
	5,  // 1: ThreadInfo.voice_info:type_name -> Voice
	6,  // 2: ThreadInfo.poll_info:type_name -> PollInfo
	7,  // 3: ThreadInfo.video_info:type_name -> VideoInfo
	8,  // 4: ThreadInfo.agree:type_name -> Agree
	1,  // 5: ThreadInfo.origin_thread_info:type_name -> ThreadInfo.OriginThreadInfo
	9,  // 6: ThreadInfo.first_post_content:type_name -> PbContent
	2,  // 7: ThreadInfo.custom_figure:type_name -> ThreadInfo.CustomFigure
	3,  // 8: ThreadInfo.custom_state:type_name -> ThreadInfo.CustomState
	10, // 9: ThreadInfo.OriginThreadInfo.media:type_name -> Media
	5,  // 10: ThreadInfo.OriginThreadInfo.voice_info:type_name -> Voice
	7,  // 11: ThreadInfo.OriginThreadInfo.video_info:type_name -> VideoInfo
	9,  // 12: ThreadInfo.OriginThreadInfo.content:type_name -> PbContent
	6,  // 13: ThreadInfo.OriginThreadInfo.poll_info:type_name -> PollInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_ThreadInfo_proto_init() }
func file_ThreadInfo_proto_init() {
	if File_ThreadInfo_proto != nil {
		return
	}
	file_User_proto_init()
	file_Voice_proto_init()
	file_PollInfo_proto_init()
	file_VideoInfo_proto_init()
	file_PbContent_proto_init()
	file_Agree_proto_init()
	file_Media_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ThreadInfo_proto_rawDesc), len(file_ThreadInfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ThreadInfo_proto_goTypes,
		DependencyIndexes: file_ThreadInfo_proto_depIdxs,
		MessageInfos:      file_ThreadInfo_proto_msgTypes,
	}.Build()
	File_ThreadInfo_proto = out.File
	file_ThreadInfo_proto_goTypes = nil
	file_ThreadInfo_proto_depIdxs = nil
}
