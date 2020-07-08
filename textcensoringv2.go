package baiduai

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type TextCensoringV2Response struct {
	LogId          int                            `json:"log_id"`         // 唯一标识符
	Data           []*TextCensoringV2ResponseData `json:"data"`           // 结果详情
	Conclusion     string                         `json:"conclusion"`     // 审核结果，可取值：合规、不合规、疑似、审核失败
	ConclusionType int                            `json:"conclusionType"` // 审核结果类型，可取值1.合规，2.不合规，3.疑似，4.审核失败
	ErrorCode      int                            `json:"error_code"`     // 错误码
	ErrorMsg       string                         `json:"error_msg"`      // 错误信息
}

type TextCensoringV2ResponseData struct {
	Type           int                               `json:"type"`           // 审核主类型，11：百度官方违禁词库、12：文本反作弊、13:自定义文本黑名单、14:自定义文本白名单
	SubType        int                               `json:"subType"`        // 审核子类型
	Conclusion     string                            `json:"conclusion"`     // 审核结果，可取值：合规、不合规、疑似、审核失败
	ConclusionType int                               `json:"conclusionType"` // 审核结果类型，可取值1.合规，2.不合规，3.疑似，4.审核失败
	Msg            string                            `json:"msg"`            // 不合规项描述信息
	Hits           []*TextCensoringV2ResponseDataHit `json:"hits"`           // 命中关键词信息
}

type TextCensoringV2ResponseDataHit struct {
	Probability float32  `json:"probability"` // 相似度
	DatasetName string   `json:"datasetName"` // 违规项目所属数据集名称
	Words       []string `json:"words"`       // 违规文本关键字
}

func TextCensoringV2(cfg *Config) (*TextCensoringV2Response, error) {
	accessToken, err := AccessToken(cfg)
	if err != nil {
		return nil, err
	}

	res, err := http.Post("https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token="+accessToken.AccessToken, "application/x-www-form-urlencoded", strings.NewReader("text="+cfg.Content))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret := new(TextCensoringV2Response)
	if err = json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	if ret.ErrorCode != 0 {
		return nil, errors.New(ret.ErrorMsg)
	}

	return ret, nil
}
