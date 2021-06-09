package imagex

import (
	"bytes"
	"fmt"
	"net/url"
)

const LicenseScene = "license"

func (c *ImageXClient) GetImageOCR(param *GetImageOCRParam) (*GetImageOCRResult, error) {
	u := url.Values{}
	c.ServiceInfo.Header.Add("X-Top-Account-Id", param.AccountId)
	u.Set("Scene", param.Scene)
	u.Set("ServiceId", param.ServiceId)
	u.Set("StoreUri", param.StoreUri)
	if param.StoreUri == "" {
		c.ServiceInfo.Header.Add("Content-type", "application/octet-stream")
	}
	data, _, err := c.PostWithBody("GetImageOCR", u, bytes.NewReader(param.Image))
	if err != nil {
		return nil, fmt.Errorf("fail to request api GetImageOCR, %v\n", err)
	}
	result := new(GetImageOCRResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
