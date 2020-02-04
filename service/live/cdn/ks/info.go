package ks

import (
	"fmt"

	"code.byted.org/videoarch/common/cdn/base"
)

type CdnHandler struct {
	base.CdnBaseInfo
}

func (c *CdnHandler) GenPullFlvUrl(domain string, appName string, stream string, suffix string) (url string) {
	if domain == "" || appName == "" || stream == "" {
		return
	}

	streamEncode := "/" + appName + "/" + stream + suffix
	url = fmt.Sprintf("http://%s%s.flv", domain, streamEncode)

	return
}

func (c *CdnHandler) GenPullHlsUrl(domain string, appName string, stream string, suffix string) (url string) {
	if domain == "" || appName == "" || stream == "" {
		return
	}

	streamEncode := "/" + appName + "/" + stream + suffix
	url = fmt.Sprintf("http://%s%s/index.m3u8", domain, streamEncode)

	return
}
func (c *CdnHandler) GenPullRtmpUrl(domain string, appName string, stream string, suffix string) (url string) {
	if domain == "" || appName == "" || stream == "" {
		return
	}

	streamEncode := "/" + appName + "/" + stream + suffix
	url = fmt.Sprintf("rtmp://%s%s", domain, streamEncode)

	return
}

func (c *CdnHandler) GenPullCmafUrl(domain string, appName string, stream string, suffix string) (url string) {
	return ""
}

func (c *CdnHandler) GenPullDashUrl(domain string, appName string, stream string, suffix string) (url string) {
	return ""
}
