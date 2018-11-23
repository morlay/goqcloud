package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询带宽包配额
// https://cloud.tencent.com/document/api/215/19210

type DescribeBandwidthPackageQuotaRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBandwidthPackageQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBandwidthPackageQuotaResponse, error) {
	resp := &DescribeBandwidthPackageQuotaResponse{}
	err := client.Request("vpc", "DescribeBandwidthPackageQuota", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeBandwidthPackageQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 带宽包配额数据结构
	QuotaSet []*Quota `json:"QuotaSet"`
}
