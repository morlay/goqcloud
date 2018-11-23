package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询订单信息
// https://cloud.tencent.com/document/api/239/30602

type DescribeInstanceDealDetailRequest struct {
	// 订单ID数组
	DealIds []*string `name:"DealIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceDealDetailRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceDealDetailResponse, error) {
	resp := &DescribeInstanceDealDetailResponse{}
	err := client.Request("redis", "DescribeInstanceDealDetail", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceDealDetailResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单详细信息
	DealDetails []*TradeDealDetail `json:"DealDetails"`
}
