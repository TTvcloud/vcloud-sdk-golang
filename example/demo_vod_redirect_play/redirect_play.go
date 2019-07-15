package main

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	vid := "your vid"
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	params := vod.RedirectPlayParam{
		Vid:        vid,
		Definition: vod.D1080P,
		Watermark:  "",
	}
	ret, err := vod.DefaultInstance.GetRedirectPlayUrl(params)
	fmt.Println(ret, err)
}
