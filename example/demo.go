package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/iam"
)

func main() {
	c := base.Credentials{
		AccessKeyID:     "This is your AK",
		SecretAccessKey: "This is your SK",
	}
	iam.DefaultInstance.SetCredential(c)

	resp, code, _ := iam.DefaultInstance.ListAccessKeys(nil)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
