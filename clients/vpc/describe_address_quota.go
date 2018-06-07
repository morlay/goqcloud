package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询弹性公网IP配额
// https://cloud.tencent.com/document/api/215/16701

type DescribeAddressQuotaRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAddressQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAddressQuotaResponse, error) {
	resp := &DescribeAddressQuotaResponse{}
	err := client.Request("vpc", "DescribeAddressQuota", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAddressQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 账户 EIP 配额信息。
	QuotaSet []*Quota `json:"QuotaSet"`
}
