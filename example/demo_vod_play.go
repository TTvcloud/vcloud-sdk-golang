package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	query := url.Values{}
	query.Set("video_id", "v0282ccd0000bjpm26vibkthkdq62qf0")

	resp, code, _ := vod.DefaultInstance.GetPlayInfo(query)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
