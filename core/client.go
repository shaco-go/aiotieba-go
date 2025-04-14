package core

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"
	"net/http"
	"net/url"
	"resty.dev/v3"
	"sync"
	"time"
)

var client *resty.Client
var once sync.Once
var limiter = rate.NewLimiter(30, 30)

var appProtoHeaders = map[string]string{
	"Accept-Encoding": "gzip",
	"Connection":      "keep-alive",
	"Host":            "tiebac.baidu.com",
	"User-Agent":      "aiotieba/4.5.2",
	"x_bd_data_type":  "protobuf",
}

var appHeaders = map[string]string{
	"Accept-Encoding": "gzip",
	"Connection":      "keep-alive",
	"Host":            "tiebac.baidu.com",
	"User-Agent":      "aiotieba/4.5.2",
}

func GetInstance() *resty.Client {
	once.Do(func() {
		client = resty.New().SetRetryCount(3).
			SetRetryWaitTime(time.Second).
			SetRetryMaxWaitTime(time.Second).
			AddRetryConditions(func(response *resty.Response, err error) bool {
				if response.StatusCode() == http.StatusTooManyRequests {
					return true
				}
				return false
			})
	})
	return client
}

func PackProtoR(url string, data []byte, resp proto.Message) error {
	err := limiter.Wait(context.TODO())
	if err != nil {
		return errors.Wrap(err, "")
	}
	body, err := GetInstance().R().
		SetHeaders(appProtoHeaders).
		SetMultipartField("data", "file", "application/octet-stream", bytes.NewReader(data)).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if body.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("http code err:%s [%d]", body.String(), body.StatusCode()))
	}
	err = proto.Unmarshal(body.Bytes(), resp)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func PackFormR(url string, data url.Values) (gjson.Result, error) {
	err := limiter.Wait(context.TODO())
	if err != nil {
		return gjson.Result{}, errors.Wrap(err, "")
	}
	b := Sign(data).Encode()
	body, err := GetInstance().R().
		SetHeaders(appHeaders).
		SetHeader("Content-Type", "x-www-form-urlencoded").
		SetBody([]byte(b)).
		Post(url)
	if err != nil {
		return gjson.Result{}, errors.Wrap(err, "")
	}
	dom := gjson.ParseBytes(body.Bytes())
	return dom, nil
}
