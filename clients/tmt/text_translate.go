package tmt

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 文本翻译
// https://cloud.tencent.com/document/api/551/15619

type TextTranslateRequest struct {
	// 项目id
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
	// 源语言，参照Target支持语言列表
	Source string `name:"Source"`
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本会翻译失败
	SourceText string `name:"SourceText"`
	// 目标语言，参照支持语言列表 zh : 中文   en : 英文  jp : 日语   kr : 韩语  de : 德语  fr : 法语  es : 西班牙文   it : 意大利文  tr : 土耳其文  ru : 俄文  pt : 葡萄牙文  vi : 越南文  id : 印度尼西亚文  ms : 马来西亚文  th : 泰文  auto : 自动识别源语言，只能用于source字段
	Target string `name:"Target"`
}

func (req *TextTranslateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TextTranslateResponse, error) {
	resp := &TextTranslateResponse{}
	err := client.Request("tmt", "TextTranslate", "2018-03-21").Do(req, resp)
	return resp, err
}

type TextTranslateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 源语言，详见入参Target
	Source string `json:"Source"`
	// 目标语言，详见入参Target
	Target string `json:"Target"`
	// 翻译后的文本
	TargetText string `json:"TargetText"`
}
