package live

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type allAppInfosCache struct {
	data *sync.Map
}

func (l *Live) autoFlush() {
	for l.ServiceInfo.Credentials.SecretAccessKey == "" || l.ServiceInfo.Credentials.AccessKeyID == "" {
	}

	for range time.Tick(UPDATE_INTERVAL) {
		if err := l.updateAllAppInfosCache(); err != nil {
			_, _ = fmt.Fprintf(os.Stdout, err.Error())
		}
	}
}

func (l *Live) updateAllAppInfosCache() error {
	result, err := l.GetAllAppInfos()
	if err != nil || result.ResponseMetadata.Error != nil {
		return fmt.Errorf("update appinfocache failed, err=%v, resp=%v\n",
			err, result)
	}

	for pushID := range result.Result.Push2AppInfo {
		l.data.Store(genPush2AppInfoKey(pushID), result.Result.Push2AppInfo[pushID])
	}

	for pushID := range result.Result.Push2AllPlayInfos {
		l.data.Store(genPush2AllPlayInfosKey(pushID), result.Result.Push2AllPlayInfos[pushID])
	}
	return nil
}

func (l *Live) getAppInfo(pushID int64) (*DesensitizedAppInfo, bool) {
	app, ok := l.data.Load(genPush2AppInfoKey(pushID))
	if !ok {
		return nil, false
	}

	return app.(*DesensitizedAppInfo), true
}

func (l *Live) getAllPlayInfos(pushID int64) (map[int64]*DesensitizedAllPlayCdnAppInfo, bool) {
	playInfos, ok := l.data.Load(genPush2AllPlayInfosKey(pushID))
	if !ok {
		return nil, false
	}

	return playInfos.(map[int64]*DesensitizedAllPlayCdnAppInfo), true
}

func genPush2AppInfoKey(pushID int64) string {
	return fmt.Sprintf("push2App-%v", pushID)
}

func genPush2AllPlayInfosKey(pushID int64) string {
	return fmt.Sprintf("push2Play-%v", pushID)
}

func concatPlayTypes(appInfo *DesensitizedAppInfo) string {
	if appInfo == nil {
		return ""
	}

	playTypes := map[string]bool{
		playTypeRtmp: appInfo.IsPlayRtmp,
		playTypeFlv:  appInfo.IsPlayFlv,
		playTypeHls:  appInfo.IsPlayHls,
		playTypeDash: appInfo.IsPlayDash,
		playTypeCmaf: appInfo.IsPlayCmaf,
	}

	result := []string{}
	for playType, ok := range playTypes {
		if ok {
			result = append(result, playType)
		}
	}

	return strings.Join(result, ",")
}
