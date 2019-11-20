package imagex

import (
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

// GetImageTemplateConf，获取模板分组配置信息
func (c *ImageXClient) GetTemplateConf(param *GetTemplateConfParam) (*GetTemplateConfResult, error) {
	if c.ServiceInfo.Credentials.Region != base.RegionCnNorth1 {
		return nil, fmt.Errorf("Api GetImageTemplateConf not support region %s", c.ServiceInfo.Credentials.Region)
	}
	query := url.Values{}
	query.Add("GroupId", param.GroupId)
	query.Add("GroupName", param.GroupName)
	respBody, _, err := c.Query("GetImageTemplateConf", query)
	if err != nil {
		return nil, fmt.Errorf("fail to request api GetImageTemplateConf, %v", err)
	}
	result := new(GetTemplateConfResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
