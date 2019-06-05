package vod

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

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

func (p *Vod) StartTranscode(req *StartTranscodeRequest) (*StartTranscodeResp, error) {
	query := url.Values{
		"TemplateId": []string{req.TemplateId},
	}

	reqBody := struct {
		Vid      string
		Input    map[string]interface{}
		Priority int
	}{
		Vid:      req.Vid,
		Input:    req.Input,
		Priority: req.Priority,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal body failed")
	}

	respBody, status, err := p.Json("StartTranscode", query, string(body))
	if err != nil || status != http.StatusOK {
		return nil, errors.Wrap(err, "query error")
	}

	resp := new(StartTranscodeResp)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal body failed")
	}

	return resp, nil
}
