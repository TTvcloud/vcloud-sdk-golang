package rtc

import (
	"encoding/json"
	"net/url"
)

func (p *Rtc) ByteStartTranscode(body map[string]interface{}) (*ByteStartTranscodeResp, int, error) {
	bts, err := json.Marshal(body)
	if err != nil {
		return nil, 400, err
	}

	resp := &ByteStartTranscodeResp{}
	status, err := p.commonPostJson("ByteStartTranscode", nil, string(bts), resp)
	return resp, status, err
}

func (p *Rtc) ByteStopTranscode(body map[string]interface{}) (*ByteStopTranscodeResp, int, error) {
	bts, err := json.Marshal(body)
	if err != nil {
		return nil, 400, err
	}

	resp := &ByteStopTranscodeResp{}
	status, err := p.commonPostJson("ByteStopTranscode", nil, string(bts), resp)
	return resp, status, err
}

func (p *Rtc) ByteTranscodeChangeLayout(body map[string]interface{}) (*ByteTranscodeChangeLayoutResp, int, error) {
	bts, err := json.Marshal(body)
	if err != nil {
		return nil, 400, err
	}

	resp := &ByteTranscodeChangeLayoutResp{}
	status, err := p.commonPostJson("ByteTranscodeChangeLayout", nil, string(bts), resp)
	return resp, status, err
}

func (p *Rtc) commonPostJson(api string, query url.Values, body string, out interface{}) (int, error) {
	respBody, status, err := p.Json(api, query, body)
	if err != nil {
		return status, err
	}

	err2 := json.Unmarshal(respBody, out)
	return status, err2
}
