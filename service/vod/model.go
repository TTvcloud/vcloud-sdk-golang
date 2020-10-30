package vod

import (
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod/top"
)

type RedirectPlayReq struct {
	Vid        string `json:"Vid"`
	Definition string `json:"Definition"` //视频分辨率
	LogoType   string `json:"LogoType,omitempty"`
	Expires    string
}

type StartWorkflowRequest struct {
	Vid          string
	TemplateId   string
	Input        map[string]interface{}
	Priority     int
	CallbackArgs string
}

type StartWorkflowResult struct {
	RunId string
}

type StartWorkflowResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *StartWorkflowResult `json:",omitempty"`
}

type UploadMediaByUrlResult struct {
	Code    int
	Message string
}

type UploadMediaByUrlResp struct {
	base.CommonResponse
	Result UploadMediaByUrlResult
}

type UploadVideoByUrlResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.UrlUploadResponse  `json:"Result,omitempty"`
}

type QueryUploadTaskInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.UrlQueryData       `json:"Result,omitempty"`
}

type VideoFormat string

const (
	MP4  VideoFormat = "mp4"
	M3U8 VideoFormat = "m3u8"
)

type UploadMediaByUrlParams struct {
	SpaceName    string
	Format       VideoFormat
	SourceUrls   []string
	CallbackArgs string
}

type UrlUploadParams struct {
	SpaceName string   `json:"SpaceName"`
	URLSets   []URLSet `json:"URLSets"`
}

type URLSet struct {
	SourceUrl    string `json:"SourceUrl"`
	CallbackArgs string `json:"CallbackArgs"`
	Md5          string `json:"Md5"`
	TemplateId   string `json:"TemplateId"`
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	Tags         string `json:"Tags"`
	Category     string `json:"Category"`
}

type UrlQueryParams struct {
	JobIds string `json:"JobIds"`
}

type FileType string

const (
	VIDEO  FileType = "video"
	IMAGE  FileType = "image"
	OBJECT FileType = "object"
)

type ApplyUploadParam struct {
	SpaceName  string
	SessionKey string
	FileType   FileType
	FileSize   int
	UploadNum  int
}

type ApplyUploadResp struct {
	base.CommonResponse
	Result ApplyUploadResult
}

type ApplyUploadResult struct {
	RequestID     string
	UploadAddress UploadAddress
}
type UploadAddress struct {
	StoreInfos    []StoreInfo
	UploadHosts   []string
	UploadHeader  map[string]string
	SessionKey    string
	AdvanceOption AdvanceOption
}

type ApplyUploadInfoParam struct {
	SpaceName  string
	SessionKey string
	FileSize   int
}

type CommitUploadInfoParam struct {
	SpaceName    string
	SessionKey   string
	CallbackArgs string
	Functions    string
}

type ApplyUploadInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.ApplyData          `json:"Result,omitempty"`
}

type CommitUploadInfoResp struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           top.CommitData         `json:"Result,omitempty"`
}

type VideoDefinition string

const (
	D1080P VideoDefinition = "1080p"
	D720P  VideoDefinition = "720p"
	D540P  VideoDefinition = "540p"
	D480P  VideoDefinition = "480p"
	D360P  VideoDefinition = "360p"
	D240P  VideoDefinition = "240p"
)

type StoreInfo struct {
	StoreUri string
	Auth     string
}

type AdvanceOption struct {
	Parallel  int
	Stream    int
	SliceSize int
}

type ModifyVideoInfoBody struct {
	SpaceName string       `json:"SpaceName"`
	Vid       string       `json:"Vid"`
	Info      UserMetaInfo `json:"Info"`
	Tags      TagControl   `json:"Tags"`
}

type UserMetaInfo struct {
	Title       string
	Description string
	Category    string
	PosterUri   string
}

type TagControl struct {
	Deletes string
	Adds    string
}

type ModifyVideoInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ModifyVideoInfoBaseResp
}

type ModifyVideoInfoBaseResp struct {
	BaseResp *BaseResp
}

type BaseResp struct {
	StatusMessage string
	StatusCode    int
}

type CommitUploadParam struct {
	SpaceName string
	Body      CommitUploadBody
}

type CommitUploadBody struct {
	CallbackArgs string
	SessionKey   string
	Functions    []Function
}

type Function struct {
	Name  string      `json:"Name"`
	Input interface{} `json:"Input,,omitempty"`
}

type SnapshotInput struct {
	SnapshotTime float64
}

type EntryptionInput struct {
	Config       map[string]string
	PolicyParams map[string]string
}

type OptionInfo struct {
	Title       string
	Tags        string
	Description string
	Category    string
}

type WorkflowInput struct {
	TemplateId string
}

type CommitUploadResp struct {
	base.CommonResponse
	Result CommitUploadResult
}

type CommitUploadResult struct {
	RequestId    string
	CallbackArgs string
	Results      []UploadResult
}

type UploadResult struct {
	Vid         string
	VideoMeta   VideoMeta
	ImageMeta   ImageMeta
	ObjectMeta  ObjectMeta
	Encryption  Encryption
	SnapshotUri string
}
type VideoMeta struct {
	Uri      string
	Height   int
	Width    int
	Duration float64
	Bitrate  int
	Md5      string
	Format   string
	Size     int
}

type ImageMeta struct {
	Uri    string
	Height int
	Width  int
	Md5    string
}

type ObjectMeta struct {
	Uri string
	Md5 string
}

type Encryption struct {
	Uri       string
	SecretKey string
	Algorithm string
	Version   string
	SourceMd5 string
	Extra     map[string]string
}

// SetVideoPublishStatus
type SetVideoPublishStatusResp struct {
	ResponseMetadata *base.ResponseMetadata
}

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
