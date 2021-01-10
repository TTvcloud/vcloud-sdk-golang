package edit

import (
	"net/http"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

type Edit struct {
	*base.Client
}

func NewInstance() *Edit {
	instance := &Edit{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "open.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "edit"},
		},
	}

	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "edit"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"SubmitDirectEditTaskAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitDirectEditTaskAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetDirectEditResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDirectEditResult"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SubmitDirectEditTaskSync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitDirectEditTaskSync"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SubmitTemplateTaskAsync": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SubmitTemplateTaskAsync"},
				"Version": []string{"2018-01-01"},
			},
		},
	}
)
