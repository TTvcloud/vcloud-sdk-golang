package vod

import (
	"net/http"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

type Vod struct {
	*base.Client
}

var DefaultInstance = NewInstance()

// static function
func NewInstance() *Vod {
	instance := &Vod{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	return instance
}

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "vod.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
				"Version": []string{"2019-03-15"},
			},
		},
	}
)
