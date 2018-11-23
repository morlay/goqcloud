package billing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取账户余额
// https://cloud.tencent.com/document/api/555/20253

type DescribeAccountBalanceRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountBalanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountBalanceResponse, error) {
	resp := &DescribeAccountBalanceResponse{}
	err := client.Request("billing", "DescribeAccountBalance", "2018-07-09").Do(req, resp)
	return resp, err
}

type DescribeAccountBalanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云账户信息中的”展示可用余额”字段，单位为"分"
	Balance int64 `json:"Balance"`
}
