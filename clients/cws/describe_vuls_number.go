package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看当前漏洞总计数量
// https://cloud.tencent.com/document/api/692/18088

type DescribeVulsNumberRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeVulsNumberRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulsNumberResponse, error) {
	resp := &DescribeVulsNumberResponse{}
	err := client.Request("cws", "DescribeVulsNumber", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeVulsNumberResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 受影响的网站总数。
	ImpactSiteNumber int64 `json:"ImpactSiteNumber"`
	// 受影响的网站列表。
	ImpactSites []*MonitorMiniSite `json:"ImpactSites"`
	// 扫描页面总数。
	PageCount int64 `json:"PageCount"`
	// 已验证的网站总数。
	SiteNumber int64 `json:"SiteNumber"`
	// 已验证的网站列表。
	Sites []*MonitorMiniSite `json:"Sites"`
	// 高风险漏洞总数。
	VulsHighNumber int64 `json:"VulsHighNumber"`
	// 低高风险漏洞总数。
	VulsLowNumber int64 `json:"VulsLowNumber"`
	// 中风险漏洞总数。
	VulsMiddleNumber int64 `json:"VulsMiddleNumber"`
	// 风险提示总数。
	VulsNoticeNumber int64 `json:"VulsNoticeNumber"`
}
