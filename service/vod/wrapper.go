package vod

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

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

	body, err := json.Marshal(req)
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

func (p *Vod) UploadMediaByUrl(params UploadMediaByUrlParams) (*UploadMediaByUrlResp, error) {
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

	resp := &UploadMediaByUrlResp{}
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

func (p *Vod) ModifyVideoInfo(body ModifyVideoInfoBody) (*ModifyVideoInfoResp, error) {
	query := url.Values{}

	bts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	respBody, status, err := p.Json("ModifyVideoInfo", query, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}
	resp := &ModifyVideoInfoResp{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "vod"
		return resp, nil
	}
}

func (p *Vod) Upload(rd io.Reader, size int64, spaceName string, fileType FileType) (string, string, error) {
	if size == 0 {
		return "", "", fmt.Errorf("file size is zero")
	}

	params := ApplyUploadParam{
		SpaceName: spaceName,
		FileType:  fileType,
	}

	resp, err := p.ApplyUpload(params)
	if err != nil {
		return "", "", err
	}

	if resp.ResponseMetadata.Error != nil && resp.ResponseMetadata.Error.Code != "0" {
		return "", "", fmt.Errorf("%+v", resp.ResponseMetadata.Error)
	}

	if len(resp.Result.UploadAddress.UploadHosts) == 0 {
		return "", "", fmt.Errorf("no tos host found")
	}
	if len(resp.Result.UploadAddress.StoreInfos) == 0 {
		return "", "", fmt.Errorf("no store infos found")
	}

	// upload file
	tosHost := resp.Result.UploadAddress.UploadHosts[0]
	oid := resp.Result.UploadAddress.StoreInfos[0].StoreUri
	sessionKey := resp.Result.UploadAddress.SessionKey
	auth := resp.Result.UploadAddress.StoreInfos[0].Auth
	client := &http.Client{}
	if int(size) < MinChunckSize {
		bts, err := ioutil.ReadAll(rd)
		if err != nil {
			return "", "", err
		}
		if err := p.directUpload(tosHost, oid, auth, bts, client); err != nil {
			return "", "", err
		}
	} else {
		uploadPart := UploadPartCommon{
			TosHost: tosHost,
			Oid:     oid,
			Auth:    auth,
		}
		if err := p.chunkUpload(rd, uploadPart, client); err != nil {
			return "", "", err
		}
	}
	return oid, sessionKey, nil
}

func (p *Vod) directUpload(tosHost string, oid string, auth string, fileBytes []byte, client *http.Client) error {
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(fileBytes))
	url := fmt.Sprintf("http://%s/%s", tosHost, oid)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", auth)

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &UploadPartCommonResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) chunkUpload(rd io.Reader, uploadPart UploadPartCommon, client *http.Client) error {
	uploadID, err := p.initUploadPart(uploadPart.TosHost, uploadPart.Oid, uploadPart.Auth, client)
	if err != nil {
		return err
	}
	pre, cur := make([]byte, MinChunckSize), make([]byte, MinChunckSize)
	parts := make([]string, 0)

	n, err := io.ReadFull(rd, pre) // 保留前一个分片，避免最后的分片小于MinChunkSize上传失败
	if err != nil {
		return err
	}
	pre = pre[:n]
	i := 0
	for ; ; i++ {
		n, err = io.ReadFull(rd, cur)
		if err == io.EOF {
			break
		}
		if err == io.ErrUnexpectedEOF {
			//当 io 本身出现问题时，n = 0，n =0 正确情况只有EOF， ErrUnexpectedEOF 为错误，不处理会发生上传不完整情况
			if n == 0 {
				return err
			}
		} else if err != nil {
			return err
		}

		// UploadPart要求分片不能小于MinChunkSize，否则会报错
		// 所以如果当前分片小于MinChunkSize，表示文件已经读到末尾，当前分片与前一片合起来一起上传。
		if int64(n) < MinChunckSize {
			pre = append(pre, cur[:n]...)
			break
		}

		part, err := p.uploadPart(uploadPart, uploadID, i, pre, client)
		if err != nil { // retry part
			part, err = p.uploadPart(uploadPart, uploadID, i, pre, client)
		}
		if err != nil {
			return err
		}
		parts = append(parts, part)
		copy(pre, cur[:n])
		pre = pre[:n]
	}
	// 退出的条件有两个：文件读EOF；读字节数小于MinChunkSize；这两种情况都需要把pre保存下来的分片上传
	part, err := p.uploadPart(uploadPart, uploadID, i, pre, client)
	if err != nil {
		return err
	}
	parts = append(parts, part)
	return p.uploadMergePart(uploadPart, uploadID, parts, client)
}

