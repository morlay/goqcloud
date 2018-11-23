package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看漏洞数随时间变化统计信息
// https://cloud.tencent.com/document/api/692/18087

type DescribeVulsNumberTimelineRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeVulsNumberTimelineRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulsNumberTimelineResponse, error) {
	resp := &DescribeVulsNumberTimelineResponse{}
	err := client.Request("cws", "DescribeVulsNumberTimeline", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeVulsNumberTimelineResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 统计数据记录数量。
	TotalCount int64 `json:"TotalCount"`
	// 用户漏洞数随时间变化统计数据。
	VulsTimeline []*VulsTimeline `json:"VulsTimeline"`
}
