package live

import (
	"fmt"
	"net/url"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/service/live/util"
	"github.com/pkg/errors"
)

func (l *Live) mMGetStreamsFallbackPlayInfo(request *MGetStreamsPlayInfoRequest) (resp *MGetStreamsPlayInfoResp, err error) {
	if len(request.Streams) == 0 {
		return nil, errors.Errorf("invalid argument: streams=%v, IsCustomizedStream=%v", request.Streams, request.IsCustomizedStream)
	}

	getStreamsResult, err := l.mGetStreamsFallbackInfo(request.Streams)
	if err != nil {
		return nil, err
	}

	playContext := &mGetStreamsPlayContext{
		streams:     request.Streams,
		enableSSL:   request.EnableSSL,
		streamInfos: filterPlayStreamInfos(getStreamsResult),
	}

	scheduleResult, err := l.scheduleByWeight(playContext.streamInfos)
	if err != nil {
		return nil, err
	}

	playContext.scheduleResult = scheduleResult

	playInfos, err := l.deserializePlayInfos(playContext)
	if err != nil {
		return nil, err
	}
	if request.EnableStreamData {
		l.fillPullData(playInfos)
	}
	return &MGetStreamsPlayInfoResp{Result: &MGetStreamsPlayInfoResult{
		PlayInfos: playInfos,
	}}, nil
}

func filterPlayStreamInfos(streamInfos map[string]*StreamInfo) map[string]*StreamInfo {
	filteredMap := make(map[string]*StreamInfo)

	for stream := range streamInfos {
		streamInfo := streamInfos[stream]

		if streamInfo.PlayTypes == "" {
			_, _ = fmt.Fprintf(os.Stdout, "empty playTypes stream: %v", stream)
			continue
		}

		if streamInfo.Resolutions == "" {
			_, _ = fmt.Fprintf(os.Stdout, "empty resolutions stream: %v", stream)
			continue
		}

		if streamInfo.PushMainCdnappId == 0 {
			_, _ = fmt.Fprintf(os.Stdout, "push main appID empty, stream: %v", stream)
			continue
		}

		filteredMap[stream] = streamInfo
	}

	return filteredMap
}

func (l *Live) deserializePlayInfos(playContext *mGetStreamsPlayContext) (
	map[string]*PlayInfo, error) {
	playInfosResult := make(map[string]*PlayInfo)

	for _, stream := range playContext.streams {
		streamInfo, ok := playContext.streamInfos[stream]
		if !ok {
			_, _ = fmt.Fprintf(os.Stdout, "stream :%v not found", stream)
			continue
		}

		scheduleResult, ok := playContext.scheduleResult[stream]
		if !ok {
			_, _ = fmt.Fprintf(os.Stdout, "stream :%v schedules failed", stream)
			continue
		}

		main := l.deserializeElePlayInfo(&deserializePlayParams{
			enableSSL:      playContext.enableSSL,
			stream:         stream,
			streamInfo:     streamInfo,
			scheduleResult: scheduleResult.mainScheduleResult,
		})

		if len(main) == 0 {
			_, _ = fmt.Fprintf(os.Stdout, "deserialize main play info empty, stream: %v", stream)
			continue
		}

		playInfosResult[stream] = &PlayInfo{
			StreamBase: streamInfo.deserializeStreamInfo(),
			Main:       main,
		}
	}

	return playInfosResult, nil
}

func (s *StreamInfo) deserializeStreamInfo() *StreamBase {
	return &StreamBase{
		AppId:      s.Appid,
		Stream:     s.LiveId,
		Extra:      s.Description,
		CreateTime: s.CreateTime,
		Status:     s.Status,
	}
}

func (l *Live) deserializeElePlayInfo(params *deserializePlayParams) []*ElePlayInfo {
	if params.scheduleResult == nil {
		return nil
	}

	streamInfo, scheduleResult := params.streamInfo, params.scheduleResult
	playTypes, resolutions := util.Split(streamInfo.PlayTypes, ","), util.SplitToMap(streamInfo.Resolutions, ",")
	playInfoMap := map[string]*ElePlayInfo{}

	for i := range scheduleResult.templates {
		template := scheduleResult.templates[i]
		templateName := template.Name

		if _, ok := resolutions[templateName]; !ok {
			continue
		}

		genParams := &genElePlayParams{
			enableSSL: params.enableSSL,

			size:      templateName,
			playTypes: playTypes,

			streamInfo:   streamInfo,
			PlayCdnApp:   scheduleResult.playCdnApp,
			Cdn:          scheduleResult.cdn,
			templateInfo: template,
		}

		playURL, err := l.genElePlayURL(genParams)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "gen ele play url fro template: %v, error: %v", templateName, err.Error())
			continue
		}

		playInfoMap[templateName] = playURL
	}

	if len(playInfoMap) == 0 {
		_, _ = fmt.Fprintf(os.Stdout, "deserialize stream: %v play info empty", streamInfo.LiveId)
		return nil
	}

	playInfo := []*ElePlayInfo{}
	for templateName := range playInfoMap {
		playInfo = append(playInfo, playInfoMap[templateName])
	}

	return playInfo
}

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
			playUrl.RtmpUrl = formatSchema(cdnInstance.GenPullRtmpUrl(domain, app, stream, suffix), enableSSL)

		case "hls":
			domain = params.Cdn.PlayHlsDomain
			playUrl.HlsUrl = formatSchema(cdnInstance.GenPullHlsUrl(domain, app, stream, suffix), enableSSL)

		case "flv":
			domain = params.Cdn.PlayFlvDomain
			if params.templateInfo.Name == "md" && params.Cdn.AdminFlvDomain != "" {
				domain = params.Cdn.AdminFlvDomain
				enableSSL = true
			}
			playUrl.FlvUrl = formatSchema(cdnInstance.GenPullFlvUrl(domain, app, stream, suffix), enableSSL)

		case "cmaf":
			domain = params.Cdn.PlayCmafDomain
			playUrl.CmafUrl = formatSchema(cdnInstance.GenPullCmafUrl(domain, app, stream, suffix), enableSSL)

		case "dash":
			domain = params.Cdn.PlayDashDomain
			playUrl.DashUrl = formatSchema(cdnInstance.GenPullDashUrl(domain, app, stream, suffix), enableSSL)

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

func formatSchema(fullUrl string, enableSSL bool) string {
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
