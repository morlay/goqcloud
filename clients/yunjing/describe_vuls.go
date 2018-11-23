package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取漏洞列表
// https://cloud.tencent.com/document/api/296/19857

type DescribeVulsRequest struct {
	// 过滤条件。Status - String - 是否必填：否 - 状态筛选（UN_OPERATED: 待处理 | FIXED：已修复）Status过滤条件值只能取其一，不能是“或”逻辑。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 漏洞类型。WEB：Web应用漏洞SYSTEM：系统组件漏洞BASELINE：安全基线
	VulType string `name:"VulType"`
}

func (req *DescribeVulsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulsResponse, error) {
	resp := &DescribeVulsResponse{}
	err := client.Request("yunjing", "DescribeVuls", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeVulsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 漏洞数量。
	TotalCount int64 `json:"TotalCount"`
	// 漏洞列表数组。
	Vuls []*Vul `json:"Vuls"`
}
