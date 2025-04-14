package main

import (
	"fmt"
	aiotieba_go "github.com/shaco-go/aiotieba-go"
	"github.com/shaco-go/aiotieba-go/protobuf"
)

func main() {
	dom, err := aiotieba_go.GetThreads(&protobuf.FrsPageReqIdl{
		Data: &protobuf.FrsPageReqIdl_DataReq{
			Pn:       1,
			Kw:       "孙笑川",
			SortType: 0,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(dom)
}
