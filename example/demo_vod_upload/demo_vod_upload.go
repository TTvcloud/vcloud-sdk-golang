package main

import (
	"fmt"
	"io/ioutil"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

const spaceName = "your-space-name"

func main() {
	uploadAll()
	uploadVideoByUrl()
}

func uploadVideoByUrl() {
	params := vod.UploadVideoByUrlParams{
		SpaceName:  spaceName,
		Format:     vod.MP4,
		SourceUrls: []string{"video-url"},
		Extra:      "xxx",
	}
	resp, err := vod.DefaultInstance.UploadVideoByUrl(params)
	if err != nil {
		fmt.Printf("err:%s\n")
	}
	if resp.ResponseMetadata.Error != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}
	fmt.Println(resp.Result)
}

func uploadAll() {
	snapShotFunc := vod.Function{Name: "Snapshot", Input: vod.SnapshotInput{SnapshotTime: 2.3}}
	getMetaFunc := vod.Function{Name: "GetMeta"}

	resp, err := upload(spaceName, "path-to-video", vod.VIDEO, getMetaFunc, snapShotFunc)
	fmt.Printf("resp:%+v err:%s", resp, err)
	resp, err = upload(spaceName, "path-to-img", vod.IMAGE, snapShotFunc)
	fmt.Printf("resp:%+v err:%s", resp, err)
	resp, err = upload(spaceName, "path-to-obj", vod.OBJECT, vod.Function{Name: "GetMeta"}, snapShotFunc)
	fmt.Printf("resp:%+v err:%s", resp, err)
}

func upload(spaceName string, filePath string, fileType vod.FileType, funcs ...vod.Function) (*vod.CommitUploadResp, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	rsp, err := vod.DefaultInstance.Upload(dat, spaceName, fileType, funcs...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
