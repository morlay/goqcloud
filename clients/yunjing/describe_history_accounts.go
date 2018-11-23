package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取帐号变更历史列表
// https://cloud.tencent.com/document/api/296/30335

type DescribeHistoryAccountsRequest struct {
	// 过滤条件。Username - String - 是否必填：否 - 帐号名
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。
	Uuid string `name:"Uuid"`
}

func (req *DescribeHistoryAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeHistoryAccountsResponse, error) {
	resp := &DescribeHistoryAccountsResponse{}
	err := client.Request("yunjing", "DescribeHistoryAccounts", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeHistoryAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 帐号变更历史数据数组。
	HistoryAccounts []*HistoryAccount `json:"HistoryAccounts"`
	// 帐号变更历史列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
