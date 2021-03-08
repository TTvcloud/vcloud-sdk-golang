package imagex

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

const (
	ActionRefresh = 0
	ActionDisable = 1
	ActionEnable  = 2

	FunctionEncryption = "Encryption"
)

// GetImageThemeColor
type GetImageThemeColorResult struct {
	Color string `json:"color"`
}

// DeleteImageUploadFiles
type DeleteImageParam struct {
	StoreUris []string `json:"StoreUris"`
}

type DeleteImageResult struct {
	ServiceId    string   `json:"ServiceId"`
	DeletedFiles []string `json:"DeletedFiles"`
}

// ApplyImageUpload
type ApplyUploadImageParam struct {
	ServiceId  string
	SpaceName  string
	SessionKey string
	UploadNum  int
	StoreKeys  []string
	SkipMeta   bool
}

type ApplyUploadImageResult struct {
	RequestId   string      `json:"RequestId"`
	SessionKey  string      `json:"SessionKey"`
	UploadHosts []string    `json:"UploadHosts"`
	StoreInfos  []StoreInfo `json:"StoreInfos"`
}

type StoreInfo struct {
	StoreUri string `json:"StoreUri"`
	Auth     string `json:"Auth"`
}

// CommitImageUpload
type CommitUploadImageParam struct {
	ServiceId   string       `json:"-"`
	SpaceName   string       `json:"-"`
	SkipMeta    bool         `json:"-"`
	SessionKey  string       `json:"SessionKey"`
	SuccessOids []string     `json:"SuccessOids"`
	OptionInfos []OptionInfo `json:"OptionInfos"`
	Functions   []Function   `json:"Functions"`
}

type OptionInfo struct {
	StoreUri string `json:"StoreUri"`
	FileName string `json:"FileName"`
}

type Function struct {
	Name  string      `json:"Name"`
	Input interface{} `json:"Input"`
}

type EncryptionInput struct {
	Config       map[string]string `json:"Config"`
	PolicyParams map[string]string `json:"PolicyParams"`
}

type CommitUploadImageResult struct {
	Results    []Result    `json:"Results"`
	RequestId  string      `json:"RequestId"`
	ImageInfos []ImageInfo `json:"PluginResult"`
}

type Result struct {
	Uri        string     `json:"Uri"`
	UriStatus  int        `json:"UriStatus"`
	Encryption Encryption `json:"Encryption"`
}

type Encryption struct {
	Uri       string            `json:"Uri"`
	SecretKey string            `json:"SecretKey"`
	Algorithm string            `json:"Algorithm"`
	Version   string            `json:"Version"`
	SourceMd5 string            `json:"SourceMd5"`
	Extra     map[string]string `json:"Extra"`
}

type ImageInfo struct {
	FileName    string `json:"FileName"`
	ImageUri    string `json:"ImageUri"`
	ImageWidth  int    `json:"ImageWidth"`
	ImageHeight int    `json:"ImageHeight"`
	ImageMd5    string `json:"ImageMd5"`
	ImageFormat string `json:"ImageFormat"`
	ImageSize   int    `json:"ImageSize"`
	FrameCnt    int    `json:"FrameCnt"`
	Duration    int    `json:"Duration"`
}

// UpdateImageUploadFiles
type UpdateImageUrlPayload struct {
	Action    int      `json:"Action"`
	ImageUrls []string `json:"ImageUrls"`
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

//GetWeightsResp
type GetWeightsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           map[string]map[string]int `json:",omitempty"`
}

type DomainInfo struct {
	MainDomain   string
	BackupDomain string
}

type ImgUrl struct {
	MainUrl   string
	BackupUrl string
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
