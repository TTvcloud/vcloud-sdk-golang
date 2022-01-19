package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * delete images
 */
func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)
	instance := imagex.NewInstance()

	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and sk as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	serviceId := "your service_id"
	url := "resource url"

	resp, err := instance.FetchImageUrl(serviceId, url)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
