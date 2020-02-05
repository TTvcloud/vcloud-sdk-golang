package live

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func (l *Live) mGetStreamsFallbackInfo(streams []string) (map[string]*StreamInfo, error) {
	result := map[string]*StreamInfo{}
	for _, stream := range streams {
		streamInfo, err := l.getStreamFallbackInfo(stream)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, err.Error())
			continue
		}

		result[stream] = streamInfo
	}

	if len(result) == 0 {
		return nil, errors.New("get streams fall back info empty")
	}
	return result, nil
}

func (l *Live) getStreamFallbackInfo(stream string) (*StreamInfo, error) {
	createTime, pushID, err := parseStreamID(stream)
	if err != nil {
		return nil, fmt.Errorf("parse steam id failed, err=%v", err.Error())
	}

	appInfo, ok := l.getAppInfo(pushID)
	if !ok {
		return nil, fmt.Errorf("not found app info")
	}

	return &StreamInfo{
		Status:           EStreamStatus_Unknown,
		Resolutions:      appInfo.Resolutions,
		PlayTypes:        concatPlayTypes(appInfo),
		PushMainCdnappId: pushID,
		LiveId:           stream,
		Appid:            appInfo.Id,
		CreateTime:       createTime,
		Description:      "{}",
	}, nil
}

func parseStreamID(stream string) (pushID int64, createTime int64, err error) {
	fields := strings.Split(stream, "-")
	if len(fields) != 2 {
		return 0, 0, fmt.Errorf("unparsable stream id=%v", stream)
	}

	id, err := strconv.ParseInt(fields[len(fields)-1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("unparsable stream id=%v", stream)
	}

	if id>>62 != 0 {
		return 0, 0, fmt.Errorf("unparsable stream id=%v", stream)
	}

	pushID = id & pushIDMask
	createTime = (id >> (pushIDBits + countBits)) & tsMask
	return createTime, pushID, nil
}
