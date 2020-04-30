package rtc

import "github.com/TTvcloud/vcloud-sdk-golang/base"

// ByteStartTranscodeResp
type ByteStartTranscodeResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CommonResult `json:",omitempty"`
}

// ByteStopTranscodeResp
type ByteStopTranscodeResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CommonResult `json:",omitempty"`
}

//ByteTranscodeChangeLayoutResp
type ByteTranscodeChangeLayoutResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CommonResult `json:",omitempty"`
}

// CommonResult
type CommonResult struct {
	Message string
}
