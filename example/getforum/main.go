package main

import (
	"fmt"
	aiotieba_go "github.com/shaco-go/aiotieba-go"
)

func main() {
	forum, err := aiotieba_go.GetForum("孙笑川")
	fmt.Println(forum, err)
}
