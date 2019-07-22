package main

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	spaceName := "your spaceName"
	fallbackWeights := map[string]int{
		"v1.test.com": 10,
		"v3.test.com": 20,
	}

	// domain
	for i := 0; i < 20; i++ {
		ret, err := vod.DefaultInstance.GetDomainInfo(spaceName, fallbackWeights)
		fmt.Println(ret)
		if err != nil {
			fmt.Printf("errMsg:%v", err)
			return
		}

		time.Sleep(1 * time.Second)
	}

	uri := "your uri"
	// poster
	ret, err := vod.DefaultInstance.GetPosterUrl(spaceName, uri, fallbackWeights, vod.WithHttps(), vod.WithVodTplSmartCrop(600, 392), vod.WithFormat(vod.FORMAT_AWEBP))
	if err != nil {
		fmt.Printf("errMsg:%v", err)
		return
	}
	fmt.Println(ret)

	// image x
	kv := url.Values{}
	kv.Add("from", "my测试")

	sig := "your sig"

	ret, err = vod.DefaultInstance.GetImageUrl(spaceName, uri, fallbackWeights, vod.WithFormat(vod.FORMAT_AWEBP), vod.WithHttps(), vod.WithSig(sig), vod.WithKV(kv))
	if err != nil {
		fmt.Printf("errMsg:%v", err)
		return
	}
	fmt.Println(ret)
}
