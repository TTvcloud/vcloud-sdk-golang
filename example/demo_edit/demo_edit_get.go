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

	resp, err := instance.GetDirectEditResult(&edit.GetDirectEditResultRequest{[]string{"your req id"}})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	retString, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}
