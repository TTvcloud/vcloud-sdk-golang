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
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

// DeleteImageUploadFiles 删除图片
func (c *ImageXClient) DeleteImages(serviceId string, uris []string) (*DeleteImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)
	param := new(DeleteImageParam)
	param.StoreUris = uris

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal request, %v", err)
	}

	data, _, err := c.Json("DeleteImageUploadFiles", query, string(body))
	if err != nil {
		return nil, fmt.Errorf("fail to request api DeleteImageUploadFiles, %v", err)
	}

	result := new(DeleteImageResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, err
	}
	return result, nil
}

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
		return nil, fmt.Errorf("fail to request api ApplyImageUpload, %s, %v", string(respBody), err)
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
	query.Add("SkipMeta", fmt.Sprintf("%v", params.SkipMeta))

	bts, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal request, %v", err)
	}

	respBody, _, err := c.Json("CommitImageUpload", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("fail to request api CommitImageUpload, %s, %v", string(respBody), err)
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
		return fmt.Errorf("http status=%v, body=%s, url=%s", rsp.StatusCode, string(body), url)
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
		return fmt.Errorf("put to url %s err:%+v", url, putResp)
	}
	return nil
}

// 上传图片
func (c *ImageXClient) UploadImages(params *ApplyUploadImageParam, images [][]byte, functions ...Function) (*CommitUploadImageResult, error) {
	params.UploadNum = len(images)

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
	success := make([]string, 0)
	host := applyResp.UploadHosts[0]
	for i, image := range images {
		info := applyResp.StoreInfos[i]
		for n := 0; n < 3; n++ {
			err := c.upload(host, info, image)
			if err != nil {
				fmt.Printf("Fail to do upload for host %s, uri %s, %v\n", host, info.StoreUri, err)
			} else {
				success = append(success, info.StoreUri)
				break
			}
		}
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SpaceName:   params.SpaceName,
		SkipMeta:    params.SkipMeta,
		SessionKey:  applyResp.SessionKey,
		SuccessOids: success,
		Functions:   functions,
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

// 获取临时上传凭证
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

// 获取上传临时密钥
func (c *ImageXClient) GetUploadAuth(serviceIds []string) (*base.SecurityToken2, error) {
	return c.GetUploadAuthWithExpire(serviceIds, time.Hour)
}

func (c *ImageXClient) GetUploadAuthWithExpire(serviceIds []string, expire time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{
		"ImageX:ApplyImageUpload",
		"ImageX:CommitImageUpload",
	}

	resources := make([]string, 0)
	if len(serviceIds) == 0 {
		resources = append(resources, fmt.Sprintf(ResourceServiceIdTRN, "*"))
	} else {
		for _, sid := range serviceIds {
			resources = append(resources, fmt.Sprintf(ResourceServiceIdTRN, sid))
		}
	}

	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	return c.SignSts2(inlinePolicy, expire)
}

func (c *ImageXClient) updateImageUrls(serviceId string, req *UpdateImageUrlPayload) ([]string, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal request, %v", err)
	}

	respBody, _, err := c.Json("UpdateImageUploadFiles", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("fail to request api UpdateImageUploadFiles, %v", err)
	}

	result := new(UpdateImageUrlPayload)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result.ImageUrls, nil
}

func (c *ImageXClient) RefreshImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionRefresh,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}

func (c *ImageXClient) EnableImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionEnable,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}

func (c *ImageXClient) DisableImageUrls(serviceId string, urls []string) ([]string, error) {
	req := &UpdateImageUrlPayload{
		Action:    ActionDisable,
		ImageUrls: urls,
	}
	return c.updateImageUrls(serviceId, req)
}
