package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/edit"
)

func String(s string) *string {
	return &s
}

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

	// add custom header, for example
	// instance.AddHeader("X-TT-LOGID", "logid")

	// your param
	param := &edit.TemplateParamItem{
		Type:     "image",
		Position: "0",
		Source:   String("your source"),
	}

	// for example
	request := &edit.SubmitTemplateTaskAsyncRequest{
		TemplateId:   "your template id",
		Space:        "your space",
		VideoName:    []string{"your video name"},
		Params:       [][]*edit.TemplateParamItem{{param}},
		Priority:     0,
		CallbackArgs: "your callback args",
		CallbackUri:  "your callback uri",
		Type:         2, // 2 indicate mv template
	}

	resp, err := instance.SubmitTemplateTaskAsync(request)

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
