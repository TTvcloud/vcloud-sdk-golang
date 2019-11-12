package base

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	RegionCnNorth1    = "cn-north-1"
	RegionCnNorth2    = "cn-north-2"
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
	CodeN   int
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

type Policy struct {
	Statement []*Statement
}

const (
	StatementEffectAllow = "Allow"
	StatementEffectDeny  = "Deny"
)

type Statement struct {
	Effect    string
	Action    []string
	Resource  []string
	Condition string `json:",omitempty"`
}

type SecurityToken2 struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	ExpiredTime     string
}

type InnerToken struct {
	LTAccessKeyId         string
	AccessKeyId           string
	SignedSecretAccessKey string
	ExpiredTime           int64
	PolicyString          string
	Signature             string

	// Policy                *Policy `json:",omitempty"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}
