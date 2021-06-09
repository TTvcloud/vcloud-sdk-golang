package imagex

import (
	"fmt"
	"net/url"
)

func (c *ImageXClient) GetImageOCR(param *GetImageOCRParam) (*GetImageOCRResult, error) {
	u := url.Values{}
	c.ServiceInfo.Header.Add("X-Top-Account-Id", param.AccountId)
	c.ServiceInfo.Header.Add("Content-type", "application/octet-stream")
	u.Set("Scene", param.Scene)
	u.Set("ServiceId", param.ServiceId)
	u.Set("StoreUri", param.StoreUri)
	data, _, err := c.Binary("GetImageOCR", u, string(param.Image))
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
