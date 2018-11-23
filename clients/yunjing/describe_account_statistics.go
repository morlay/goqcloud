package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取帐号统计列表数据
// https://cloud.tencent.com/document/api/296/30340

type DescribeAccountStatisticsRequest struct {
	// 过滤条件。Username - String - 是否必填：否 - 帐号用户名
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountStatisticsResponse, error) {
	resp := &DescribeAccountStatisticsResponse{}
	err := client.Request("yunjing", "DescribeAccountStatistics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeAccountStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 帐号统计列表。
	AccountStatistics []*AccountStatistics `json:"AccountStatistics"`
	// 帐号统计列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
