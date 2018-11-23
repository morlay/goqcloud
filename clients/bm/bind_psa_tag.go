package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 为预授权规则绑定标签
// https://cloud.tencent.com/document/api/386/30144

type BindPsaTagRequest struct {
	// 预授权规则ID
	PsaId string `name:"PsaId"`
	// 区域
	Region string `name:"Region"`
	// 需要绑定的标签key
	TagKey string `name:"TagKey"`
	// 需要绑定的标签value
	TagValue string `name:"TagValue"`
}

func (req *BindPsaTagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*BindPsaTagResponse, error) {
	resp := &BindPsaTagResponse{}
	err := client.Request("bm", "BindPsaTag", "2018-04-23").Do(req, resp)
	return resp, err
}

type BindPsaTagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
