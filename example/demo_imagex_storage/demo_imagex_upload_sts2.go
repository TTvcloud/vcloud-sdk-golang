package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * get image upload token
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

	// service id list allowed to do upload, pass empty list if no restriction
	serviceIds := []string{"your service id"}

	// set expire time by GetUploadAuthWithExpire, default is 1 hour
	token, err := instance.GetUploadAuth(serviceIds)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("token %+v", token)
	}
}
