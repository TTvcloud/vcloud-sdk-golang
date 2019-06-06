package vod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
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

func (p *Vod) UploadVideoByUrl(params UploadVideoByUrlParams) (*UploadVideoByUrlResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)
	query.Add("Format", string(params.Format))
	query.Add("SourceUrls", strings.Join(params.SourceUrls, ","))
	query.Add("Extra", params.Extra)
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

	resp, err := DefaultInstance.ApplyUpload(params)
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
	commitResp, err := DefaultInstance.CommitUpload(param)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}
