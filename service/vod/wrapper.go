package vod

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//GetPlayInfo 获取播放信息
func (p *Vod) GetPlayInfo(query url.Values) (*GetPlayInfoResp, int, error) {
	respBody, status, err := p.Query("GetPlayInfo", query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetPlayInfoResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "vod"
		return output, status, nil
	}
}

//GetOriginVideoPlayInfo 获取原片播放信息
func (p *Vod) GetOriginVideoPlayInfo(query url.Values) (*GetOriginVideoPlayInfoResp, int, error) {
	respBody, status, err := p.Query("GetOriginVideoPlayInfo", query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetOriginVideoPlayInfoResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "vod"
		return output, status, nil
	}
}

func (p *Vod) StartTranscode(req *StartTranscodeRequest) (*StartTranscodeResp, error) {
	query := url.Values{
		"TemplateId": []string{req.TemplateId},
	}

	reqBody := struct {
		Vid      string
		Input    map[string]interface{}
		Priority int
	}{
		Vid:      req.Vid,
		Input:    req.Input,
		Priority: req.Priority,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal body failed")
	}

	respBody, status, err := p.Json("StartTranscode", query, string(body))
	if err != nil || status != http.StatusOK {
		return nil, errors.Wrap(err, "query error")
	}

	resp := new(StartTranscodeResp)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal body failed")
	}

	return resp, nil
}

func (p *Vod) UploadVideoByUrl(params UploadVideoByUrlParams) (*UploadVideoByUrlResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)
	query.Add("Format", string(params.Format))
	query.Add("SourceUrls", strings.Join(params.SourceUrls, ","))
	query.Add("CallbackArgs", params.CallbackArgs)
	respBody, status, err := p.Query("UploadMediaByUrl", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &UploadVideoByUrlResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) ApplyUpload(params ApplyUploadParam) (*ApplyUploadResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)
	query.Add("SessionKey", params.SessionKey)
	query.Add("FileType", string(params.FileType))
	if params.FileSize != 0 {
		query.Add("FileSize", string(params.FileSize))
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", string(params.UploadNum))
	}

	respBody, status, err := p.Query("ApplyUpload", query)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &ApplyUploadResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) CommitUpload(params CommitUploadParam) (*CommitUploadResp, error) {
	query := url.Values{}
	query.Add("SpaceName", params.SpaceName)

	bts, err := json.Marshal(params.Body)
	if err != nil {
		return nil, err
	}

	respBody, status, err := p.Json("CommitUpload", query, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &CommitUploadResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) Upload(fileBytes []byte, spaceName string, fileType FileType, funcs ...Function) (*CommitUploadResp, error) {
	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("file size is zero")
	}

	params := ApplyUploadParam{
		SpaceName: spaceName,
		FileType:  fileType,
	}

	resp, err := p.ApplyUpload(params)
	if err != nil {
		return nil, err
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return nil, fmt.Errorf("%+v", resp.ResponseMetadata.Error)
	}

	if len(resp.Result.UploadAddress.UploadHosts) == 0 {
		return nil, fmt.Errorf("no tos host found")
	}
	if len(resp.Result.UploadAddress.StoreInfos) == 0 {
		return nil, fmt.Errorf("no store infos found")
	}

	// upload file
	checkSum := fmt.Sprintf("%x", crc32.ChecksumIEEE(fileBytes))
	tosHost := resp.Result.UploadAddress.UploadHosts[0]
	oid := resp.Result.UploadAddress.StoreInfos[0].StoreUri
	auth := resp.Result.UploadAddress.StoreInfos[0].Auth
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("http status=%v, body=%s, remote_addr=%v", rsp.StatusCode, string(b), req.Host)
	}
	defer rsp.Body.Close()

	var tosResp struct {
		Success int         `json:"success"`
		Payload interface{} `json:"payload"`
	}
	err = json.NewDecoder(rsp.Body).Decode(&tosResp)
	if err != nil {
		return nil, err
	}

	if tosResp.Success != 0 {
		return nil, fmt.Errorf("tos err:%+v", tosResp)
	}

	param := CommitUploadParam{
		SpaceName: spaceName,
		Body: CommitUploadBody{
			CallbackArgs: "",
			SessionKey:   resp.Result.UploadAddress.SessionKey,
			Functions:    funcs,
		},
	}
	commitResp, err := p.CommitUpload(param)
	if err != nil {
		return nil, err
	}
	return commitResp, nil
}

