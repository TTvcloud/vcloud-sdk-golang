## 使用方式
#### composer引用
```shell
go get github.com/TTvcloud/vcloud-sdk-golang
```
#### aksk配置

1. 配置在业务代码中，直接使用

2. 配置相关的环境变量`VCLOUD_ACCESSKEY`,`VCLOUD_SECRETKEY`

3. 配置在默认的系统文件中`~/.vcloud/config`

   config文件结构

   ```json
   {
       "ak":"your ak",
       "sk":"your sk"
   }
   ```

## 功能列表

>敬请期待

## Demo

1. 直接调用，会去获取`~/.vcloud/config`下的aksk信息，并且使用服务默认的region信息(这里使用cn-north-1)。

```java
package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/service/iam"
)

func main() {
	resp, code, _ := iam.DefaultInstance.ListAccessKeys(nil)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
```

2. 显示的设置aksk的模式

```java
package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	query := url.Values{}
	query.Set("video_id", "your vid")

	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk"})

	resp, code, _ := vod.DefaultInstance.GetPlayInfo(query)
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
```

