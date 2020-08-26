package main

import (
	"encoding/json"
	"fmt"
	"net/url"

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
	query := vod.GetPlayInfoReq{Vid: vid}
	resp, code, _ := instance.GetPlayInfo(query)
	fmt.Printf("resp:%+v code:%d\n", resp, code)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	// GetOriginVideoPlayInfo
	query2 := url.Values{}
	query2.Set("Vid", vid)

	resp2, code, _ := instance.GetOriginVideoPlayInfo(query2)
	fmt.Printf("resp:%+v code:%d\n", resp2, code)
	fmt.Println(code)
	b2, _ := json.Marshal(resp2)
	fmt.Println(string(b2))

	// GetRedirectPlayUrl
	params := vod.RedirectPlayParam{
		Vid:        vid,
		Definition: vod.D1080P,
		Watermark:  "",
		// set expires time of the redirect play url, defalut is 15min(900),
		// set if if you know the params' meaning exactly.
		Expires: "60",
	}
	ret, err := vod.NewInstanceWithRegion("us-east-1").GetRedirectPlayUrl(params)
	fmt.Println(ret, err)
}
