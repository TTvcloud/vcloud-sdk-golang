package main

import (
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	vid := "your vid"

	query := url.Values{}
	query.Set("video_id", vid)

	ret, _ := vod.DefaultInstance.GetPlayAuthToken(query)
	fmt.Println(ret)

	space := "your space"
	ret, _ = vod.DefaultInstance.GetUploadAuthToken(space)
	fmt.Println(ret)
}
