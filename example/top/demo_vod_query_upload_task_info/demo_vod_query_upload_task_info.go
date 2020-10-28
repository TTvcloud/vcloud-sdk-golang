package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "AKLTNDQ2YTRlNTBiYTg1NDcyNmE3MDA1MTUzNzc5MWMwNmI",
		SecretAccessKey: "1ZOtyBZ89VERZdOfiUrPf24a3tTjRo1XIJbzccVHMrBvZo1jEn60LjClP2t05qWz",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	jobIds := make([]string, 0)
	jobId := "45be782c3c2042e6abd22cf172a61f0a"
	jobIds = append(jobIds, jobId)
	str := strings.Join(jobIds, ",")

	params := vod.UrlQueryParams{JobIds: str}
	resp, err := instance.QueryUploadTaskInfo(params)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	bts, _ := json.Marshal(resp)
	fmt.Printf("resp = %s", bts)
}
