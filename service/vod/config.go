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
		"RedirectPlay": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RedirectPlay"},
				"Version": []string{"2019-03-15"},
			},
		},
		"GetOriginVideoPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetOriginVideoPlayInfo"},
				"Version": []string{"2019-03-15"},
			},
		},
		"StartTranscode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartTranscode"},
				"Version": []string{"2018-01-01"},
			},
		},
		"UploadMediaByUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
				"Version": []string{"2018-01-01"},
			},
		},
		"ApplyUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"CommitUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SetVideoPublishStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetVideoPublishStatus"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetCdnDomainWeights": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCdnDomainWeights"},
				"Version": []string{"2019-07-01"},
			},
		},
	}
)
