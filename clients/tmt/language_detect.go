package tmt

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 语种识别
// https://cloud.tencent.com/document/api/551/15620
type LanguageDetectRequest struct {
	// 项目id
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
	// 待识别的文本
	Text string `name:"Text"`
}

func (req *LanguageDetectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*LanguageDetectResponse, error) {
	resp := &LanguageDetectResponse{}
	err := client.Request("tmt", "LanguageDetect", "2018-03-21").Do(req, resp)
	return resp, err
}

type LanguageDetectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 识别出的语言种类，参考语言列表 zh : 中文   en : 英文  jp : 日语   kr : 韩语  de : 德语  fr : 法语  es : 西班牙文   it : 意大利文  tr : 土耳其文  ru : 俄文  pt : 葡萄牙文  vi : 越南文  id : 印度尼西亚文  ms : 马来西亚文  th : 泰文
	Lang string `json:"Lang"`
}
