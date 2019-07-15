package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//GetPlayInfo 获取播放信息
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

//GetOriginVideoPlayInfo 获取原片播放信息
func (p *Vod) GetOriginVideoPlayInfo(query url.Values) (*GetOriginVideoPlayInfoResp, int, error) {
	respBody, status, err := p.Query("GetOriginVideoPlayInfo", query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetOriginVideoPlayInfoResp)
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

func (p *Vod) UploadVideoByUrl(params UploadVideoByUrlParams) (*UploadVideoByUrlResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)
	query.Add("Format", string(params.Format))
	query.Add("SourceUrls", strings.Join(params.SourceUrls, ","))
	query.Add("CallbackArgs", params.CallbackArgs)
	respBody, status, err := p.Query("UploadMediaByUrl", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &UploadVideoByUrlResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) ApplyUpload(params ApplyUploadParam) (*ApplyUploadResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)
	query.Add("SessionKey", params.SessionKey)
	query.Add("FileType", string(params.FileType))
	if params.FileSize != 0 {
		query.Add("FileSize", string(params.FileSize))
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", string(params.UploadNum))
	}

	respBody, status, err := p.Query("ApplyUpload", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &ApplyUploadResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) CommitUpload(params CommitUploadParam) (*CommitUploadResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)

	bts, err := json.Marshal(params.Body)
	if err != nil {
		return nil, err
	}

	respBody, status, err := p.Json("CommitUpload", query, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &CommitUploadResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) Upload(fileBytes []byte, spaceName string, fileType FileType, funcs ...Function) (*CommitUploadResp, error) {
	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("file size is zero")
	}

	params := ApplyUploadParam{
		SpaceName: spaceName,
		FileType:  fileType,
	}

	resp, err := p.ApplyUpload(params)
	if err != nil {
		return nil, err
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return nil, fmt.Errorf("%+v", resp.ResponseMetadata.Error)
	}

	if len(resp.Result.UploadAddress.UploadHosts) == 0 {
		return nil, fmt.Errorf("no tos host found")
	}
	if len(resp.Result.UploadAddress.StoreInfos) == 0 {
		return nil, fmt.Errorf("no store infos found")
	}

	// upload file
	checkSum := fmt.Sprintf("%x", crc32.ChecksumIEEE(fileBytes))
	tosHost := resp.Result.UploadAddress.UploadHosts[0]
	oid := resp.Result.UploadAddress.StoreInfos[0].StoreUri
	auth := resp.Result.UploadAddress.StoreInfos[0].Auth
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("http status=%v, body=%s, remote_addr=%v", rsp.StatusCode, string(b), req.Host)
	}
	defer rsp.Body.Close()

	var tosResp struct {
		Success int         `json:"success"`
		Payload interface{} `json:"payload"`
	}
	err = json.NewDecoder(rsp.Body).Decode(&tosResp)
	if err != nil {
		return nil, err
	}

	if tosResp.Success != 0 {
		return nil, fmt.Errorf("tos err:%+v", tosResp)
	}

	param := CommitUploadParam{
		SpaceName: spaceName,
		Body: CommitUploadBody{
			CallbackArgs: "",
			SessionKey:   resp.Result.UploadAddress.SessionKey,
			Functions:    funcs,
		},
	}
	commitResp, err := p.CommitUpload(param)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

// SetVideoPublishStatus 媒资相关
func (p *Vod) SetVideoPublishStatus(SpaceName, Vid, Status string) (*SetVideoPublishStatusResp, int, error) {
	jsonTemp := `{
		"SpaceName" : "%v",
		"Vid" : "%v",
		"Status" : "%v"
	}`

	body := fmt.Sprintf(jsonTemp, SpaceName, Vid, Status)
	respBody, status, err := p.Json("SetVideoPublishStatus", nil, body)
	if err != nil {
		return nil, status, err
	}

	output := new(SetVideoPublishStatusResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	}
	output.ResponseMetadata.Service = "vod"
	return output, status, nil
}

func (p *Vod) GetPlayAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}
	if getPlayInfoToken, err := p.GetSignUrl("GetPlayInfo", query); err == nil {
		ret["GetPlayInfoToken"] = getPlayInfoToken
	} else {
		return "", err
	}

	b, _ := json.Marshal(ret)
	return base64.StdEncoding.EncodeToString(b), nil
}

//GetRedirectPlayUrl get redirected playback addres
func (p *Vod) GetRedirectPlayUrl(params RedirectPlayParam) (string, error) {
	query := url.Values{}
	query.Add("video_id", params.VideoID)
	query.Add("expire", strconv.FormatInt(time.Now().Add(params.Expire).Unix(), 10))
	if params.Definition == "" {
		return "", errors.New("Defintion not set")
	}
	query.Add("definition", string(params.Definition))

	token, err := p.GetSignUrl("RedirectPlay", query)
	if err != nil {
		return "", err
	}

	apiInfo := p.ApiInfoList["RedirectPlay"]
	url := fmt.Sprintf("http://%s%s?%s", p.ServiceInfo.Host, apiInfo.Path, token)
	return url, nil
}

func (p *Vod) GetUploadAuthToken(space string) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}
	query := url.Values{}
	query.Set("SpaceName", space)

	if applyUploadToken, err := p.GetSignUrl("ApplyUpload", query); err == nil {
		ret["ApplyUploadToken"] = applyUploadToken
	} else {
		return "", err
	}

	if commitUploadToken, err := p.GetSignUrl("CommitUpload", query); err == nil {
		ret["CommitUploadToken"] = commitUploadToken
	} else {
		return "", err
	}

	b, _ := json.Marshal(ret)
	return base64.StdEncoding.EncodeToString(b), nil
}

func (p *Vod) GetCdnDomainWeights(spaceName string) (*GetWeightsResp, error) {
	query := url.Values{}
	query.Set("SpaceName", spaceName)
	respBody, _, err := p.Query("GetCdnDomainWeights", query)
	if err != nil {
		return nil, err
	}

	output := new(GetWeightsResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, err
	}
	return output, nil
}
