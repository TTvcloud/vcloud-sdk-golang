package imagex

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
)

// ApplyImageUpload 获取图片上传地址
func (c *ImageXClient) ApplyUploadImage(params *ApplyUploadImageParam) (*ApplyUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	if params.SpaceName != "" {
		query.Add("SpaceName", params.SpaceName)
	}
	if params.SessionKey != "" {
		query.Add("SessionKey", params.SessionKey)
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", strconv.Itoa(params.UploadNum))
	}
	for _, key := range params.StoreKeys {
		query.Add("StoreKeys", key)
	}

	respBody, _, err := c.Query("ApplyImageUpload", query)
	if err != nil {
		return nil, fmt.Errorf("fail to request api ApplyImageUpload, %v", err)
	}

	result := new(struct {
		UploadAddress ApplyUploadImageResult `json:"UploadAddress"`
		RequestId     string                 `json:"RequestId"`
	})
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	result.UploadAddress.RequestId = result.RequestId
	return &result.UploadAddress, nil
}

// CommitImageUpload 图片上传完成上报
func (c *ImageXClient) CommitUploadImage(params *CommitUploadImageParam) (*CommitUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	if params.SpaceName != "" {
		query.Add("SpaceName", params.SpaceName)
	}

	bts, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal request, %v", err)
	}

	respBody, _, err := c.Json("CommitImageUpload", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("fail to request api CommitImageUpload, %v", err)
	}

	result := new(CommitUploadImageResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageXClient) upload(host string, storeInfo StoreInfo, imageBytes []byte) error {
	if len(imageBytes) == 0 {
		return fmt.Errorf("file size is zero")
	}

	checkSum := fmt.Sprintf("%x", crc32.ChecksumIEEE(imageBytes))
	url := fmt.Sprintf("http://%s/%s", host, storeInfo.StoreUri)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(imageBytes))
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

// 上传图片
func (c *ImageXClient) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
	if params.UploadNum == 0 {
		params.UploadNum = 1
	}
	if len(images) != params.UploadNum {
		return nil, fmt.Errorf("images num %d != upload num %d", len(images), params.UploadNum)
	}

	// 1. apply
	applyResp, err := c.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	} else if len(applyResp.UploadHosts) == 0 {
		return nil, fmt.Errorf("no upload host found")
	} else if len(applyResp.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("store infos num %d != upload num %d", len(applyResp.StoreInfos), params.UploadNum)
	}

	// 2. upload
	for i, image := range images {
		err := c.upload(applyResp.UploadHosts[0], applyResp.StoreInfos[i], image)
		if err != nil {
			return nil, err
		}
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:  params.ServiceId,
		SpaceName:  params.SpaceName,
		SessionKey: applyResp.SessionKey,
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

func (c *ImageXClient) GetUploadAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}

	applyUploadToken, err := c.GetSignUrl("ApplyImageUpload", query)
	if err != nil {
		return "", err
	}
	ret["ApplyUploadToken"] = applyUploadToken

	commitUploadToken, err := c.GetSignUrl("CommitImageUpload", query)
	if err != nil {
		return "", err
	}
	ret["CommitUploadToken"] = commitUploadToken

	b, err := json.Marshal(ret)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