func (p *Vod) initUploadPart(tosHost string, oid string, auth string, client *http.Client) (string, error) {
	url := fmt.Sprintf("http://%s/%s?uploads", tosHost, oid)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", auth)
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return res.PayLoad.UploadID, nil
}

func (p *Vod) uploadPart(uploadPart UploadPartCommon, uploadID string, partNumber int, data []byte, client *http.Client) (string, error) {
	url := fmt.Sprintf("http://%s/%s?partNumber=%d&uploadID=%s", uploadPart.TosHost, uploadPart.Oid, partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", uploadPart.Auth)

	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}
	res := &UploadPartResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return "", err
	}
	if res.Success != 0 {
		return "", errors.New(res.Error.Message)
	}
	return checkSum, nil
}

func (p *Vod) uploadMergePart(uploadPart UploadPartCommon, uploadID string, checkSum []string, client *http.Client) error {
	url := fmt.Sprintf("http://%s/%s?uploadID=%s", uploadPart.TosHost, uploadPart.Oid, uploadID)
	body, err := p.genMergeBody(checkSum)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", uploadPart.Auth)
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}
	res := &UploadMergeResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}
	if res.Success != 0 {
		return errors.New(res.Error.Message)
	}
	return nil
}

func (p *Vod) genMergeBody(checkSum []string) (string, error) {
	if len(checkSum) == 0 {
		return "", errors.New("body crc32 empty")
	}
	s := make([]string, len(checkSum))
	for partNumber, crc := range checkSum {
		s[partNumber] = fmt.Sprintf("%d:%s", partNumber, crc)
	}
	return strings.Join(s, ","), nil
}

func (p *Vod) UploadPoster(vid string, fileBytes []byte, spaceName string, fileType FileType) (string, error) {
	oid, _, err := p.Upload(bytes.NewReader(fileBytes), int64(len(fileBytes)), spaceName, fileType)
	if err != nil {
		return "", err
	}

	body := ModifyVideoInfoBody{
		SpaceName: spaceName,
		Vid:       vid,
		Info: UserMetaInfo{
			PosterUri: oid,
		},
	}
	_, err = p.ModifyVideoInfo(body)
	if err != nil {
		return "", err
	}
	return oid, nil
}

func (p *Vod) UploadVideoWithCallbackArgs(filePath string, spaceName string, fileType FileType, callbackArgs string, funcs ...Function) (*CommitUploadResp, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return p.UploadVideoInner(file, stat.Size(), spaceName, fileType, callbackArgs, funcs...)
}

