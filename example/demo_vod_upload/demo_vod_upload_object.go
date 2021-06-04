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
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "your spaceName"
	filePath := "your filePath"

	// 自行设置 fileName 必须保证名称的不同，非自行设置传空串即可
	fileName := "123456789.mp4"
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	resp, err := instance.UploadObjectWithCallbackArgs(dat, spaceName, "my callback args", fileName)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
