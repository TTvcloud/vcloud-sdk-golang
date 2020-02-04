package cdn

import (
	"fmt"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/ali"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/byte"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/fcdn"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/fws"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/ks"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn/ws"
)

const (
	CDN_ALI  = "ali"
	CDN_WS   = "ws"
	CDN_KS   = "ks"
	CDN_FCDN = "fcdn"
	CDN_BYTE = "byte"
	CDN_FWS  = "fws"
)

type CDNInterface interface {
	//获取拉流地址
	GenPullFlvUrl(domain string, appName string, stream string, suffix string) (url string)
	GenPullHlsUrl(domain string, appName string, stream string, suffix string) (url string)
	GenPullRtmpUrl(domain string, appName string, stream string, suffix string) (url string)
	GenPullCmafUrl(domain string, appName string, stream string, suffix string) (url string)
	GenPullDashUrl(domain string, appName string, stream string, suffix string) (url string)
}

func registerCdnInstance(mapCdn map[string]CDNInterface, cdnName string, CI CDNInterface) {
	if cdnName == "" || CI == nil {
		_, _ = fmt.Fprintf(os.Stdout, "Register key[%s] Fail! Input nil param\n", cdnName)
		return
	}

	if _, ok := mapCdn[cdnName]; ok {
		_, _ = fmt.Fprintf(os.Stdout, "Register key[%s] multi-times\n", cdnName)
		return
	}

	mapCdn[cdnName] = CI
	_, _ = fmt.Fprintf(os.Stdout, "Register Cdn: "+cdnName)
}

func Init() map[string]CDNInterface {
	mapCdn := make(map[string]CDNInterface)
	registerCdnInstance(mapCdn, CDN_ALI, &ali.CdnHandler{})
	registerCdnInstance(mapCdn, CDN_WS, &ws.CdnHandler{})
	registerCdnInstance(mapCdn, CDN_KS, &ks.CdnHandler{})
	registerCdnInstance(mapCdn, CDN_FCDN, &fcdn.CdnHandler{})
	registerCdnInstance(mapCdn, CDN_BYTE, &byte.CdnHandler{})
	registerCdnInstance(mapCdn, CDN_FWS, &fws.CdnHandler{})
	_, _ = fmt.Fprintf(os.Stdout, "init cdn handlers finished")
	return mapCdn
}
