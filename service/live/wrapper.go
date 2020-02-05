package live

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

func (l *Live) GetAllAppInfos() (*GetDesensitizedAllAppInfosResp, error) {
	respBody, status, err := l.Query("GetAllAppInfos", nil)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}
	resp := &GetDesensitizedAllAppInfosResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) CreateStream(request *CreateStreamRequest) (*CreateStreamResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := l.Json("CreateStream", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &CreateStreamResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) MGetStreamsPushInfo(request *MGetStreamsPushInfoRequest) (*MGetStreamsPushInfoResp, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := l.Json("MGetStreamsPushInfo", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &mGetStreamsPushInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] resp unmarshal failed")
	} else {
		resp.ResponseMetadata.Service = "live"
	}

	if resp.Result == nil {
		return &MGetStreamsPushInfoResp{ResponseMetadata: resp.ResponseMetadata}, nil
	}
	pushInfos := map[string]*PushInfo{}
	err = json.Unmarshal(resp.Result.PushInfos, &pushInfos)
	if err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] pushinfo unmarshal failed")
	}

	return &MGetStreamsPushInfoResp{Result: pushInfos, ResponseMetadata: resp.ResponseMetadata}, nil
}

func (l *Live) MGetStreamsPlayInfo(request *MGetStreamsPlayInfoRequest) (response *MGetStreamsPlayInfoResp, err error) {
	defer func() {
		if err != nil && !request.IsCustomizedStream {
			fallbackResp, fallbackErr := l.mMGetStreamsFallbackPlayInfo(request)
			if fallbackErr != nil {
				return
			}
			fallbackResp.ResponseMetadata = response.ResponseMetadata
			fallbackResp.ResponseMetadata.Error = nil
			response = fallbackResp
			return
		}
	}()

	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	respBody, status, err := l.Json("MGetStreamsPlayInfo", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &mGetStreamsPlayInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] resp unmarshal failed")
	} else {
		resp.ResponseMetadata.Service = "live"
	}

	if resp.Result == nil {
		return &MGetStreamsPlayInfoResp{ResponseMetadata: resp.ResponseMetadata}, nil
	}
	playInfos := map[string]*PlayInfo{}
	err = json.Unmarshal(resp.Result.PlayInfos, &playInfos)
	if err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] playinfo unmarshal failed")
	}

	return &MGetStreamsPlayInfoResp{Result: playInfos, ResponseMetadata: resp.ResponseMetadata}, nil
}

func (l *Live) GetVODs(request *GetVODsRequest) (*GetVODsResponse, error) {
	query := url.Values{}
	query.Add("Stream", request.Stream)

	respBody, status, err := l.Query("GetVODs", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetVODsResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) GetRecords(request *GetRecordsRequest) (*GetRecordsResponse, error) {
	query := url.Values{}
	query.Add("Stream", request.Stream)

	respBody, status, err := l.Query("GetRecords", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetRecordsResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) GetSnapshots(request *GetSnapshotsRequest) (*GetSnapshotsResponse, error) {
	query := url.Values{}
	query.Add("Stream", request.Stream)

	respBody, status, err := l.Query("GetSnapshots", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetSnapshotsResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) GetOnlineUserNum(request *GetOnlineUserNumRequest) (*GetOnlineUserNumResponse, error) {
	query := url.Values{}
	query.Add("Stream", request.Stream)
	query.Add("StartTime", strconv.FormatInt(request.StartTime, 10))
	query.Add("EndTime", strconv.FormatInt(request.EndTime, 10))

	respBody, status, err := l.Query("GetStreamTimeShiftInfo", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetOnlineUserNumResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) GetStreamTimeShiftInfo(request *GetStreamTimeShiftInfoRequest) (*GetStreamTimeShiftInfoResponse, error) {
	query := url.Values{}
	query.Add("Stream", request.Stream)
	query.Add("StartTime", strconv.FormatInt(request.StartTime, 10))
	query.Add("EndTime", strconv.FormatInt(request.EndTime, 10))

	respBody, status, err := l.Query("GetStreamTimeShiftInfo", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetStreamTimeShiftInfoResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}
