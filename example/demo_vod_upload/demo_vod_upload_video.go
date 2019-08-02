package main

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"io/ioutil"
	"os"

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

	spaceName := "your spaceName"
	filePath := "your filePath"

	snapShotFunc := vod.Function{Name: "Snapshot", Input: vod.SnapshotInput{SnapshotTime: 2.3}}
	getMetaFunc := vod.Function{Name: "GetMeta"}

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	resp, err := vod.DefaultInstance.UploadVideo(dat, spaceName, vod.VIDEO, getMetaFunc, snapShotFunc)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
