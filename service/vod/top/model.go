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