// SetVideoPublishStatus 媒资相关
func (p *Vod) SetVideoPublishStatus(SpaceName, Vid, Status string) (*SetVideoPublishStatusResp, int, error) {
	jsonTemp := `{
		"SpaceName" : "%v",
		"Vid" : "%v",
		"Status" : "%v"
	}`

	body := fmt.Sprintf(jsonTemp, SpaceName, Vid, Status)
	respBody, status, err := p.Json("SetVideoPublishStatus", nil, body)
	if err != nil {
		return nil, status, err
	}

	output := new(SetVideoPublishStatusResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	}
	output.ResponseMetadata.Service = "vod"
	return output, status, nil
}

func (p *Vod) GetPlayAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}
	if getPlayInfoToken, err := p.GetSignUrl("GetPlayInfo", query); err == nil {
		ret["GetPlayInfoToken"] = getPlayInfoToken
	} else {
		return "", err
	}

	b, _ := json.Marshal(ret)
	return base64.StdEncoding.EncodeToString(b), nil
}

//GetRedirectPlayUrl get redirected playback addres
func (p *Vod) GetRedirectPlayUrl(params RedirectPlayParam) (string, error) {
	query := url.Values{}
	query.Add("Vid", params.Vid)

	if params.Definition == "" {
		return "", errors.New("Defintion not set")
	}
	query.Add("Definition", string(params.Definition))

	query.Add("Watermark", params.Watermark)

	token, err := p.GetSignUrl("RedirectPlay", query)
	if err != nil {
		return "", err
	}

	apiInfo := p.ApiInfoList["RedirectPlay"]
	url := fmt.Sprintf("http://%s%s?%s", p.ServiceInfo.Host, apiInfo.Path, token)
	return url, nil
}

func (p *Vod) GetUploadAuthToken(space string) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}
	query := url.Values{}
	query.Set("SpaceName", space)

	if applyUploadToken, err := p.GetSignUrl("ApplyUpload", query); err == nil {
		ret["ApplyUploadToken"] = applyUploadToken
	} else {
		return "", err
	}

	if commitUploadToken, err := p.GetSignUrl("CommitUpload", query); err == nil {
		ret["CommitUploadToken"] = commitUploadToken
	} else {
		return "", err
	}

	b, _ := json.Marshal(ret)
	return base64.StdEncoding.EncodeToString(b), nil
}

func (p *Vod) GetCdnDomainWeights(spaceName string) (*GetWeightsResp, error) {
	query := url.Values{}
	query.Set("SpaceName", spaceName)
	respBody, _, err := p.Query("GetCdnDomainWeights", query)
	if err != nil {
		return nil, err
	}

	output := new(GetWeightsResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, err
	}
	return output, nil
}

// poster image related
func (p *Vod) GetDomainInfo(spaceName string, fallbackWeights map[string]int) (*DomainInfo, error) {

	now := time.Now().Unix()
	if p.LastDomainUpdateTime == -1 || now-p.LastDomainUpdateTime > UPDATE_INTERVAL {
		resp, err := p.GetCdnDomainWeights(spaceName)
		if err != nil {
			p.DomainCache[spaceName] = fallbackWeights
		}
		if err := resp.ResponseMetadata.Error; err != nil {
			p.DomainCache[spaceName] = fallbackWeights
		}
		weightsMap, exist := resp.Result[spaceName]
		if !exist || len(weightsMap) == 0 {
			p.DomainCache[spaceName] = fallbackWeights
		}

		p.DomainCache[spaceName] = weightsMap
		p.LastDomainUpdateTime = now
	}

	var (
		mainDomain   string
		backupDomain string
	)
	mainDomain = randWeights(p.DomainCache[spaceName], "")
	if mainDomain == "" {
		return nil, errors.New("rand domain failed")
	}

	backupDomain = randWeights(p.DomainCache[spaceName], mainDomain)
	if backupDomain == "" {
		backupDomain = mainDomain
	}
	return &DomainInfo{MainDomain: mainDomain, BackupDomain: backupDomain}, nil
}

