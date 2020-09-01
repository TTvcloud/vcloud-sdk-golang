package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
	"time"
)

func main() {
	instance := vod.NewInstance()
	ret, _ := instance.GetUploadAuth()
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetUploadAuthWithExpiredTime(time.Minute)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
