package main

import (
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	vod.NewInstance().SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	vid := "your vid"
	spaceName := "your spaceName"

	query := url.Values{}
	query.Set("video_id", vid)
	// set expires time of the play auth token, defalut is 15min(900),
	// set if if you know the params' meaning exactly.
	query.Set("X-Amz-Expires", "60")

	ret, _ := vod.NewInstance().GetPlayAuthToken(query)
	fmt.Println(ret)

	query = url.Values{}
	query.Set("SpaceName", spaceName)
	// set expires time of the upload token, defalut is 15min(900),
	// set if if you know the params' meaning exactly.
	query.Set("X-Amz-Expires", "60")

	ret, _ = vod.NewInstance().GetUploadAuthToken(query)
	fmt.Println(ret)
}
