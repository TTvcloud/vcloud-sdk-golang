package live

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

type allAppInfosCache struct {
	lastCredentials base.Credentials
	count           int

	data *sync.Map
}

func (l *Live) autoFlush() {
	for range time.Tick(time.Second) {
		l.count = (l.count + 1) % UPDATE_INTERVAL_SECOND
		switch {
		case l.count == 0:
			if err := l.updateAllAppInfosCache(); err != nil {
				_, _ = fmt.Fprintf(os.Stdout, err.Error())
				l.count--
				continue
			}

		case !l.isLatestCredentials():
			if err := l.updateAllAppInfosCache(); err != nil {
				_, _ = fmt.Fprintf(os.Stdout, err.Error())
				continue
			}
			l.setLatestCredentials()
		}
	}
}

func (l *Live) isLatestCredentials() bool {
	return l.ServiceInfo.Credentials == l.lastCredentials
}

func (l *Live) setLatestCredentials() {
	l.lastCredentials = l.ServiceInfo.Credentials
}

func (l *Live) updateAllAppInfosCache() error {
	credentials := l.ServiceInfo.Credentials
	if credentials.AccessKeyID == "" || credentials.SecretAccessKey == "" {
		return fmt.Errorf("please set credientials' ak sk before use it")
	}

	result, err := l.GetAllAppInfos()
	if err != nil || result.ResponseMetadata.Error != nil {
		return fmt.Errorf("update appinfocache failed, err=%v, resp=%v\n",
			err, result)
	}

	for pushID := range result.Result.Push2AppInfo {
		l.data.Store(l.genPush2AppInfoKey(pushID), result.Result.Push2AppInfo[pushID])
	}

	for pushID := range result.Result.Push2AllPlayInfos {
		l.data.Store(l.genPush2AllPlayInfosKey(pushID), result.Result.Push2AllPlayInfos[pushID])
	}
	return nil
}

func (l *Live) mustUpdateAllAppInfosCache() {
	if err := l.updateAllAppInfosCache(); err != nil {
		panic(err)
	}
}

func (l *Live) getAppInfo(pushID int64) (*DesensitizedAppInfo, bool) {
	app, ok := l.data.Load(l.genPush2AppInfoKey(pushID))
	if !ok {
		return nil, false
	}

	return app.(*DesensitizedAppInfo), true
}

func (l *Live) getAllPlayInfos(pushID int64) (map[int64]*DesensitizedAllPlayCdnAppInfo, bool) {
	playInfos, ok := l.data.Load(l.genPush2AllPlayInfosKey(pushID))
	if !ok {
		return nil, false
	}

	return playInfos.(map[int64]*DesensitizedAllPlayCdnAppInfo), true
}

func (l *Live) genPush2AppInfoKey(pushID int64) string {
	return fmt.Sprintf("push2App-%v", pushID)
}

func (l *Live) genPush2AllPlayInfosKey(pushID int64) string {
	return fmt.Sprintf("push2Play-%v", pushID)
}

func (l *Live) concatPlayTypes(appInfo *DesensitizedAppInfo) string {
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
