package live

import (
	"fmt"
	"os"

	"github.com/TTvcloud/vcloud-sdk-golang/service/live/util"
	"github.com/pkg/errors"
)

func (l *Live) MGetStreamsFallbackPlayInfo(request *MGetStreamsPlayInfoRequest) (resp *MGetStreamsPlayInfoResp, err error) {
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

	return l.deserializePlayInfos(playContext)
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
	*MGetStreamsPlayInfoResp, error) {
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

	return &MGetStreamsPlayInfoResp{
		Result: playInfosResult,
	}, nil
}

func (s *StreamInfo) deserializeStreamInfo() *StreamBase {
	return &StreamBase{
		AppId:      s.Appid,
		Stream:     s.LiveId,
		Extra:      s.Description,
		CreateTime: s.CreateTime,
	}
}

func (l *Live) deserializeElePlayInfo(params *deserializePlayParams) []*ElePlayInfo {

	if params.scheduleResult == nil {
		return nil
	}

	streamInfo, scheduleResult := params.streamInfo, params.scheduleResult
	playTypes, resolutions := util.Split(streamInfo.PlayTypes, ","), util.SplitToMap(streamInfo.Resolutions, ",")
	playInfo := []*ElePlayInfo{}

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

		playInfo = append(playInfo, playURL)
	}

	if len(playInfo) == 0 {
		_, _ = fmt.Fprintf(os.Stdout, "deserialize stream: %v play info empty", streamInfo.LiveId)
		return nil
	}

	return playInfo
}
