package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取订单信息
// https://cloud.tencent.com/document/api/409/18104

type DescribeOrdersRequest struct {
	// 订单名集合
	DealNames []*string `name:"DealNames"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeOrdersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOrdersResponse, error) {
	resp := &DescribeOrdersResponse{}
	err := client.Request("postgres", "DescribeOrders", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeOrdersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单数组
	Deals []*PgDeal `json:"Deals"`
	// 订单数量
	TotalCount int64 `json:"TotalCount"`
}
