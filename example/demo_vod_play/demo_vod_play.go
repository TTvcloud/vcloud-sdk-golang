package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

const spaceName = "your-space-name"

func main() {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	query := url.Values{}
	query.Set("video_id", "your vid")

	resp, code, _ := vod.DefaultInstance.GetPlayInfo(query)
	fmt.Printf("resp:%+v code:%d\n", resp, code)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
