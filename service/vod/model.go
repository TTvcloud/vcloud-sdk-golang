package vod

import (
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"time"
)

// GetPlayInfo
type GetPlayInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetPlayInfoData `json:",omitempty"`
}

type GetPlayInfoData struct {
	Data *GetPlayInfoInner
}

type GetPlayInfoInner struct {
	Status       int
	VideoID      string
	CoverUrl     string
	Duration     float32
	MediaType    string
	PlayInfoList []*PlayInfo
	TotalCount   int
}

type PlayInfo struct {
	Bitrate         int
	FileHash        string
	Size            int
	Height          int
	Width           int
	Format          string
	Codec           string
	Logo            string
	Definition      string
	Quality         string
	PlayAuth        string
	MainPlayUrl     string
	BackupPlayUrl   string
	FileID          string
	P2pVerifyURL    string
	PreloadInterval int
	PreloadMaxStep  int
	PreloadMinStep  int
	PreloadSize     int
}

type StartTranscodeRequest struct {
	Vid        string
	TemplateId string
	Input      map[string]interface{}
	Priority   int
}

type StartTranscodeResult struct {
	RunId string
}

type StartTranscodeResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *StartTranscodeResult `json:",omitempty"`
}

type UploadVideoByUrlResult struct {
	Code    int
	Message string
}

type UploadVideoByUrlResp struct {
	base.CommonResponse
	Result UploadVideoByUrlResult
}

type VideoFormat string

const (
	MP4  VideoFormat = "mp4"
	M3U8 VideoFormat = "m3u8"
)

type UploadVideoByUrlParams struct {
	SpaceName    string
	Format       VideoFormat
	SourceUrls   []string
	CallbackArgs string
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

type VideoDefinition string

const (
	D1080P VideoDefinition = "1080p"
	D720P  VideoDefinition = "720p"
	D540P  VideoDefinition = "540p"
	D480P  VideoDefinition = "480p"
	D360P  VideoDefinition = "360p"
	D240P  VideoDefinition = "240p"
)

type RedirectPlayParam struct {
	VideoID    string
	Definition VideoDefinition
	Watermark  string
	Expire     time.Duration
}

type StoreInfo struct {
	StoreUri string
	Auth     string
}

type AdvanceOption struct {
	Parallel  int
	Stream    int
	SliceSize int
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
	Name  string
	Input interface{}
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
	RequestId string
	Results   []UploadResult
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
