package imagex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

const (
	ImageXHostCn = "imagex.bytedanceapi.com"
	ImageXHostVa = "imagex.us-east-1.bytedanceapi.com"
	ImageXHostSg = "imagex.ap-singapore-1.bytedanceapi.com"

	ImageXTimeout     = 5 * time.Second
	ImageXServiceName = "ImageX"
	ImageXApiVersion  = "2018-08-01"
)

type ImageX struct {
	*base.Client
}

func NewInstance() *ImageX {
	instance := &ImageX{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *ImageX {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("can't find region %s, please check it carefully", region))
	}
	instance := &ImageX{
		Client: base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: ImageXTimeout,
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
			Host:    ImageXHostSg,
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
		"ApplyUploadImageFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUploadImageFile"},
				"Version": []string{ImageXApiVersion},
			},
		},
		"CommitUploadImageFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUploadImageFile"},
				"Version": []string{ImageXApiVersion},
			},
		},
	}
)
