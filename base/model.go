package base

import (
	"net/http"
	"net/url"
	"time"
)

const (
	RegionCnNorth1    = "cn-north-1"
	RegionUsEast1     = "us-east-1"
	RegionApSingapore = "ap-singapore-1"

	timeFormatV4 = "20060102T150405Z"
)

type ServiceInfo struct {
	Timeout     time.Duration
	Host        string
	Header      http.Header
	Credentials Credentials
}

type ApiInfo struct {
	Method  string
	Path    string
	Query   url.Values
	Form    url.Values
	Timeout time.Duration
	Header  http.Header
}

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	Service         string
	Region          string
}

type metadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}

// 统一的JSON返回结果
type CommonResponse struct {
	ResponseMetadata ResponseMetadata
	Result           interface{} `json:"Result,omitempty"`
}

type BaseResp struct {
	Status      string
	CreatedTime int64
	UpdatedTime int64
}

type ErrorObj struct {
	Code    string
	Message string
}

type ResponseMetadata struct {
	RequestId string
	Service   string    `json:",omitempty"`
	Region    string    `json:",omitempty"`
	Action    string    `json:",omitempty"`
	Version   string    `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}
