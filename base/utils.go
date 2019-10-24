package base

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	// 初始化随机种子
	rand.Seed(time.Now().Unix())
}

func CreateTempAKSK() (accessKeyId string, plainSk string, err error) {
	// 生成AccessKeyId
	if accessKeyId, err = GenerateAccessKeyId("AKTP"); err != nil {
		return
	}

	// 生成SecretKey明文
	plainSk, err = GenerateSecretKey()
	if err != nil {
		return
	}
	return
}

func GenerateAccessKeyId(prefix string) (string, error) {
	// 生成uuid，如：a1fe1d4f-eb56-4a06-86e8-3e5068a1a838
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	// 滤掉'-'后，做base64，输出：YTFmZTFkNGZlYjU2NGEwNjg2ZTgzZTUwNjhhMWE4Mzg=
	uidBase64 := base64.StdEncoding.EncodeToString([]byte(strings.Replace(uid.String(), "-", "", -1)))

	// 去掉"-+/="特殊字符，加上prefix
	s := strings.Replace(uidBase64, "=", "", -1)
	s = strings.Replace(s, "/", "", -1)
	s = strings.Replace(s, "+", "", -1)
	s = strings.Replace(s, "-", "", -1)
	return prefix + s, nil
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateSecretKey() (string, error) {
	randString32 := randStringRunes(32)
	return AesEncryptCBCWithBase64([]byte(randString32), []byte("ttcloudbestcloud"))
}

/*
func CreateInnerToken(credential Credentials, secretAccessKey string, inlinePolicy *Policy) (*InnerToken, error) {
	var err error
	innerToken := new(InnerToken)
	innerToken.LTAccessKeyId = credential.AccessKeyID
	key := md5.Sum([]byte(credential.SecretAccessKey))
	innerToken.SignedSecretAccessKey, err = AesEncryptCBCWithBase64([]byte(secretAccessKey), key[:])
	innerToken.Policy = inlinePolicy
	if err == nil {
		return innerToken, nil
	} else {
		return nil, err
	}
}
*/

func CreateInnerToken(credentials Credentials, sts *SecurityToken2, inlinePolicy *Policy, t int64) (*InnerToken, error) {
	var err error
	innerToken := new(InnerToken)

	innerToken.LTAccessKeyId = credentials.AccessKeyID
	innerToken.AccessKeyId = sts.AccessKeyId
	innerToken.ExpiredTime = t

	key := md5.Sum([]byte(credentials.SecretAccessKey))
	innerToken.SignedSecretAccessKey, err = AesEncryptCBCWithBase64([]byte(sts.SecretAccessKey), key[:])
	if err != nil {
		return nil, err
	}

	if inlinePolicy != nil {
		b, _ := json.Marshal(inlinePolicy)
		innerToken.PolicyString = string(b)
	}

	// sign signature
	signStr := fmt.Sprintf("%s|%s|%d|%s|%s", innerToken.LTAccessKeyId, innerToken.AccessKeyId, innerToken.ExpiredTime, innerToken.SignedSecretAccessKey, innerToken.PolicyString)

	innerToken.Signature = hex.EncodeToString(hmacSHA256(key[:], signStr))
	return innerToken, nil
}

func getTimeout(serviceTimeout, apiTimeout time.Duration) time.Duration {
	timeout := time.Second
	if serviceTimeout != time.Duration(0) {
		timeout = serviceTimeout
	}
	if apiTimeout != time.Duration(0) {
		timeout = apiTimeout
	}
	return timeout
}

func mergeQuery(query1, query2 url.Values) (query url.Values) {
	query = url.Values{}
	if query1 != nil {
		for k, vv := range query1 {
			for _, v := range vv {
				query.Add(k, v)
			}
		}
	}

	if query2 != nil {
		for k, vv := range query2 {
			for _, v := range vv {
				query.Add(k, v)
			}
		}
	}
	return
}

func mergeHeader(header1, header2 http.Header) (header http.Header) {
	header = http.Header{}
	if header1 != nil {
		for k, v := range header1 {
			header.Set(k, strings.Join(v, ";"))
		}
	}
	if header2 != nil {
		for k, v := range header2 {
			header.Set(k, strings.Join(v, ";"))
		}
	}

	return
}

func NewAllowStatement(actions, resources []string) *Statement {
	sts := new(Statement)
	sts.Effect = "Allow"
	sts.Action = actions
	sts.Resource = resources

	return sts
}

func NewDenyStatement(actions, resources []string) *Statement {
	sts := new(Statement)
	sts.Effect = "Deny"
	sts.Action = actions
	sts.Resource = resources

	return sts
}
