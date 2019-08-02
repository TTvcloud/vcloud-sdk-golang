package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.DefaultInstance.SetAccessKey("")
	//vod.DefaultInstance.SetSecretKey("")

	vid := "your vid"
	spaceName := "your spaceName"
	filePath := "your filePath"

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	posterUri, err := vod.DefaultInstance.UploadPoster(vid, dat, spaceName, vod.IMAGE)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", posterUri)
	}
}
