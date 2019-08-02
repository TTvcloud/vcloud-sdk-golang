## 使用方式
#### go引用
```shell
go get github.com/TTvcloud/vcloud-sdk-golang
```
### AK/SK设置
- 在代码里显示调用VodService的方法SetAccessKey/SetSecretKey

- 在当前环境变量中分别设置 VCLOUD_ACCESSKEY="your ak"  VCLOUD_SECRETKEY = "your sk"

- json格式放在～/.vcloud/config中，格式为：{"ak":"your ak","sk":"your sk"}

以上优先级依次降低，建议在代码里显示设置，以便问题排查

### API

#### 上传

- 通过指定url地址上传

[UploadMediaByUrl](https://open.bytedance.com/docs/4/4652/)

- 服务端直接上传


上传视频包括 [ApplyUpload](https://open.bytedance.com/docs/4/2915/) 和 [CommitUpload](https://open.bytedance.com/docs/4/2916/) 两步

上传封面图包括 [ApplyUpload](https://open.bytedance.com/docs/4/2915/) 和 [ModifyVideoInfo](https://open.bytedance.com/docs/4/4367/) 两步


为方便用户使用，封装方法 UploadVideo 和 UploadPoster， 一步上传


#### 转码
[StartTranscode](https://open.bytedance.com/docs/4/1670/)


#### 发布
[SetVideoPublishStatus](https://open.bytedance.com/docs/4/4709/)


#### 播放
[GetPlayInfo](https://open.bytedance.com/docs/4/2918/)

[GetOriginVideoPlayInfo](https://open.bytedance.com/docs/4/11148/)

[GetRedirectPlay](https://open.bytedance.com/docs/4/9205/)

#### 封面图
[GetPosterUrl]()

#### token相关
[GetUploadAuthToken](https://open.bytedance.com/docs/4/6275/)

[GetPlayAuthToken](https://open.bytedance.com/docs/4/6275/)

PS: 上述两个接口和 [GetRedirectPlay](https://open.bytedance.com/docs/4/9205/) 接口中均含有 X-Amz-Expires 这个参数

关于这个参数的解释为：设置返回的playAuthToken或uploadToken或follow 302地址的有效期，目前服务端默认该参数为15min（900s），如果用户认为该有效期过长，可以传递该参数来控制过期时间
。

#### 更多示例参见
example目录

### Change log

#### 0.0.5
- 去掉image X 相关
- 增加封面图上传接口
- getUploadAuthToken/getPlayAuthToken/RedirectPlay支持X-Amz-Expires参数
- 代码格式优化