func randWeights(weightsMap map[string]int, excludeDomain string) string {
	var weightSum int
	for domain, weight := range weightsMap {
		if domain == excludeDomain {
			continue
		}
		weightSum += weight
	}
	if weightSum <= 0 {
		return ""
	}
	r := rand.Intn(weightSum) + 1
	for domains, weight := range weightsMap {
		if domains == excludeDomain {
			continue
		}
		r -= weight
		if r <= 0 {
			return domains
		}
	}
	return ""
}

func (p *Vod) GetPosterUrl(spaceName string, uri string, fallbackWeights map[string]int, opts ...OptionFun) (*ImgUrl, error) {
	domainInfos, err := p.GetDomainInfo(spaceName, fallbackWeights)
	if err != nil {
		return nil, err
	}
	opt := &option{
		isHttps: false,
		format:  FORMAT_ORIGINAL,
		tpl:     VOD_TPL_NOOP,
	}
	for _, op := range opts {
		op(opt)
	}
	proto := HTTP
	if opt.isHttps {
		proto = HTTPS
	}
	var tpl string

	if opt.tpl == VOD_TPL_OBJ || opt.tpl == VOD_TPL_NOOP {
		tpl = opt.tpl
	} else {
		tpl = fmt.Sprintf("%s:%d:%d", opt.tpl, opt.w, opt.h)
	}

	return &ImgUrl{
		MainUrl:   fmt.Sprintf("%s://%s/%s~%s.%s", proto, domainInfos.MainDomain, uri, tpl, opt.format),
		BackupUrl: fmt.Sprintf("%s://%s/%s~%s.%s", proto, domainInfos.BackupDomain, uri, tpl, opt.format),
	}, nil
}

func (p *Vod) GetImageUrl(spaceName string, uri string, fallbackWeights map[string]int, opts ...OptionFun) (*ImgUrl, error) {
	domainInfos, err := p.GetDomainInfo(spaceName, fallbackWeights)
	if err != nil {
		return nil, err
	}
	opt := &option{
		isHttps: false,
		format:  FORMAT_ORIGINAL,
		sigKey:  "",
		kv:      nil,
		tpl:     VOD_TPL_SIG,
	}
	for _, op := range opts {
		op(opt)
	}
	path := fmt.Sprintf("/%s~%s.%s", uri, opt.tpl, opt.format)
	sigTxt := path
	if opt.kv != nil {
		if opt.sigKey != "" && opt.kv.Get(KEY_SIG) != "" {
			return nil, ErrKvSig
		}
		sigTxt = fmt.Sprintf("%s?%s", path, opt.kv.Encode())
	}

	if opt.sigKey != "" {
		h := hmac.New(sha1.New, []byte(opt.sigKey))
		h.Write([]byte(sigTxt))
		sig := base64.URLEncoding.EncodeToString(h.Sum(nil))
		if opt.kv == nil {
			opt.kv = url.Values{}
		}
		opt.kv.Add(KEY_SIG, sig)
		path = fmt.Sprintf("%s?%s", path, opt.kv.Encode())
	} else {
		path = sigTxt
	}

	return domainInfos.makeImageUrl(opt.isHttps, path), nil
}

func (p *DomainInfo) makeImageUrl(isHttps bool, path string) *ImgUrl {
	proto := HTTP
	if isHttps {
		proto = HTTPS
	}
	return &ImgUrl{
		MainUrl:   fmt.Sprintf("%s://%s%s", proto, p.MainDomain, path),
		BackupUrl: fmt.Sprintf("%s://%s%s", proto, p.BackupDomain, path),
	}
}
