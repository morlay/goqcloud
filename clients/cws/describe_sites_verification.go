package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看站点列表的验证信息列表
// https://cloud.tencent.com/document/api/692/16753

type DescribeSitesVerificationRequest struct {
	// 区域
	Region string `name:"Region"`
	// 站点的url列表
	Urls []*string `name:"Urls"`
}

func (req *DescribeSitesVerificationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSitesVerificationResponse, error) {
	resp := &DescribeSitesVerificationResponse{}
	err := client.Request("cws", "DescribeSitesVerification", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeSitesVerificationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 验证信息列表。
	SitesVerification []*SitesVerification `json:"SitesVerification"`
	// 验证信息数量。
	TotalCount int64 `json:"TotalCount"`
}
