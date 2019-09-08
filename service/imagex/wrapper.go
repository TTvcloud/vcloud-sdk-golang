package imagex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// ApplyUploadImageFile 获取文件上传地址
func (x *ImageX) ApplyUploadImage(params *ApplyUploadImageParam) (*ApplyUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	if params.SessionKey != "" {
		query.Add("SessionKey", params.SessionKey)
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", strconv.Itoa(params.UploadNum))
	}
	for _, key := range params.StoreKeys {
		query.Add("StoreKeys", key)
	}

	respBody, status, err := x.Query("ApplyUploadImageFile", query)
	if err != nil {
		return nil, fmt.Errorf("fail to request api ApplyUploadImageFile, %v", err)
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := new(ApplyUploadImageResp)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, fmt.Errorf("fail to unmarshal response, %v", err)
	} else if err := resp.ResponseMetadata.Error; err != nil && err.CodeN != 0 {
		return nil, fmt.Errorf("apply upload image request %s error %s",
			resp.ResponseMetadata.RequestId, err.Message)
	}
	return resp.Result, nil
}

// CommitUploadImageFile 文件上传完成上报
func (x *ImageX) CommitUploadImage(params *CommitUploadImageParam) (*CommitUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	query.Add("SessionKey", params.SessionKey)

	bts, err := json.Marshal(params.OptionInfos)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal request, %v", err)
	}

	respBody, status, err := x.Json("CommitUploadImageFile", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("fail to request api CommitUploadImageFile, %v", err)
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := new(CommitUploadImageResp)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, fmt.Errorf("fail to unmarshal response, %v", err)
	} else if err := resp.ResponseMetadata.Error; err != nil && err.CodeN != 0 {
		return nil, fmt.Errorf("commit upload image request %s error %s",
			resp.ResponseMetadata.RequestId, err.Message)
	}
	return resp.Result, nil
}

func (x *ImageX) Upload(host string, storeInfo StoreInfo, imageBytes []byte) error {
	if len(imageBytes) == 0 {
		return fmt.Errorf("file size is zero")
	}

	checkSum := fmt.Sprintf("%x", crc32.ChecksumIEEE(imageBytes))
	url := fmt.Sprintf("http://%s/%s", host, storeInfo.StoreUri)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(imageBytes))
	if err != nil {
		return fmt.Errorf("fail to new put request, %v", err)
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", storeInfo.Auth)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to do request, %v", err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("fail to read response body, %v", err)
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status=%v, body=%s, remote_addr=%v", rsp.StatusCode, string(body), req.Host)
	}
	defer rsp.Body.Close()

	var putResp struct {
		Success int         `json:"success"`
		Payload interface{} `json:"payload"`
	}
	if err = json.Unmarshal(body, &putResp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	if putResp.Success != 0 {
		return fmt.Errorf("put to host %s err:%+v", host, putResp)
	}
	return nil
}

func (x *ImageX) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
	if params.UploadNum == 0 {
		params.UploadNum = 1
	}
	if len(images) != params.UploadNum {
		return nil, fmt.Errorf("images num %d != upload num %d", len(images), params.UploadNum)
	}

	// 1. apply
	applyResp, err := x.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	} else if len(applyResp.UploadHosts) == 0 {
		return nil, fmt.Errorf("no upload host found")
	} else if len(applyResp.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("store infos num %d != upload num %d", len(applyResp.StoreInfos), params.UploadNum)
	}

	// 2. upload
	for i, image := range images {
		err := x.Upload(applyResp.UploadHosts[0], applyResp.StoreInfos[i], image)
		if err != nil {
			return nil, err
		}
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:  params.ServiceId,
		SessionKey: applyResp.SessionKey,
	}
	commitResp, err := x.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}
