package top

type ApplyResponse struct {
	UploadAddress *UploadAddress `json:"UploadAddress"`
	PluginResult  interface{}    `json:"PluginResult,omitempty"`
}

type UploadAddress struct {
	StoreInfos   []StoreInfo       `json:"StoreInfos"`
	UploadHosts  []string          `json:"UploadHosts"`
	UploadHeader map[string]string `json:"UploadHeader"`
	SessionKey   string            `json:"SessionKey"`
}

type StoreInfo struct {
	StoreUri string `json:"StoreUri"`
	Auth     string `json:"Auth"`
}

type ApplyData struct {
	ApplyResponse `json:"Data"`
}

type CommitData struct {
	CommitResponse `json:"Data"`
}

type CommitResponse struct {
	Vid          string      `json:"Vid,omitempty"`
	SourceInfo   *SourceInfo `json:"SourceInfo,omitempty"`
	PosterUri    string      `json:"PosterUri,omitempty"`
	Encryption   *Encryption `json:"Encryption,omitempty"`
	CallbackArgs string      `json:"CallbackArgs"`
}

type Encryption struct {
	Uri       string            `json:"Uri"`
	SecretKey string            `json:"SecretKey"`
	Algorithm string            `json:"Algorithm"`
	Version   string            `json:"Version"`
	SourceMD5 string            `json:"SourceMd5"`
	Extra     map[string]string `json:"Extra"`
}

type SourceInfo struct {
	StoreUri string  `json:"StoreUri"`
	Md5      string  `json:"Md5,omitempty"`
	Width    int64   `json:"Width"`
	Height   int64   `json:"Height"`
	Duration float32 `json:"Duration,omitempty"`
	Bitrate  float32 `json:"Bitrate,omitempty"`
	Format   string  `json:"Format,omitempty"`
	Size     int64   `json:"Size,omitempty"`
	FileType string  `json:"FileType"`
	Extra    *Extra  `json:"Extra,omitempty"`
}

type Extra struct {
	Filed      string  `json:"Filed"`
	Codec      string  `json:"Codec"`
	Definition string  `json:"Definition"`
	Fps        float32 `json:"Fps"`
	CreateTime string  `json:"CreateTime"`
}

type QueryJob struct {
	JobId     string `json:"JobId"`
	SourceUrl string `json:"SourceUrl"`
}

type UrlUploadResponse struct {
	ValuePairs []QueryJob `json:"Data,omitempty"`
}

type UrlQueryData struct {
	QueryDataResponse `json:"Data"`
}

type QueryDataResponse struct {
	VideoInfoList  []URLSet `json:"VideoInfoList"`
	NotExistJobIds []string `json:"NotExistJobIds"`
}
type URLSet struct {
	RequestId  string                 `json:"RequestId"`
	JobId      string                 `json:"JobId"`
	SourceUrl  string                 `json:"SourceUrl"`
	State      string                 `json:"State"`
	Vid        string                 `json:"Vid"`
	SpaceName  string                 `json:"SpaceName"`
	AccountId  string                 `json:"AccountId"`
	SourceInfo SourceInfo             `json:"SourceInfo"`
	Extra      map[string]interface{} `json:"Extra,omitempty"`
}
