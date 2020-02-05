package live

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

func (l *Live) scheduleByWeight(streamInfos map[string]*StreamInfo) (
	map[string]*schedulePlayResult, error) {

	result := make(map[string]*schedulePlayResult)

	for stream := range streamInfos {
		streamInfo := streamInfos[stream]

		scheduleResult := &schedulePlayResult{
			streamInfo: streamInfo,
		}

		playCdnAppInfos, ok := l.getAllPlayInfos(streamInfo.PushMainCdnappId)
		if !ok {
			_, _ = fmt.Fprintf(os.Stdout, "get play cdn app info failed, pushID=%v", streamInfo.PushMainCdnappId)
			continue
		}

		main, err := scheduleStreamByWeight(playCdnAppInfos)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "schedule main for stream: %v err: %v", stream, err.Error())
			continue
		}

		scheduleResult.mainScheduleResult = main
		result[stream] = scheduleResult
	}

	if len(result) == 0 {
		return nil, errors.New("schedule by weight result empty")
	}

	return result, nil
}

func scheduleStreamByWeight(playCdnAppInfos map[int64]*DesensitizedAllPlayCdnAppInfo) (*scheduleElePlayResult, error) {

	if len(playCdnAppInfos) == 0 {
		return nil, errors.New("empty playCdnAppInfos")
	}

	totalWeight := int64(0)

	for _, playCdnAppInfo := range playCdnAppInfos {
		weight := int64(playCdnAppInfo.PlayCdnApp.PlayProportion)
		totalWeight += weight
	}

	randNum := rand.Int63n(totalWeight)
	var tempWeight int64

	for i := range playCdnAppInfos {
		candidate := playCdnAppInfos[i]

		if candidate.PlayCdnApp.PlayProportion == 0 {
			continue
		}

		tempWeight += int64(candidate.PlayCdnApp.PlayProportion)

		if tempWeight > randNum {
			return &scheduleElePlayResult{
				playCdnApp: candidate.PlayCdnApp,
				cdn:        candidate.Cdn,
				templates:  addOriginToTemplates(candidate.Templates),
			}, nil
		}
	}

	return nil, errors.New("schedule play by weight failed")
}

func addOriginToTemplates(origin []*DesensitizedTemplateInfo) []*DesensitizedTemplateInfo {
	return append(origin, &DesensitizedTemplateInfo{
		Name: "origin",
		Size: "origin",
	})
}
