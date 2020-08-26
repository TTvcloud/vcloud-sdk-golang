package live

import (
	"fmt"
	"os"
	"sort"

	json "github.com/json-iterator/go"
)

type PullData struct {
	Data map[string]*PullURLS `json:"data"` // 流信息
}

type PullURLS struct {
	Main   *URLData `json:"main,omitempty"`   // 主流信息
	Backup *URLData `json:"backup,omitempty"` // 备流信息
}

type URLData struct {
	Flv       string  `json:"flv"`                  // flv 拉流地址
	Hls       string  `json:"hls"`                  // hls 拉流地址
	Cmaf      string  `json:"cmaf"`                 // cmaf 拉流地址
	Dash      string  `json:"dash"`                 // dash 拉流地址
	SDKParams *string `json:"sdk_params,omitempty"` // sdk拉流参数 拉流地址
}

func (l *Live) fillPullData(origin map[string]*PlayInfo) {
	allowedSize2Priority := map[string]int{
		"ao":     10,
		"ld":     20,
		"sd":     30,
		"hd":     40,
		"uhd":    50,
		"origin": 60,
	}
	for stream := range origin {
		playInfo := origin[stream]

		pullData := &PullData{
			Data: map[string]*PullURLS{},
		}
		fillData(pullData.Data, true, playInfo.Main, allowedSize2Priority)
		fillData(pullData.Data, false, playInfo.Backup, allowedSize2Priority)
		if len(playInfo.Backup) != 0 {
			filterOutIntersectionData(pullData.Data)
		}

		jsonPullData, err := json.Marshal(pullData)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "pullData json encode failed, error= %v", err)
			return
		}
		res := string(jsonPullData)
		playInfo.StreamData = &res

		playInfo.StreamSizes = fillSizes(pullData.Data, allowedSize2Priority)
	}
}

func fillData(data map[string]*PullURLS, isMain bool, eleInfos []*ElePlayInfo, allowedSize map[string]int) {
	if len(eleInfos) == 0 {
		return
	}

	for i := range eleInfos {
		eleInfo := eleInfos[i]
		if _, ok := allowedSize[eleInfo.GetSize()]; !ok {
			continue
		}

		pullURLS := data[eleInfo.GetSize()]

		if pullURLS == nil {
			pullURLS = &PullURLS{}
			data[eleInfo.GetSize()] = pullURLS
		}

		urlData := deserializeURLData(eleInfo)
		if isMain {
			pullURLS.Main = &urlData
		} else {
			pullURLS.Backup = &urlData
		}
	}
}

func filterOutIntersectionData(data map[string]*PullURLS) {
	for size, pullURLS := range data {
		if pullURLS.Main == nil || pullURLS.Backup == nil {
			delete(data, size)
		}
	}
}

func deserializeURLData(eleInfo *ElePlayInfo) URLData {
	sdkParams := "{}"
	return URLData{
		Flv:       eleInfo.GetUrl().FlvUrl,
		Hls:       eleInfo.GetUrl().HlsUrl,
		Cmaf:      eleInfo.GetUrl().CmafUrl,
		Dash:      eleInfo.GetUrl().DashUrl,
		SDKParams: &sdkParams,
	}
}

func fillSizes(data map[string]*PullURLS, priority map[string]int) []string {
	sizes := []string{}
	for size := range data {
		if priority[size] == 0 {
			continue
		}
		sizes = append(sizes, size)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return priority[sizes[i]] < priority[sizes[j]]
	})
	return sizes
}
