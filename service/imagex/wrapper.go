package imagex

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

const (
	imagexProductLine = "imagex"
)

/*GetImagexURL Get the image url
 * @param serviceId 	your serviceId
 * @param uri 			the image uri that your upload image
 * @param tpl			the template of your config on the imagex
 * @param opts			options, for example WithHttps(), WithFormat()...
 */
func (c *ImageXClient) GetImagexURL(serviceId, uri, tpl string, opts ...OptionFun) (*ImgUrl, error) {

	opt := &option{
		isHttps:         false,
		format:          FORMAT_ORIGINAL,
		sigKey:          "",
		kv:              nil,
		fallbackWeights: nil,
	}
	for _, op := range opts {
		op(opt)
	}

	domainInfos, err := c.GetDomainInfo(serviceId, opt.fallbackWeights)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/%s~%s.%s", uri, tpl, opt.format)
	sigTxt := path
	if opt.kv != nil {
		if opt.sigKey != "" && opt.kv.Get(KEY_SIG) != "" {
			return nil, ErrKvSig
		}
		sigTxt = fmt.Sprintf("%s?%s", path, opt.kv.Encode())
	} else {
		sigTxt = path + "?"
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

	return domainInfos.makeImageURL(opt.isHttps, path), nil
}

func (p *DomainInfo) makeImageURL(isHttps bool, path string) *ImgUrl {
	proto := HTTP
	if isHttps {
		proto = HTTPS
	}
	return &ImgUrl{
		MainUrl:   fmt.Sprintf("%s://%s%s", proto, p.MainDomain, path),
		BackupUrl: fmt.Sprintf("%s://%s%s", proto, p.BackupDomain, path),
	}
}

/*GetCdnDomainWeights Get the domain weight from imagex
 * @param serviceId 	your serviceId
 */
func (c *ImageXClient) GetCdnDomainWeights(serviceId string) (*GetWeightsResp, error) {
	query := url.Values{}
	query.Set("ServiceId", serviceId)
	query.Set("ProductLine", imagexProductLine)
	respBody, _, err := c.Query("GetCdnDomainWeights", query)
	if err != nil {
		return nil, err
	}

	output := new(GetWeightsResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, err
	}
	return output, nil
}

/*GetDomainInfo Get the domain from imagex
 * @param serviceId 		your serviceId
 * @param fallbackWeights	if get from imagex failed then use it
							key: domain val: weight
*/
func (c *ImageXClient) GetDomainInfo(serviceId string, fallbackWeights map[string]int) (*DomainInfo, error) {
	var cache map[string]int
	var ok bool
	c.Lock.RLock()
	if cache, ok = c.DomainCache[serviceId]; !ok {
		c.Lock.RUnlock()
		c.Lock.Lock()
		if cache, ok = c.DomainCache[serviceId]; !ok {
			var weightsMap map[string]int
			var exist bool
			resp, err := c.GetCdnDomainWeights(serviceId)
			if err != nil {
				weightsMap = fallbackWeights
			}
			if resp != nil {
				if err := resp.ResponseMetadata.Error; err != nil {
					weightsMap = fallbackWeights
				}
				weightsMap, exist = resp.Result[serviceId]
			}
			if !exist || len(weightsMap) == 0 {
				weightsMap = fallbackWeights
			}
			c.DomainCache[serviceId] = weightsMap

			c.Lock.Unlock()
			cache = c.DomainCache[serviceId]

			go func() {
				for range time.Tick(updateInterval * time.Second) {
					var weightsMap map[string]int
					var exist bool
					resp, err := c.GetCdnDomainWeights(serviceId)
					if err != nil {
						weightsMap = fallbackWeights
					}
					if resp != nil {
						if err := resp.ResponseMetadata.Error; err != nil {
							weightsMap = fallbackWeights
						}
						weightsMap, exist = resp.Result[serviceId]
					}
					if !exist || len(weightsMap) == 0 {
						weightsMap = fallbackWeights
					}
					c.Lock.Lock()
					c.DomainCache[serviceId] = weightsMap
					c.Lock.Unlock()
				}
			}()
		} else {
			c.Lock.Unlock()
		}
	} else {
		c.Lock.RUnlock()
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
