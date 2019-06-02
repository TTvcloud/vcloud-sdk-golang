package vod

import "github.com/TTvcloud/vcloud-sdk-golang/base"

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
