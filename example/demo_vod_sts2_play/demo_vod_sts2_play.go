package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	instance := vod.NewInstance()
	ret, _ := instance.GetVideoPlayAuth([]string{}, []string{}, []string{})
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
}
