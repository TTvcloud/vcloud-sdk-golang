package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "you ak",
		SecretAccessKey: "you sk"})

	query := url.Values{}
	query.Set("Vid", "you vid")

	resp, code, _ := vod.DefaultInstance.GetOriginVideoPlayInfo(query)
	fmt.Printf("resp:%+v code:%d\n", resp, code)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
