package ecs

import (
	"net/http"
	"net/url"
	"time"
	"vcloud-sdk-golang/base"
	"vcloud-sdk-golang/bytedance"
)

// ECS methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ECS struct {
	*base.Client
}

func New(session bytedance.Session) *ECS {
	client := &ECS{
		Client: base.NewClient(ServiceInfoMap[session.Config.Region], ApiInfoList),
	}
	return client
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth2: {
			Timeout: 5 * time.Second,
			Host:    "ecs.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth2, Service: "ecs"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetInstances": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInstances"},
				"Version": []string{"2018-11-01"},
			},
		},
		"RunInstances": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RunInstances"},
				"Version": []string{"2018-11-01"},
			},
		},
	}
)
