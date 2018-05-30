package tmt

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 语音翻译
// https://cloud.tencent.com/document/api/551/16611
type SpeechTranslateRequest struct {
	// pcm : 146   amr : 33554432   mp3 : 83886080
	AudioFormat int64 `name:"AudioFormat"`
	// 语音分片内容的base64字符串
	Data string `name:"Data"`
	// 是否最后一片
	IsEnd int64 `name:"IsEnd"`
	// 项目id
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 语音分片后的第几片
	Seq int64 `name:"Seq"`
	// 一段完整的语音对应一个SessionUuid
	SessionUuid string `name:"SessionUuid"`
	// 音频中的语言类型，支持语言列表 zh : 中文   en : 英文
	Source string `name:"Source"`
	// 翻译目标语⾔言类型 ，支持的语言列表 zh : 中文   en : 英文
	Target string `name:"Target"`
}

func (req *SpeechTranslateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SpeechTranslateResponse, error) {
	resp := &SpeechTranslateResponse{}
	err := client.Request("tmt", "SpeechTranslate", "2018-03-21").Do(req, resp)
	return resp, err
}

type SpeechTranslateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 语音识别状态 1-进行中 0-完成
	RecognizeStatus int64 `json:"RecognizeStatus"`
	// 第几个语音分片
	Seq int64 `json:"Seq"`
	// 请求的SessionUuid直接返回
	SessionUuid string `json:"SessionUuid"`
	// 源语言
	Source string `json:"Source"`
	// 识别出的源文
	SourceText string `json:"SourceText"`
	// 目标语言
	Target string `json:"Target"`
	// 翻译出的译文
	TargetText string `json:"TargetText"`
}