func (p *Vod) UploadVideoInner(rd io.Reader, size int64, spaceName string, fileType FileType, callbackArgs string, funcs ...Function) (*CommitUploadResp, error) {
	_, sessionKey, err := p.Upload(rd, size, spaceName, fileType)
	if err != nil {
		return nil, err
	}

	param := CommitUploadParam{
		SpaceName: spaceName,
		Body: CommitUploadBody{
			CallbackArgs: callbackArgs,
			SessionKey:   sessionKey,
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
	if params.Vid == "" {
		return "", errors.New("Vid not set")
	}
	query.Add("Vid", params.Vid)
	if params.Definition != "" {
		query.Add("Definition", string(params.Definition))
	}
	if params.Watermark != "" {
		query.Add("Watermark", params.Watermark)
	}
	if params.Expires != "" {
		query.Add("X-Amz-Expires", params.Expires)
	}

	token, err := p.GetSignUrl("RedirectPlay", query)
	if err != nil {
		return "", err
	}

	apiInfo := p.ApiInfoList["RedirectPlay"]
	url := fmt.Sprintf("http://%s%s?%s", p.ServiceInfo.Host, apiInfo.Path, token)
	return url, nil
}

func (p *Vod) GetUploadAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}

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
	query.Set("ProductLine", "vcloud")
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

	var cache map[string]int
	var ok bool
	p.Lock.RLock()
	if cache, ok = p.DomainCache[spaceName]; !ok {
		p.Lock.RUnlock()

		p.Lock.Lock()
		if cache, ok = p.DomainCache[spaceName]; !ok {
			var weightsMap map[string]int
			var exist bool
			resp, err := p.GetCdnDomainWeights(spaceName)
			if err != nil || resp == nil || resp.ResponseMetadata == nil || resp.ResponseMetadata.Error != nil {
				weightsMap = fallbackWeights
			} else {
				weightsMap, exist = resp.Result[spaceName]
				if !exist || len(weightsMap) == 0 {
					weightsMap = fallbackWeights
				}
			}
			p.DomainCache[spaceName] = weightsMap

			p.Lock.Unlock()
			cache = p.DomainCache[spaceName]

			go func() {
				for range time.Tick(UPDATE_INTERVAL * time.Second) {
					var weightsMap map[string]int
					resp, err := p.GetCdnDomainWeights(spaceName)
					if err != nil || resp == nil || resp.ResponseMetadata == nil || resp.ResponseMetadata.Error != nil {
						weightsMap = fallbackWeights
					} else {
						weightsMap, exist := resp.Result[spaceName]
						if !exist || len(weightsMap) == 0 {
							weightsMap = fallbackWeights
						}
					}
					p.Lock.Lock()
					p.DomainCache[spaceName] = weightsMap
					p.Lock.Unlock()
				}
			}()
		} else {
			p.Lock.Unlock()
		}
	} else {
		p.Lock.RUnlock()
	}

	var (
		mainDomain   string
		backupDomain string
	)
	mainDomain = randWeights(cache, "")
	if mainDomain == "" {
		return nil, errors.New("rand domain failed")
	}

	backupDomain = randWeights(cache, mainDomain)
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

func (p *Vod) GetVideoPlayAuthWithExpiredTime(vidList, streamTypeList, watermarkList []string, expiredTime time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{ActionGetPalyInfo}
	resources := make([]string, 0)

	// 设置vid的resource权限
	resources = p.AddResourcesFormat(vidList, resources, ResourceVideoFormat)

	// 设置streamType的resource权限
	resources = p.AddResourcesFormat(streamTypeList, resources, ResourceStreamTypeFormat)

	// 设置watermark的resource权限
	resources = p.AddResourcesFormat(watermarkList, resources, ResourceWatermarkFormat)

	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)

	return p.SignSts2(inlinePolicy, expiredTime)
}

func (p *Vod) AddResourcesFormat(list []string, resources []string, resourceFormat string) []string {
	if len(list) == 0 {
		resources = append(resources, fmt.Sprintf(resourceFormat, Star))
	} else {
		for _, v := range list {
			resources = append(resources, fmt.Sprintf(resourceFormat, v))
		}
	}
	return resources
}

func (p *Vod) GetVideoPlayAuth(vidList, streamTypeList, watermarkList []string) (*base.SecurityToken2, error) {
	return p.GetVideoPlayAuthWithExpiredTime(vidList, streamTypeList, watermarkList, time.Hour)
}

func (p *Vod) GetUploadAuthWithExpiredTime(expiredTime time.Duration) (*base.SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{"vod:ApplyUpload", "vod:CommitUpload"}
	resources := make([]string, 0)
	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	return p.SignSts2(inlinePolicy, expiredTime)
}

func (p *Vod) GetUploadAuth() (*base.SecurityToken2, error) {
	return p.GetUploadAuthWithExpiredTime(time.Hour)
}
