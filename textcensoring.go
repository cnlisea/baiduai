package baiduai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// 文本审核

type TextCensoringResponse struct {
	LogId     uint64                       `json:"log_id"`     // 唯一标识符
	Result    *TextCensoringResponseResult `json:"result"`     /// 结果详情
	ErrorCode int64                        `json:"error_code"` // 错误码
	ErrorMsg  string                       `json:"error_msg"`  // 错误信息
}

type TextCensoringResponseResult struct {
	Spam   int                            `json:"spam"`   // 请求中是否包含违禁
	Reject []*TextCensoringResponseVerify `json:"reject"` // 审核未通过的类别列表
	Review []*TextCensoringResponseVerify `json:"review"` // 待人工复审的类别列表
	Pass   []*TextCensoringResponseVerify `json:"pass"`   // 审核通过的类别列表
}

type TextCensoringResponseVerify struct {
	Label int      `json:"label"` // 请求中的违禁类型
	Score float32  `json:"score"` // 违禁检测分, 范围0-1
	Hit   []string `json:"hit"`   // 违禁类型对应命中的违禁词集合, 可能为空
}

func TextCensoring(cfg *Config) (*TextCensoringResponse, error) {
	accessToken, err := AccessToken(cfg)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	b.WriteString("https://aip.baidubce.com/rest/2.0/antispam/v2/spam?access_token=")
	b.WriteString(accessToken.AccessToken)

	res, err := http.Post(b.String(), "application/x-www-form-urlencoded", strings.NewReader("content="+cfg.Content))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret := new(TextCensoringResponse)
	if err = json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	if ret.ErrorCode != 0 {
		return nil, errors.New(ret.ErrorMsg)
	}

	return ret, nil
}
