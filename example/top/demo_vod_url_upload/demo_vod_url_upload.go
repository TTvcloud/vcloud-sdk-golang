package main

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod/top/models"
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

	spaceName := "james-test"
	urlSets := make([]*models.URLSet, 0)
	urlSet := &models.URLSet{
		SourceUrl: "https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4",
	}
	urlSets = append(urlSets, urlSet)

	urlRequest := models.VodUrlUploadRequest{
		SpaceName: spaceName,
		URLSets:   urlSets,
	}

	resp, err := instance.UploadVideoByUrl(urlRequest)
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
