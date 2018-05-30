package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 验证站点
// https://cloud.tencent.com/document/api/692/16755
type VerifySitesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 站点的url列表
	Urls []*string `name:"Urls"`
}

func (req *VerifySitesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*VerifySitesResponse, error) {
	resp := &VerifySitesResponse{}
	err := client.Request("cws", "VerifySites", "2018-03-12").Do(req, resp)
	return resp, err
}

type VerifySitesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 验证失败的根域名数量。
	FailNumber int64 `json:"FailNumber"`
	// 验证成功的根域名数量。
	SuccessNumber int64 `json:"SuccessNumber"`
}
