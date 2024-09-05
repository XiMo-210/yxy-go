package yxyClient

import (
	"yxy-go/pkg/xerr"

	"github.com/go-resty/resty/v2"
)

func HttpSendPost(url string, req map[string]interface{}, headers map[string]string, resp interface{}) (*resty.Response, error) {
	client := resty.New()

	r, err := client.R().
		SetHeaders(headers).
		SetBody(req).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	return r, nil
}