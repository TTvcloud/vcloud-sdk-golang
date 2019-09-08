package imagex

import "github.com/TTvcloud/vcloud-sdk-golang/base"

// ApplyUploadImageFile
type ApplyUploadImageParam struct {
	ServiceId  string
	SessionKey string
	UploadNum  int
	StoreKeys  []string
}

type ApplyUploadImageResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ApplyUploadImageResult `json:",omitempty"`
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

type CommitUploadImageResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CommitUploadImageResult `json:",omitempty"`
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
