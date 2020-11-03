package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/request"
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

	query := &request.VodGetPlayInfoRequest{
		Vid:        vid,
		Format:     "",
		Codec:      "",
		Definition: "360p",
		FileType:   "",
		LogoType:   "",
		Base64:     "1",
		Ssl:        "",
	}
	resp, code, _ := instance.GetPlayInfo(query)
	fmt.Printf("resp:%+v code:%s\n", resp, code)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))


	// GetOriginalPlayInfo
	query2 := &request.VodGetOriginalPlayInfoRequest{Vid: vid}

	resp2, code2, _ := instance.GetOriginalPlayInfo(query2)
	fmt.Printf("resp:%+v code:%s\n", resp2, code2)
	fmt.Println(code)
	b2, _ := json.Marshal(resp2)
	fmt.Println(string(b2))

}
