package imagex

import (
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
	data, _, err := c.Post("GetImageOCR", u, url.Values{})
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
