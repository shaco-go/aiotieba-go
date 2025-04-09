package aiotieba_go

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shaco-go/aiotieba-go/core"
	"github.com/shaco-go/aiotieba-go/protobuf"
	"net/url"
)

func GetForum(name string) (*protobuf.FrsPageResIdl_DataRes, error) {
	data := url.Values{
		"kw": {name},
	}
	dom, err := core.PackFormR("https://tiebac.baidu.com/c/f/frs/frsBottom", data)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	fmt.Println(dom)
	return nil, err
}
