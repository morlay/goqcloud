package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取漏洞检测结果
// https://cloud.tencent.com/document/api/296/19858

type DescribeVulScanResultRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeVulScanResultRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVulScanResultResponse, error) {
	resp := &DescribeVulScanResultResponse{}
	err := client.Request("yunjing", "DescribeVulScanResult", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeVulScanResultResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 受影响的专业版主机数。
	ImpactedHostNum int64 `json:"ImpactedHostNum"`
	// 专业版机器数。
	ProVersionNum int64 `json:"ProVersionNum"`
	// 漏洞数量。
	VulNum int64 `json:"VulNum"`
}
