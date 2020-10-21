package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
	"time"
)

func main() {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "AKLTZDI1NTZkYWI2MzEwNGI5MGE1MmRjNGJmYzg2MmQyYmE",
		SecretAccessKey: "D+5K+SOYf+L232Se+h4yRbhZu/P7pVeti9QNF138R4zSVFWeqtClX4XAdgcGplt+",
	})
	ret, _ := instance.GetUploadAuth()
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetUploadAuthWithExpiredTime(time.Minute)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
