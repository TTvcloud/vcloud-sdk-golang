package live

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

func (l *Live) genElePlayURL(params *genElePlayParams) (*ElePlayInfo, error) {
	playUrl := &PlayUrlInfo{}

	app := params.PlayCdnApp.PlayApp
	stream := params.streamInfo.LiveId
	suffix := params.templateInfo.Suffix
	enableSSL := params.enableSSL

	cdnInstance, ok := l.cdnMap[params.Cdn.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported cdn: %v", params.Cdn.Name)
	}

	domain := ""
	for i := range params.playTypes {
		playType := params.playTypes[i]

		switch playType {
		case "rtmp":
			domain = params.Cdn.PlayRtmpDomain
			playUrl.RtmpUrl = replaceSchema(cdnInstance.GenPullRtmpUrl(domain, app, stream, suffix), enableSSL)

		case "hls":
			domain = params.Cdn.PlayHlsDomain
			playUrl.HlsUrl = replaceSchema(cdnInstance.GenPullHlsUrl(domain, app, stream, suffix), enableSSL)

		case "flv":
			domain = params.Cdn.PlayFlvDomain
			if params.templateInfo.Name == "md" && params.Cdn.AdminFlvDomain != "" {
				domain = params.Cdn.AdminFlvDomain
				enableSSL = true
			}
			playUrl.FlvUrl = replaceSchema(cdnInstance.GenPullFlvUrl(domain, app, stream, suffix), enableSSL)

		case "cmaf":
			domain = params.Cdn.PlayCmafDomain
			playUrl.CmafUrl = replaceSchema(cdnInstance.GenPullCmafUrl(domain, app, stream, suffix), enableSSL)

		case "dash":
			domain = params.Cdn.PlayDashDomain
			playUrl.DashUrl = replaceSchema(cdnInstance.GenPullDashUrl(domain, app, stream, suffix), enableSSL)

		default:
			_, _ = fmt.Fprintf(os.Stdout, "unsupported play type: %v", playType)
		}
	}
	if isURLsEmpty(playUrl) {
		return nil, errors.New("all urls empty")
	}

	playInfo := &ElePlayInfo{
		Size: &params.size,
		Url:  playUrl,
	}
	return playInfo, nil
}

func isURLsEmpty(urls *PlayUrlInfo) bool {
	if urls.RtmpUrl == "" && urls.FlvUrl == "" && urls.HlsUrl == "" && urls.CmafUrl == "" && urls.DashUrl == "" {
		return true
	}
	return false
}

func replaceSchema(fullUrl string, enableSSL bool) string {
	parsedURL, err := url.Parse(fullUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "parse url failed: err=%v", err.Error())
		return ""
	}

	// https overwrites http
	if enableSSL && parsedURL.Scheme == string(HTTP) {
		parsedURL.Scheme = string(HTTPS)
	}
	return parsedURL.String()
}
