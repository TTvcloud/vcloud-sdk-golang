package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	vod.NewInstance().SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	StartTranscodeExample()
}

func StartTranscodeExample() {
	input := map[string]interface{}{
		"watermark_str": "test",
	}

	req := &vod.StartTranscodeRequest{
		Vid:        "your vid",
		TemplateId: "your template id",
		Input:      input,
		Priority:   0,
	}

	resp, err := vod.NewInstance().StartTranscode(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	fmt.Println(resp.Result)
}
