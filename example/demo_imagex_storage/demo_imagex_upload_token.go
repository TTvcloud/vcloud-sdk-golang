package main

import (
	"fmt"
	"net/url"

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

	query := url.Values{}
	query.Set("ServiceId", "your service id")
	// set expire time of the upload token, default is 15min(900),
	// set only if you know the params' meaning exactly.
	query.Set("X-Amz-Expires", "60")

	token, err := instance.GetUploadAuthToken(query)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("token %s", token)
	}
}
