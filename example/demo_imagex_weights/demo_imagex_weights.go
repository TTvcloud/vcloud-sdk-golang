package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := imagex.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	serviceId := "your serviceId"

	urls, err := imagex.NewInstance().GetImagexURL(serviceId, "your uri", "your tpl[:param]", imagex.WithHttps(), imagex.WithFormat(imagex.FORMAT_WEBP))
	if err != nil {
		fmt.Printf("GetImagexURL err :%v\n", err)
	} else {
		fmt.Printf("%v\n", urls)
	}
}
