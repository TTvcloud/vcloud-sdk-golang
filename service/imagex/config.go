package imagex

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

const (
	updateInterval = 10

	ImageXHostCn = "imagex.bytedanceapi.com"
	ImageXHostVa = "imagex.us-east-1.bytedanceapi.com"
	ImageXHostSg = "imagex.ap-singapore-1.bytedanceapi.com"

	ImageXInnerHostCn = "imagex.byted.org"
	ImageXInnerHostVa = "imagex.us-east-1.byted.org"
	ImageXInnerHostSg = "imagex.ap-singapore-1.byted.org"

	ImageXTimeout              = 10 * time.Second
	ImageXServiceName          = "ImageX"
	ImageXApiVersion           = "2018-08-01"
	ImageXDomainWeightsVersion = "2019-07-01"

	ResourceServiceIdTRN = "trn:ImageX:*:*:ServiceId/%s"
)

type ImageXClient struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

func NewInstance() *ImageXClient {
	instance := &ImageXClient{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *ImageXClient {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("can't find region %s, please check it carefully", region))
	}
	instance := &ImageXClient{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: ImageXTimeout,
			Scheme:  "https",
			Host:    ImageXHostCn,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ImageXServiceName,
			},
		},
		base.RegionUsEast1: {
			Timeout: ImageXTimeout,
			Scheme:  "https",
			Host:    ImageXHostVa,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionUsEast1,
				Service: ImageXServiceName,
			},
		},
		base.RegionApSingapore: {
			Timeout: ImageXTimeout,
			Scheme:  "https",
			Host:    ImageXHostSg,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionApSingapore,
				Service: ImageXServiceName,
			},
		},
		base.InnerRegionCnNorth1: {
			Timeout: ImageXTimeout,
			Scheme:  "http",
			Host:    ImageXInnerHostCn,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ImageXServiceName,
			},
		},
		base.InnerRegionUsEast1: {
			Timeout: ImageXTimeout,
			Scheme:  "http",
			Host:    ImageXInnerHostVa,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionUsEast1,
				Service: ImageXServiceName,
			},
		},
		base.InnerRegionApSingapore: {
			Timeout: ImageXTimeout,
			Scheme:  "http",
			Host:    ImageXInnerHostSg,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionApSingapore,
				Service: ImageXServiceName,
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// 资源管理相关
		"DeleteImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageUploadFiles"},
				"Version": []string{ImageXApiVersion},
			},
		},
		"ApplyImageUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyImageUpload"},
				"Version": []string{ImageXApiVersion},
			},
		},
		"CommitImageUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitImageUpload"},
				"Version": []string{ImageXApiVersion},
			},
		},
		"UpdateImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageUploadFiles"},
				"Version": []string{ImageXApiVersion},
			},
		},

		// 模板相关
		"GetImageTemplateConf": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTemplateConf"},
				"Version": []string{ImageXApiVersion},
			},
		},
		//域名调度相关
		"GetCdnDomainWeights": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCdnDomainWeights"},
				"Version": []string{ImageXDomainWeightsVersion},
			},
		},
		"GetImageThemeColor": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageThemeColor"},
				"Version": []string{ImageXApiVersion},
			},
		},
	}
)
