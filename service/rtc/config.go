package rtc

import (
	"net/http"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

const (
	ServiceVersion20180101 = "2018-01-01"
)

type Rtc struct {
	*base.Client
}

// static function
func NewInstance() *Rtc {
	instance := &Rtc{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	return instance
}

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "rtc.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "rtc"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"ByteStartTranscode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ByteStartTranscode"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ByteStopTranscode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ByteStopTranscode"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ByteTranscodeChangeLayout": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ByteTranscodeChangeLayout"},
				"Version": []string{ServiceVersion20180101},
			},
		},
	}
)
