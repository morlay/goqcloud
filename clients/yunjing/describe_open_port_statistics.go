package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取端口统计列表
// https://cloud.tencent.com/document/api/296/30328

type DescribeOpenPortStatisticsRequest struct {
	// 过滤条件。Port - Uint64 - 是否必填：否 - 端口号
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeOpenPortStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOpenPortStatisticsResponse, error) {
	resp := &DescribeOpenPortStatisticsResponse{}
	err := client.Request("yunjing", "DescribeOpenPortStatistics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeOpenPortStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 端口统计数据列表
	OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics"`
	// 端口统计列表总数
	TotalCount int64 `json:"TotalCount"`
}
