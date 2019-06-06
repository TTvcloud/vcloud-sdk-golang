package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func searchVideo() {

}

func deleteVideo() {

}

func modifyVideoInfo() {

}

func setVideoPlayStatus() {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your-ak",
		SecretAccessKey: "your-sk"})

	resp, code, _ := vod.DefaultInstance.SetVideoPublishStatus("space", "vidxxxxx", "Published")
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func DescribeVideoInfos() {

}

func main() {
	setVideoPlayStatus()
}
