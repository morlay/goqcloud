package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取漏洞受影响机器列表
// https://cloud.tencent.com/document/api/296/19860

type DescribeImpactedHostsRequest struct {
	// 过滤条件。Status - String - 是否必填：否 - 状态筛选（UN_OPERATED：待处理 | FIXED：已修复）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 漏洞种类ID。
	VulId int64 `name:"VulId"`
}

func (req *DescribeImpactedHostsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeImpactedHostsResponse, error) {
	resp := &DescribeImpactedHostsResponse{}
	err := client.Request("yunjing", "DescribeImpactedHosts", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeImpactedHostsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 漏洞影响机器列表数组
	ImpactedHosts []*ImpactedHost `json:"ImpactedHosts"`
	// 记录总数
	TotalCount int64 `json:"TotalCount"`
}
