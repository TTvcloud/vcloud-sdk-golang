package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/models"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	vid := "your vid"

	// GetPlayInfo

	instance := vod.NewInstance()
	query := &models.VodGetPlayInfoRequest{Vid: vid}
	resp, code, _ := instance.GetPlayInfo(query)
	fmt.Printf("resp:%+v code:%s\n", resp, code)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	// GetOriginVideoPlayInfo
	query2 := &models.VodGetOriginalPlayInfoRequest{Vid: vid}

	resp2, code2, _ := instance.GetOriginVideoPlayInfo(query2)
	fmt.Printf("resp:%+v code:%s\n", resp2, code2)
	fmt.Println(code)
	b2, _ := json.Marshal(resp2)
	fmt.Println(string(b2))

}
