package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/edit"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := edit.NewInstance()
	//instance.SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	// your param str
	// for example
	paramStr := `
{
	"Upload": {
		"Uploader": "your uploader",
		"VideoName": "video"
	},
	"Output": {
		"Fps": 25,
		"Height": 720,
		"Quality": "medium",
		"Width": 1280
	},
	"Segments": [{
		"BackGround": "0xFFFFFFFF",
		"Duration": 3,
		"Elements": [],
		"Volume": 1
	}],
	"GlobalElements": []
}`
	var param interface{}

	err := json.Unmarshal([]byte(paramStr), &param)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request := &edit.SubmitDirectEditTaskRequest{
		Param:    param,
		Priority: 0,
	}

	resp, err := instance.SubmitDirectEditTaskSync(request)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(1)

	retString, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}
