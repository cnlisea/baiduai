package baiduai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type AccessTokenResponse struct {
	ExpiresIn        int64  `json:"expires_in"`        // 过期时间
	AccessToken      string `json:"access_token"`      // 访问码
	Error            string `json:"error"`             // 错误码
	ErrorDescription string `json:"error_description"` // 错误信息
}

func AccessToken(cfg *Config) (*AccessTokenResponse, error) {
	var b bytes.Buffer

	b.WriteString("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=")
	b.WriteString(cfg.ApiKey)
	b.WriteString("&client_secret=")
	b.WriteString(cfg.SecretKey)

	res, err := http.Post(b.String(), "application/json; charset=UTF-8", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret := new(AccessTokenResponse)
	if err = json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	if ret.Error != "" {
		return nil, errors.New(ret.Error + ": " + ret.ErrorDescription)
	}

	return ret, nil
}
