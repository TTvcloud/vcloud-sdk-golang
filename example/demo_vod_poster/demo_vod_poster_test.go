package main

import (
	"testing"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func BenchmarkGetPoster(b *testing.B) {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk"})

	spaceName := "space name"
	fallbackWeights := map[string]int{
		"v1.test.com": 10,
		"v3.test.com": 20,
	}

	uri := "uri"

	for i := 0; i < b.N; i++ {
		_, err := vod.DefaultInstance.GetPosterUrl(spaceName, uri, fallbackWeights, vod.WithHttps(), vod.WithVodTplSmartCrop(600, 392), vod.WithFormat(vod.FORMAT_AWEBP))
		if err != nil {
			return
		}
	}

}
