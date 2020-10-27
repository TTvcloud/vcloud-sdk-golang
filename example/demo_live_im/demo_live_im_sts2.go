package demo_live_im

import (
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"net/http"
	"time"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "live.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "live"},
	}
	ApiInfoList = map[string]*base.ApiInfo{}
)

/**
返回值示例
{
    "data": {
        "AccessKeyId": "AKTPNmY0NzllZDE0GYzNGM2ODgwYmE5Y5MGZlYTcxYzQ",
        "ExpiredTime": 1579013824,
        "SecretAccessKey": "K6SaRqHsmO691Xc8gFll7ZIkiQap//3bHb9LvwMSyvzQ/GXwgulhM7xOLj",
        "SessionToken": "STS2eyJMVEFjY2zc0tleUlkIjoYWRlZmVmIiwiQWNjZXNzSlZ5MkhWNFhPbUdMT20yVlovNUFZU8yeDloZmI3bFlWWFlUeNjcwNTljMjkwNWUifQ=="
    },
    "status_code": 0
}
 */
func GenerateSts() (map[string]interface{}, error) {
	response := make(map[string]interface{})

	//1.新建一个client
	client := base.NewClient(ServiceInfo, ApiInfoList)
	client.SetAccessKey("你的ak")
	client.SetSecretKey("你的sk")

	//2.设置权限 给一个权限是所有action参数可访问的statement
	inlinePolicy := new(base.Policy)
	statement := base.NewAllowStatement([]string{"live:*"}, []string{})
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)

	//3.生成签名，其中第二个参数有效期可以自己设定
	ret, _ := client.SignSts2(inlinePolicy, time.Hour)

	//4.构造返回值,共4个字段
	response["AccessKeyId"] = ret.AccessKeyId
	response["SecretAccessKey"] = ret.SecretAccessKey
	response["SessionToken"] = ret.SessionToken
	expireTime, _ := time.Parse(time.RFC3339, ret.ExpiredTime)
	response["ExpiredTime"] = expireTime.Unix()
	return response, nil
}

func main()  {
	sts, err := GenerateSts()
	if err != nil {
		fmt.Printf("GenerateSts err: %v\n", err)
	}
	fmt.Printf("sts: %v\n", sts)
}