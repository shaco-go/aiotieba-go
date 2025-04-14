package aiotieba_go

import (
	"github.com/pkg/errors"
	"github.com/shaco-go/aiotieba-go/core"
	"github.com/tidwall/gjson"
	"net/url"
)

func GetForum(name string) (gjson.Result, error) {
	data := url.Values{
		"kw": {name},
	}
	dom, err := core.PackFormR("https://tiebac.baidu.com/c/f/frs/frsBottom", data)
	if err != nil {
		return gjson.Result{}, errors.Wrap(err, "")
	}
	return dom, nil
}
