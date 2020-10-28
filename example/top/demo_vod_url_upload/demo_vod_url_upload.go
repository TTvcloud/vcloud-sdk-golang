package main

import (
	"encoding/json"
	"fmt"
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

	spaceName := "your space name"
	urlSets := make([]vod.URLSet, 0)
	urlSet := vod.URLSet{
		SourceUrl: "url",
	}
	urlSets = append(urlSets, urlSet)

	urlRequest := vod.UrlUploadParams{
		SpaceName: spaceName,
		URLSets:   urlSets,
	}

	resp, err := instance.UploadVideoByUrl(urlRequest)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s", bts)

}
