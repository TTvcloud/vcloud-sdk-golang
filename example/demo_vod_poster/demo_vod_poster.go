package main

import (
	"fmt"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your space"
	fallbackWeights := map[string]int{
		"v1.test.com": 10,
		"v3.test.com": 20,
	}

	// domain
	for i := 0; i < 20; i++ {
		ret, err := instance.GetDomainInfo(spaceName, fallbackWeights)
		fmt.Println(ret)
		if err != nil {
			fmt.Printf("errMsg:%v", err)
			return
		}

		time.Sleep(1 * time.Second)
	}

	uri := "your uri"
	// poster
	ret, err := instance.GetPosterUrl(spaceName, uri, fallbackWeights, vod.WithHttps(), vod.WithVodTplSmartCrop(600, 392), vod.WithFormat(vod.FORMAT_AWEBP))
	if err != nil {
		fmt.Printf("errMsg:%v", err)
		return
	}
	fmt.Println(ret)
}
