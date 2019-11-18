package imagex

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

// ApplyUploadImageFile
type ApplyUploadImageParam struct {
	ServiceId  string
	SessionKey string
	UploadNum  int
	StoreKeys  []string
}

type ApplyUploadImageResult struct {
	ServiceId   string      `json:"ServiceId"`
	SessionKey  string      `json:"SessionKey"`
	UploadHosts []string    `json:"UploadHosts"`
	StoreInfos  []StoreInfo `json:"StoreInfos"`
}

type StoreInfo struct {
	StoreUri string `json:"StoreUri"`
	Auth     string `json:"Auth"`
}

// CommitUploadImageFile
type CommitUploadImageParam struct {
	ServiceId   string
	SessionKey  string
	OptionInfos []OptionInfo `json:"OptionInfos"`
}

type OptionInfo struct {
	StoreUri string `json:"StoreUri"`
	FileName string `json:"FileName"`
}

type CommitUploadImageResult struct {
	ServiceId  string      `json:"ServiceId"`
	ImageInfos []ImageInfo `json:"ImageInfos"`
}

type ImageInfo struct {
	FileName    string `json:"FileName"`
	ImageUri    string `json:"ImageUri"`
	ImageWidth  int    `json:"ImageWidth"`
	ImageHeight int    `json:"ImageHeight"`
	ImageMd5    string `json:"ImageMd5"`
}

// GetImageTemplateConf
type GetTemplateConfParam struct {
	GroupId   string `json:"GroupId"`   // 可选参数，若指定则返回该分组的相关配置信息
	GroupName string `json:"GroupName"` // 可选参数，若指定则返回名称中包含该值的所有分组的相关配置信息
}

type GetTemplateConfResult struct {
	Groups []Group `json:"Groups"`
}

type Group struct {
	Name    string `json:"Name"`
	GroupId string `json:"GroupId"`
	Confs   []Conf `json:"Confs"`
}

type Conf struct {
	ConfId           string   `json:"ConfId"`
	ServiceId        string   `json:"ServiceId"`
	ServiceName      string   `json:"ServiceName"`
	TemplateName     string   `json:"TemplateName"`
	TemplateAbstract string   `json:"TemplateAbstract"`
	TemplateParams   []string `json:"TemplateParams"`
	Weight           int      `json:"Weight"`
	ImageFormat      string   `json:"ImageFormat"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}
