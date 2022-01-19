package live

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/pkg/errors"
)

func (l *Live) GetDesensitizedAllAppInfos() (*GetDesensitizedAllAppInfosResp, error) {
	respBody, status, err := l.Query("GetDesensitizedAllAppInfos", nil)
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

func (l *Live) MGetStreamsInfo(request *MGetStreamsInfoRequest) (*MGetStreamsInfoResp, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := l.Json("MGetStreamsInfo", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &MGetStreamsInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] resp unmarshal failed")
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

	resp := &MGetStreamsPushInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] resp unmarshal failed")
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}

func (l *Live) MGetStreamsPlayInfo(request *MGetStreamsPlayInfoRequest) (response *MGetStreamsPlayInfoResp, err error) {
	defer func() {
		if err != nil && !request.IsCustomizedStream {
			fallbackResp, fallbackErr := l.mMGetStreamsFallbackPlayInfo(request)
			if fallbackErr != nil {
				_, _ = fmt.Fprintf(os.Stdout, "[vcloud-live] mget stream fall back play info failed, err=%v", fallbackErr.Error())
				return
			}

			fallbackResp.ResponseMetadata = &base.ResponseMetadata{
				Service: "live",
				Region:  l.ServiceInfo.Credentials.Region,
				Action:  "MGetStreamsPlayInfo",
				Version: "2019-10-01",
			}
			response = fallbackResp
			err = nil
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

	resp := &MGetStreamsPlayInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "[vcloud-live] resp unmarshal failed")
	} else {
		resp.ResponseMetadata.Service = "live"
		if resp.ResponseMetadata.Error != nil {
			return resp, fmt.Errorf("[vcloud-live] MGetStreamsPlayInfo failed, errCodeNum:%d, errCode:%s, errMessage:%s",
				resp.ResponseMetadata.Error.CodeN, resp.ResponseMetadata.Error.Code, resp.ResponseMetadata.Error.Message)
		}
		return resp, nil
	}
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

func (l *Live) CreateVOD(request *CreateVODRequest) (*CreateVODResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := l.Json("CreateVOD", nil, string(bts))
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &CreateVODResponse{}
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

	respBody, status, err := l.Query("GetOnlineUserNum", query)
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

func (l *Live) ForbidStream(request *ForbidStreamRequest) (*ForbidStreamResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := l.Json("ForbidStream", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &ForbidStreamResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "live"
		return resp, nil
	}
}
