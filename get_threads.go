package aiotieba_go

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shaco-go/aiotieba-go/core"
	"github.com/shaco-go/aiotieba-go/protobuf"
	"google.golang.org/protobuf/proto"
)

// ThreadSortType 帖子排序类型
// 对于有热门分区的贴吧 0热门排序(HOT) 1按发布时间(CREATE) 2关注的人(FOLLOW) 34热门排序(HOT) >=5是按回复时间(REPLY)
// 对于无热门分区的贴吧 0按回复时间(REPLY) 1按发布时间(CREATE) 2关注的人(FOLLOW) >=3按回复时间(REPLY)
type ThreadSortType int

const (
	THREAD_SORT_TYPE_REPLY  ThreadSortType = 5
	THREAD_SORT_TYPE_CREATE ThreadSortType = 1
	THREAD_SORT_TYPE_HOT    ThreadSortType = 3
	THREAD_SORT_TYPE_FOLLOW ThreadSortType = 2
)

// GetThreads 获取帖子列表
// fname_or_fid (str | int): 贴吧名或fid 优先贴吧名
// pn (int, optional): 页码. Defaults to 1.
// rn (int, optional): 请求的条目数. Defaults to 13. Max to 13.
// sort (ThreadSortType, optional): HOT热门排序 REPLY按回复时间 CREATE按发布时间 FOLLOW关注的人. Defaults to ThreadSortType.REPLY.
// is_good (bool, optional): True则获取精品区帖子 False则获取普通区帖子. Defaults to False.
func GetThreads(req *protobuf.FrsPageReqIdl) (*protobuf.FrsPageResIdl_DataRes, error) {
	req.Data.Common = &protobuf.CommonReq{
		XClientType:    2,
		XClientVersion: "12.64.1.1",
	}
	req.Data.Rn = 13
	req.Data.RnNeed = 13
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	var resp protobuf.FrsPageResIdl
	err = core.PackProtoR("http://tiebac.baidu.com/c/f/frs/page?cmd=301001", data, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if resp.Error != nil && resp.Error.GetErrorno() != 0 {
		return nil, errors.New(fmt.Sprintf("tie ba err:%s [%d]", resp.Error.Errmsg, resp.Error.Errorno))
	}
	return resp.Data, err
}
