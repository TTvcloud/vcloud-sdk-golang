package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * get template conf
 */
func main() {
	instance := imagex.NewInstance()

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and sk as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	params := &imagex.GetTemplateConfParam{
		GroupId: "your group id",
	}

	resp, err := instance.GetTemplateConf(params)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
