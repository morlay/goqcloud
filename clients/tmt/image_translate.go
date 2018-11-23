package tmt

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 图片翻译
// https://cloud.tencent.com/document/api/551/17232

type ImageTranslateRequest struct {
	// 图片数据的Base64字符串
	Data string `name:"Data"`
	// 项目id
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
	// doc:文档扫描
	Scene string `name:"Scene"`
	// 唯一id，返回时原样返回
	SessionUuid string `name:"SessionUuid"`
	// 源语言，支持语言列表 zh : 中文   en : 英文
	Source string `name:"Source"`
	// 目标语言，支持语言列表 zh : 中文   en : 英文
	Target string `name:"Target"`
}

func (req *ImageTranslateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ImageTranslateResponse, error) {
	resp := &ImageTranslateResponse{}
	err := client.Request("tmt", "ImageTranslate", "2018-03-21").Do(req, resp)
	return resp, err
}

type ImageTranslateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 图片翻译结果，翻译结果按识别的文本每一行独立翻译，后续会推出按段落划分并翻译的版本
	ImageRecord ImageRecord `json:"ImageRecord"`
	// 请求的SessionUuid返回
	SessionUuid string `json:"SessionUuid"`
	// 源语言
	Source string `json:"Source"`
	// 目标语言
	Target string `json:"Target"`
}
