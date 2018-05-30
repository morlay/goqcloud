package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询订单信息
// https://cloud.tencent.com/document/api/557/16139
type DescribeOrdersRequest struct {
	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `name:"DealNames"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeOrdersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOrdersResponse, error) {
	resp := &DescribeOrdersResponse{}
	err := client.Request("dcdb", "DescribeOrders", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeOrdersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单信息列表。
	Deals []*Deal `json:"Deals"`
	// 返回的订单数量。
	TotalCount []*int64 `json:"TotalCount"`
}
