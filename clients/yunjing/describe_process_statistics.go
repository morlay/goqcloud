package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取进程统计列表
// https://cloud.tencent.com/document/api/296/30334

type DescribeProcessStatisticsRequest struct {
	// 过滤条件。ProcessName - String - 是否必填：否 - 进程名
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeProcessStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProcessStatisticsResponse, error) {
	resp := &DescribeProcessStatisticsResponse{}
	err := client.Request("yunjing", "DescribeProcessStatistics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeProcessStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 进程统计列表数据数组。
	ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics"`
	// 进程统计列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
