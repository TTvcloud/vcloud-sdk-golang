package main

import (
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	StartWorkflowExample()
}

func StartWorkflowExample() {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	input := map[string]interface{}{
		"watermark_str": "test",
	}

	req := &vod.StartWorkflowRequest{
		Vid:          "your vid",
		TemplateId:   "your template id",
		Input:        input,
		Priority:     0,
		CallbackArgs: "",
	}

	resp, err := instance.StartWorkflow(req)
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
