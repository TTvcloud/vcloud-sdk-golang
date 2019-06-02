package vod

import (
	"encoding/json"
	"net/url"
)

func (p *Vod) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Vod) GetPlayInfo(query url.Values) (*GetPlayInfoResp, int, error) {
	respBody, status, err := p.Query("GetPlayInfo", query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetPlayInfoResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "vod"
		return output, status, nil
	}
}
