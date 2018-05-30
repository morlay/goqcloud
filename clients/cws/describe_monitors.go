package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看监控任务
// https://cloud.tencent.com/document/api/692/16745
type DescribeMonitorsRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 监控任务ID列表
	MonitorIds []*int64 `name:"MonitorIds,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeMonitorsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMonitorsResponse, error) {
	resp := &DescribeMonitorsResponse{}
	err := client.Request("cws", "DescribeMonitors", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeMonitorsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 监控任务列表。
	Monitors []*MonitorsDetail `json:"Monitors"`
	// 监控任务数量。
	TotalCount int64 `json:"TotalCount"`
}
