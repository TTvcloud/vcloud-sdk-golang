package main

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

const spaceName = "your-space-name"

func main() {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	resp, _ := vod.DefaultInstance.GetCdnDomainWeights(spaceName)
	fmt.Println(resp)
}
