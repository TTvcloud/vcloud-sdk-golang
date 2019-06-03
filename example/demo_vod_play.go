package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	query := url.Values{}
	query.Set("video_id", "your vid")

	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	resp, code, _ := vod.DefaultInstance.GetPlayInfo(query)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
