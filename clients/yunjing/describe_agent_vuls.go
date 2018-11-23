package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取单台主机的漏洞列表
// https://cloud.tencent.com/document/api/296/19861

type DescribeAgentVulsRequest struct {
	// 过滤条件。Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 客户端UUID。
	Uuid string `name:"Uuid"`
	// 漏洞类型。WEB: Web应用漏洞SYSTEM：系统组件漏洞BASELINE：安全基线
	VulType string `name:"VulType"`
}

func (req *DescribeAgentVulsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAgentVulsResponse, error) {
	resp := &DescribeAgentVulsResponse{}
	err := client.Request("yunjing", "DescribeAgentVuls", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeAgentVulsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 主机漏洞信息
	AgentVuls []*AgentVul `json:"AgentVuls"`
	// 记录总数
	TotalCount int64 `json:"TotalCount"`
}
