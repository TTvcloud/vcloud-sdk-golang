package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * upload local image files
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

	params := &imagex.ApplyUploadImageParam{
		ServiceId: "your service id",
		StoreKeys: []string{"your file name"},
	}

	filePath := "your file path"
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		os.Exit(-1)
	}

	resp, err := instance.UploadImages(params, [][]byte{dat})
	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %v", resp)
	}
}
