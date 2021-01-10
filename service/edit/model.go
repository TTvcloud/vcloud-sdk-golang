package edit

import "github.com/TTvcloud/vcloud-sdk-golang/base"

type SubmitDirectEditTaskRequest struct {
	Uploader     string      `json:"Uploader,omitempty"`
	Application  string      `json:"Application,omitempty"`
	VideoName    string      `json:"VideoName,omitempty"`
	Param        interface{} `json:"EditParam"`
	Priority     int32       `json:"Priority"`
	CallbackUri  string      `json:"CallbackUri,omitempty"`
	CallbackArgs string      `json:"CallbackArgs,omitempty"`
}

type GetDirectEditResultRequest struct {
	ReqIds []string `json:"ReqIds"`
}

type SubmitDirectEditTaskAsyncResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           struct {
		ReqId string `json:"ReqId"`
	} `json:"Result"`
}

type SubmitDirectEditTaskSyncResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           struct {
		ReqId  string      `json:"ReqId"`
		Output interface{} `json:"Output"`
		Code   int         `json:"Code"`
		TaskId string      `json:"TaskId"`
	} `json:"Result"`
}

type GetDirectEditResultResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           []struct {
		ReqId        string      `json:"ReqId"`
		Param        interface{} `json:"EditParam"`
		Priority     int32       `json:"Priority"`
		CallbackUri  string      `json:"CallbackUri"`
		CallbackArgs string      `json:"CallbackArgs"`
		Status       string      `json:"Status"`
		OutputVid    string      `json:"OutputVid"`
		TaskId       string      `json:"TaskId"`
	} `json:"Result"`
}

type SubmitTemplateTaskAsyncResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           []string `json:"Result"`
}

type TemplateParamItem struct {
	Name     string  `json:"Name,omitempty"`
	Type     string  `json:"Type"`
	Position string  `json:"Position"`
	Text     *string `json:"Text,omitempty"`
	Source   *string `json:"Source,omitempty"`
}

type SubmitTemplateTaskAsyncRequest struct {
	TemplateId   string                 `json:"TemplateId" valid:"required"`
	Space        string                 `json:"Space" valid:"required"`
	VideoName    []string               `json:"VideoName" valid:"-"`
	Params       [][]*TemplateParamItem `json:"Params" valid:"required"`
	Priority     int                    `json:"Priority" valid:"-"`
	CallbackUri  string                 `json:"CallbackUri" valid:"-"`
	CallbackArgs string                 `json:"CallbackArgs" valid:"-"`
	Type         int                    `json:"Type" valid:"-"`
}
