package imagex

import (
	"errors"
	"net/url"
)

const (
	FORMAT_JPEG     = "jpeg"
	FORMAT_PNG      = "png"
	FORMAT_WEBP     = "webp"
	FORMAT_AWEBP    = "awebp"
	FORMAT_GIF      = "gif"
	FORMAT_HEIC     = "heic"
	FORMAT_ORIGINAL = "image"

	HTTP  = "http"
	HTTPS = "https"

	KEY_SIG = "sig"
)

var (
	ErrKvSig = errors.New("Input kv already has sig query")
)

type option struct {
	isHttps         bool
	format          string
	sigKey          string
	parm            string
	kv              url.Values
	fallbackWeights map[string]int
}

type OptionFun func(*option)

//WithHttps if you need https
func WithHttps() OptionFun {
	return func(opt *option) {
		opt.isHttps = true
	}
}

//WithFormat transcode to the image in the desired format
func WithFormat(format string) OptionFun {
	return func(opt *option) {
		opt.format = format
	}
}

//WithSig used  if your tpl need authentication, the key is your template authentication key
func WithSig(key string) OptionFun {
	return func(opt *option) {
		opt.sigKey = key
	}
}

//WithKV used kv in querystring, the key must not use "sig" if your tpl need authentication
func WithKV(kv url.Values) OptionFun {
	return func(opt *option) {
		opt.kv = kv
	}
}

//WithFallBackWeights used if the domain weight get from imagex failed
func WithFallBackWeights(w map[string]int) OptionFun {
	return func(opt *option) {
		opt.fallbackWeights = w
	}
}
