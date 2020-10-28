package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod/top/functions"
	"io/ioutil"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	//TODO golang 检测引发冲突

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "AKLTNDQ2YTRlNTBiYTg1NDcyNmE3MDA1MTUzNzc5MWMwNmI",
		SecretAccessKey: "1ZOtyBZ89VERZdOfiUrPf24a3tTjRo1XIJbzccVHMrBvZo1jEn60LjClP2t05qWz",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	spaceName := "james-test"
	filePath := "/Users/bytedance/Downloads/objects.mp4"

	//TODO 预定义Function
	snapShotFunc := functions.SnapshotFunc(2.3)
	getMetaFunc := functions.GetMeatFunc()

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	resp, err := instance.UploadVideoWithCallback(dat, spaceName, "my callback", getMetaFunc, snapShotFunc)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		bts, _ := json.Marshal(resp)
		fmt.Printf("\nresp = %s", bts)
	}
}