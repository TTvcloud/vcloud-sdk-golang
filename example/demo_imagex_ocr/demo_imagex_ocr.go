package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * get image ocr
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

	//use the param when image is stored in tos
	param := &imagex.GetImageOCRParam{
		ServiceId: "xx",
		Scene: "license",
		StoreUri: "xx",
	}

	//use the param when post binary data of a image
	// param := &imagex.GetImageOCRParam{
	// 	ServiceId: "xx",
	// 	Scene: "license",
	// 	Image: []byte("图片数据"),
	// }

	resp, err := instance.GetImageOCR(param)
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}