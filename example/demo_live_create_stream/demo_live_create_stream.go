package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	ret, err := instance.CreateStream(&live.CreateStreamRequest{
		AppID: 200002,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	retString, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}
