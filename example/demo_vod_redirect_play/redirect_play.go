package main

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
	"time"
)

func main() {
	vid := "your vid"
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	params := vod.RedirectPlayParam{
		VideoID:    vid,
		Definition: vod.D1080P,
		Wartermark: "",
	}
	ret, err := vod.DefaultInstance.GetRedirectPlayUrl(params)
	fmt.Println(ret, err)
}
