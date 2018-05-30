package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看漏洞列表
// https://cloud.tencent.com/document/api/692/16741
type DescribeVulsRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 监控任务ID
	MonitorId *int64 `name:"MonitorId,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 站点ID
	SiteId *int64 `name:"SiteId,omitempty"`
}

func (req *DescribeVulsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulsResponse, error) {
	resp := &DescribeVulsResponse{}
	err := client.Request("cws", "DescribeVuls", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeVulsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 漏洞数量。
	TotalCount int64 `json:"TotalCount"`
	// 漏洞信息列表。
	Vuls []*Vul `json:"Vuls"`
}
