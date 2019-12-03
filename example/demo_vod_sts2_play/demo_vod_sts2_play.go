package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	instance := vod.NewInstance()
	ret, _ := instance.GetVideoPlayAuth([]string{}, []string{}, []string{})
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetVideoPlayAuthWithExpiredTime([]string{}, []string{}, []string{}, time.Minute)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
