package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	ret, _ := vod.DefaultInstance.SignSts2(nil, time.Hour)
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
}
