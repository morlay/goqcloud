package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取组件统计列表
// https://cloud.tencent.com/document/api/296/30337

type DescribeComponentStatisticsRequest struct {
	// 过滤条件。ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComponentStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComponentStatisticsResponse, error) {
	resp := &DescribeComponentStatisticsResponse{}
	err := client.Request("yunjing", "DescribeComponentStatistics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeComponentStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 组件统计列表数据数组。
	ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics"`
	// 组件统计列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
