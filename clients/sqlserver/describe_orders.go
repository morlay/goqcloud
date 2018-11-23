package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询订单信息
// https://cloud.tencent.com/document/api/238/19937

type DescribeOrdersRequest struct {
	// 订单数组。发货时会返回订单名字，利用该订单名字调用DescribeOrders接口查询发货情况
	DealNames []*string `name:"DealNames"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeOrdersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOrdersResponse, error) {
	resp := &DescribeOrdersResponse{}
	err := client.Request("sqlserver", "DescribeOrders", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeOrdersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单信息数组
	Deals []*DealInfo `json:"Deals"`
	// 返回多少个订单的信息
	TotalCount int64 `json:"TotalCount"`
}
