package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	AccessKey = "VCLOUD_ACCESSKEY"
	SecretKey = "VCLOUD_SECRETKEY"
)

type Client struct {
	Client      http.Client
	ServiceInfo *ServiceInfo
	ApiInfoList map[string]*ApiInfo
}

func NewClient(info *ServiceInfo, apiInfoList map[string]*ApiInfo) *Client {
	transport := &http.Transport{
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     10 * time.Second,
	}

	c := http.Client{Transport: transport}
	client := &Client{Client: c, ServiceInfo: info, ApiInfoList: apiInfoList}

	if os.Getenv(AccessKey) != "" && os.Getenv(SecretKey) != "" {
		client.ServiceInfo.Credentials.AccessKeyID = AccessKey
		client.ServiceInfo.Credentials.SecretAccessKey = SecretKey
	} else if _, err := os.Stat(os.Getenv("HOME") + "/.vcloud/config"); err == nil {
		if content, err := ioutil.ReadFile(os.Getenv("HOME") + "/.vcloud/config"); err == nil {
			m := make(map[string]string)
			json.Unmarshal(content, &m)
			if accessKey, ok := m["ak"]; ok {
				client.ServiceInfo.Credentials.AccessKeyID = accessKey
			}
			if secretKey, ok := m["sk"]; ok {
				client.ServiceInfo.Credentials.SecretAccessKey = secretKey
			}
		}
	}

	return client
}

func (client *Client) SetCredential(c Credentials) {
	if c.AccessKeyID != "" {
		client.ServiceInfo.Credentials.AccessKeyID = c.AccessKeyID
	}

	if c.SecretAccessKey != "" {
		client.ServiceInfo.Credentials.SecretAccessKey = c.SecretAccessKey
	}

	if c.Region != "" {
		client.ServiceInfo.Credentials.Region = c.Region
	}
}

func (client *Client) GetSignUrl(api string, query url.Values) (string, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return "", errors.New("相关api不存在")
	}

	query = mergeQuery(query, apiInfo.Query)

	url := fmt.Sprintf("http://%s%s?%s", client.ServiceInfo.Host, apiInfo.Path, query.Encode())
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), url, nil)

	if err != nil {
		return "", errors.New("构建request失败")
	}

	return client.ServiceInfo.Credentials.SignUrl(req), nil
}

func (client *Client) SignSts2(inlinePolicy *Policy, expire time.Duration) (*SecurityToken2, error) {
	var err error
	sts := new(SecurityToken2)
	if sts.AccessKeyId, sts.SecretAccessKey, err = CreateTempAKSK(); err != nil {
		return nil, err
	}

	if innerToken, err := CreateInnerToken(client.ServiceInfo.Credentials, sts.SecretAccessKey, inlinePolicy); err != nil {
		return nil, err
	} else {
		if expire < time.Minute {
			expire = time.Minute
		}
		expireTime := time.Now().Add(expire)
		sts.ExpiredTime = expireTime.Format("20060102T150405Z")
		innerToken.ExpiredTime = expireTime.Unix()

		b, _ := json.Marshal(innerToken)
		sts.SessionToken = "STS2" + base64.StdEncoding.EncodeToString(b)
		return sts, nil
	}
}

func (client *Client) Query(api string, query url.Values) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}

	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	url := fmt.Sprintf("http://%s%s?%s", client.ServiceInfo.Host, apiInfo.Path, query.Encode())
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), url, nil)
	req.Header = header

	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}

	return client.makeRequest(api, req, timeout)
}

func (client *Client) Json(api string, query url.Values, body string) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	url := fmt.Sprintf("http://%s%s?%s", client.ServiceInfo.Host, apiInfo.Path, query.Encode())
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), url, strings.NewReader(body))
	req.Header = header

	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}
	req.Header.Set("Content-Type", "application/json")

	return client.makeRequest(api, req, timeout)
}

func (client *Client) Post(api string, query url.Values, form url.Values) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)
	form = mergeQuery(form, apiInfo.Form)

	url := fmt.Sprintf("http://%s%s?%s", client.ServiceInfo.Host, apiInfo.Path, query.Encode())
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), url, strings.NewReader(form.Encode()))
	req.Header = header

	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return client.makeRequest(api, req, timeout)
}

func (client *Client) makeRequest(api string, req *http.Request, timeout time.Duration) ([]byte, int, error) {
	req = client.ServiceInfo.Credentials.Sign(req)

	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	req = req.WithContext(ctx)

	resp, err := client.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else {
		return []byte(""), 500, errors.New("Bad Request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
