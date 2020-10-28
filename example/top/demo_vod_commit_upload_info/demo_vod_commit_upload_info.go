package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	space := "your space"
	session := ""

	functions := make([]vod.Function, 0)

	snapShotFunc := vod.Function{Name: "Snapshot", Input: vod.SnapshotInput{SnapshotTime: 2.3}}
	getMetaFunc := vod.Function{Name: "GetMeta"}

	functions = append(functions, snapShotFunc)
	functions = append(functions, getMetaFunc)

	fbts, err := json.Marshal(functions)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("\n%s", fbts)

	params := vod.CommitUploadInfoParam{
		SpaceName:    space,
		SessionKey:   session,
		Functions:    string(fbts),
		CallbackArgs: "",
	}
	resp, err := instance.CommitUploadInfo(params)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s", bts)
}
