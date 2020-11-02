## Go SDK使用方式
#### 安装
```
go get github.com/TTvcloud/vcloud-sdk-golang
```
### AK/SK设置
- 在代码里显示调用VodService的方法SetAccessKey/SetSecretKey

- 在当前环境变量中分别设置 VCLOUD_ACCESSKEY="your ak"  VCLOUD_SECRETKEY = "your sk"

- json格式放在～/.vcloud/config中，格式为：{"ak":"your ak","sk":"your sk"}

以上优先级依次降低，建议在代码里显示设置，以便问题排查

### 地域Region设置
- 目前已开放三个地域设置，分别为
  ```
  - cn-north-1 (默认)
  - ap-singapore-1
  - us-east-1
  ```
- 默认为cn-north-1（NewInstance初始化即默认为该地域），如果需要调用其它地域服务，请在初始化函数NewInstanceWithRegion中传入指定地域region，例如：
  ```
  ret, err := vod.NewInstanceWithRegion("us-east-1").GetRedirectPlayUrl(params)
  ```
- 注意1：IAM模块目前只开放cn-north-1区域
- 注意2：不要同时调用NewInstanceWithRegion 和 NewInstance，因为初始化为单例模式，会导致第二次调用不生效

### API

#### 上传

- 通过指定url地址上传

[UploadMediaByUrl](https://open.bytedance.com/docs/4/4652/)

- 服务端直接上传


上传视频包括 [ApplyUpload](https://open.bytedance.com/docs/4/2915/) 和 [CommitUpload](https://open.bytedance.com/docs/4/2916/) 两步

上传封面图包括 [ApplyUpload](https://open.bytedance.com/docs/4/2915/) 和 [ModifyVideoInfo](https://open.bytedance.com/docs/4/4367/) 两步


为方便用户使用，封装方法 UploadVideo 和 UploadPoster， 一步上传


- STS2 鉴权

点播提供的 API
  
GetVideoPlayAuth ( vidList,streamTypeList,watermarkList []string)
  
vidList、streamTypeList、watermarkList 为3种资源，分别代表视频vid、stream type和水印三种资源，切片为空是代表允许访问所有资源。
  
默认的 action 为 vod::GetPlayInfo（不需手动设置）
  
默认过期时间为1小时，可以通过如下 API 自定义过期时间
  
GetVideoPlayAuthWithExpiredTime(vidList, streamTypeList, watermarkList []string, expiredTime time.Duration)
  
示例代码：

```
func main() {
       instance := vod.NewInstance()
       vidList, streamTypeList, watermarkList := make([]string, 0), make([]string, 0), make([]string, 0)
       ret, _ := instance.GetVideoPlayAuth(vidList, streamTypeList, watermarkList)
       b, _ := json.Marshal(ret)
       fmt.Println(string(b))
}
```
  
自定义 STS2 授权模式

```
// 第1步 创建 Policy
inlinePolicy := new(base.Policy)

// 第2步 创建  actions 和 resources
actions := []string{"service:Method"} // eg: vod:GetPlayInfo
resources := make([]string, 0)
// 其中每个 resource 格式类似 "trn:vod::*:video_id/%s"，若允许全部则用 * 替代，否则用实际字符串替代，本例可以填写实际的 vid
if len(vidList) == 0 {
       resources = append(resources, fmt.Sprintf(ResourceVideoFormat, "*"))
} else {
       for _, vid := range vidList {
              resources = append(resources, fmt.Sprintf(ResourceVideoFormat, vid))
       }
}

// 第3步 创建 Statement,允许的 NewAllowStatement, 拒绝的 NewDenyStatement，并添加到 Policy 对应的 Statement 切片里
statement := base.NewAllowStatement(actions, resources)
inlinePolicy.Statement = append(inlinePolicy.Statement, statement)

// 第4步 调用 SignSts2 生成签名
return SignSts2(inlinePolicy, expiredTime)
```

#### 转码
[StartTranscode](https://open.bytedance.com/docs/4/1670/)


#### 发布
[SetVideoPublishStatus](https://open.bytedance.com/docs/4/4709/)


#### 播放
[GetPlayInfo](https://open.bytedance.com/docs/4/2918/)

[GetOriginalPlayInfo](https://open.bytedance.com/docs/4/11148/)

[GetRedirectPlay](https://open.bytedance.com/docs/4/9205/)

#### 封面图
[GetPosterUrl](https://open.bytedance.com/docs/4/5335/)

#### token相关
[GetUploadAuthToken](https://open.bytedance.com/docs/4/6275/)

[GetPlayAuthToken](https://open.bytedance.com/docs/4/6275/)

PS: 上述两个接口和 [GetRedirectPlay](https://open.bytedance.com/docs/4/9205/) 接口中均含有 X-Amz-Expires 这个参数

关于这个参数的解释为：设置返回的playAuthToken或uploadToken或follow 302地址的有效期，目前服务端默认该参数为15min（900s），如果用户认为该有效期过长，可以传递该参数来控制过期时间
。

#### 直播相关
[CreateStream](https://vcloud.bytedance.net/docs/3171/151/)

[MGetStreamsPushInfo](https://vcloud.bytedance.net/docs/3171/184/)

[MGetStreamsPlayInfo](https://vcloud.bytedance.net/docs/3171/185/)

[GetVODs](https://vcloud.bytedance.net/docs/3171/27991/)

[GetRecords](https://vcloud.bytedance.net/docs/3171/27990/)

[GetSnapshots](https://vcloud.bytedance.net/docs/3171/27989/)

[GetOnlineUserNum](https://vcloud.bytedance.net/docs/3171/28269/)

[GetStreamTimeShiftInfo](https://vcloud.bytedance.net/docs/3171/27992/)

#### 更多示例参见
example目录
