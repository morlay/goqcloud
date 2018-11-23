package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取漏洞详情
// https://cloud.tencent.com/document/api/296/19859

type DescribeVulInfoRequest struct {
	// 区域
	Region string `name:"Region"`
	// 漏洞种类ID。
	VulId int64 `name:"VulId"`
}

func (req *DescribeVulInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulInfoResponse, error) {
	resp := &DescribeVulInfoResponse{}
	err := client.Request("yunjing", "DescribeVulInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeVulInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 漏洞CVE。
	CveId string `json:"CveId"`
	// 漏洞描述。
	Description string `json:"Description"`
	// 参考链接。
	Reference string `json:"Reference"`
	// 修复方案。
	RepairPlan string `json:"RepairPlan"`
	// 漏洞种类ID。
	VulId int64 `json:"VulId"`
	// 漏洞等级。
	VulLevel string `json:"VulLevel"`
	// 漏洞名称。
	VulName string `json:"VulName"`
	// 漏洞类型。
	VulType string `json:"VulType"`
}
