package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/rtc"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := rtc.NewInstance()
	/*
		instance.SetCredential(base.Credentials{
			AccessKeyID:     "your-ak",
			SecretAccessKey: "your-sk",
		})
	*/

	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")
	req := map[string]interface{}{
		"AppID":  "your app id",
		"RoomID": "your room id",
		"UserID": "your user id",
	}
	ret, _, _ := instance.ByteStartTranscode(req)
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
	return
}
