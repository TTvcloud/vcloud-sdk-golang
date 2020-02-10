package live

import (
	"encoding/json"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

// GetDesensitizedAllAppInfos
type GetDesensitizedAllAppInfosResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *DesensitizedAllAppInfos `json:",omitempty"`
}

type DesensitizedAllAppInfos struct {
	Push2AppInfo      map[int64]*DesensitizedAppInfo
	Push2AllPlayInfos map[int64]map[int64]*DesensitizedAllPlayCdnAppInfo
}

type DesensitizedAppInfo struct {
	Id          int64
	Name        string
	Description string
	IsPlayFlv   bool
	IsPlayRtmp  bool
	IsPlayHls   bool
	IsPlayCmaf  bool
	IsPlayDash  bool
	Resolutions string
}

type DesensitizedAllPlayCdnAppInfo struct {
	PlayCdnApp *DesensitizedPlayCdnAppInfo
	Cdn        *DesensitizedCdnInfo
	Templates  []*DesensitizedTemplateInfo
}

type DesensitizedPlayCdnAppInfo struct {
	Id             int32
	PlayApp        string
	PlayProportion int32
}

type DesensitizedCdnInfo struct {
	Id             int32
	Name           string
	PlayFlvDomain  string
	PlayHlsDomain  string
	PlayRtmpDomain string
	PlayCmafDomain string
	PlayDashDomain string
	AdminHlsDomain string
	AdminFlvDomain string
}

type DesensitizedTemplateInfo struct {
	Id     int32
	Name   string
	Size   string
	Suffix string
}

// CreateStream
type CreateStreamRequest struct {
	AppID     int32
	Stream    string
	Extra     string
	DelayTime int64
}

type CreateStreamResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CreateStreamResult `json:",omitempty"`
}

type CreateStreamResult struct {
	AppID      int64
	Stream     string
	CreateTime int64
}

// MGetStreamsPushInfo
type MGetStreamsPushInfoRequest struct {
	Streams []string
}

type mGetStreamsPushInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *mGetStreamsPushInfoResult `json:",omitempty"`
}

type mGetStreamsPushInfoResult struct {
	PushInfos json.RawMessage
}

type MGetStreamsPushInfoResp struct {
	Result           map[string]*PushInfo `json:",omitempty"`
	ResponseMetadata *base.ResponseMetadata
}

type PushInfo struct {
	StreamBase *StreamBase
	Main       *ElePushInfo
	Backup     *ElePushInfo
	Suggest    *string
}

type StreamBase struct {
	AppId      int64
	Stream     string
	RefId      string
	Status     EStreamStatus
	Extra      string
	CreateTime int64
}

type StreamInfo struct {
	Status           EStreamStatus
	Resolutions      string
	PlayTypes        string
	PushMainCdnappId int64
	LiveId           string
	Appid            int64
	CreateTime       int64
	Description      string
}

type EStreamStatus int64

const (
	EStreamStatus_Unknown EStreamStatus = -1
	EStreamStatus_Create  EStreamStatus = 0
	EStreamStatus_Living  EStreamStatus = 1
	EStreamStatus_Stoped  EStreamStatus = 8
)

type ElePushInfo struct {
	Urls      []string
	VCodec    []string
	SdkParams *string
	RtmpUrl   string
}

// MGetStreamsPlayInfo
type MGetStreamsPlayInfoRequest struct {
	Streams            []string
	EnableSSL          bool
	IsCustomizedStream bool
}

type mGetStreamsPlayInfoResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *mGetStreamsPlayInfoResult `json:",omitempty"`
}

type mGetStreamsPlayInfoResult struct {
	PlayInfos json.RawMessage
}

type MGetStreamsPlayInfoResp struct {
	Result           map[string]*PlayInfo `json:",omitempty"`
	ResponseMetadata *base.ResponseMetadata
}

type PlayInfo struct {
	StreamBase          *StreamBase
	Main                []*ElePlayInfo
	Backup              []*ElePlayInfo
	Suggest             *string
	StreamData          *string
	StreamSizes         []string
	MainRecommendInfo   *Recommendation
	BackupRecommendInfo *Recommendation
}

type ElePlayInfo struct {
	VCodec    *string
	Size      *string
	Url       *PlayUrlInfo
	SdkParams *string
}

type PlayUrlInfo struct {
	FlvUrl  string
	HlsUrl  string
	RtmpUrl string
	CmafUrl string
	DashUrl string
}

type Recommendation struct {
	DefaultSize *string
}

// GetVODs
type GetVODsRequest struct {
	Stream string
}

type GetVODsResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetVODsResult `json:",omitempty"`
}

type GetVODsResult struct {
	VODs []*VOD
}

type VOD struct {
	SourceURL string
	VID       string
	Duration  float64
	StartTime int64
	EndTime   int64
}

// GetRecords
type GetRecordsRequest struct {
	Stream string
}

type GetRecordsResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetRecordsResult `json:",omitempty"`
}

type GetRecordsResult struct {
	Records []*Record
}

type Record struct {
	URL       string
	Type      string
	Duration  float64
	StartTime int64
	EndTime   int64
}

// GetSnapshots
type GetSnapshotsRequest struct {
	Stream string
}

type GetSnapshotsResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetSnapshotsResult `json:",omitempty"`
}

type GetSnapshotsResult struct {
	Snapshots []*Snapshot
}

type Snapshot struct {
	URL       string
	Timestamp int64
}

// GetStreamTimeShiftInfo
type GetStreamTimeShiftInfoRequest struct {
	Stream    string
	StartTime int64
	EndTime   int64
}

type GetStreamTimeShiftInfoResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetStreamTimeShiftInfoResult `json:",omitempty"`
}

type GetStreamTimeShiftInfoResult struct {
	URL       string
	StartTime int64
	EndTime   int64
	VCodec    string
}

// GetOnlineUserNum
type GetOnlineUserNumRequest struct {
	Stream    string
	StartTime int64
	EndTime   int64
}

type GetOnlineUserNumResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *GetOnlineUserNumResult `json:",omitempty"`
}

type GetOnlineUserNumResult struct {
	OnlineUserNum []OnlineUserNum
}

type OnlineUserNum struct {
	Timestamp int64
	Num       int64
}

// fallback model
type mGetStreamsPlayContext struct {
	streams []string

	streamInfos    map[string]*StreamInfo
	scheduleResult map[string]*schedulePlayResult
	enableSSL      bool
}

type schedulePlayResult struct {
	streamInfo         *StreamInfo
	mainScheduleResult *scheduleElePlayResult
}

type scheduleElePlayResult struct {
	playCdnApp *DesensitizedPlayCdnAppInfo
	cdn        *DesensitizedCdnInfo
	templates  []*DesensitizedTemplateInfo
}

type deserializePlayParams struct {
	enableSSL     bool
	expireSeconds int64

	stream     string
	streamInfo *StreamInfo

	scheduleResult *scheduleElePlayResult
}

type genElePlayParams struct {
	playTypes []string
	size      string
	sdkParams string

	streamInfo   *StreamInfo
	PlayCdnApp   *DesensitizedPlayCdnAppInfo
	Cdn          *DesensitizedCdnInfo
	templateInfo *DesensitizedTemplateInfo

	enableSSL bool
}

// ForbidStream
type ForbidStreamRequest struct {
	AppID          int64
	Stream         string
	ForbidInterval int64
}

type ForbidStreamResponse struct {
	ResponseMetadata *base.ResponseMetadata
}
