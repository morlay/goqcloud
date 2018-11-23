package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 解除标签与预授权规则的绑定
// https://cloud.tencent.com/document/api/386/30139

type UnbindPsaTagRequest struct {
	// 预授权规则ID
	PsaId string `name:"PsaId"`
	// 区域
	Region string `name:"Region"`
	// 需要解绑的标签key
	TagKey string `name:"TagKey"`
	// 需要解绑的标签value
	TagValue string `name:"TagValue"`
}

func (req *UnbindPsaTagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UnbindPsaTagResponse, error) {
	resp := &UnbindPsaTagResponse{}
	err := client.Request("bm", "UnbindPsaTag", "2018-04-23").Do(req, resp)
	return resp, err
}

type UnbindPsaTagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
