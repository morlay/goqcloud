package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取概览统计数据
// https://cloud.tencent.com/document/api/296/19849

type DescribeOverviewStatisticsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeOverviewStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOverviewStatisticsResponse, error) {
	resp := &DescribeOverviewStatisticsResponse{}
	err := client.Request("yunjing", "DescribeOverviewStatistics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeOverviewStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全基线数。
	BaseLineNum int64 `json:"BaseLineNum"`
	// 暴力破解成功数。
	BruteAttackSuccessNum int64 `json:"BruteAttackSuccessNum"`
	// 木马文件数。
	MalwareNum int64 `json:"MalwareNum"`
	// 异地登录数。
	NonlocalLoginNum int64 `json:"NonlocalLoginNum"`
	// 服务器在线数。
	OnlineMachineNum int64 `json:"OnlineMachineNum"`
	// 专业服务器数。
	ProVersionMachineNum int64 `json:"ProVersionMachineNum"`
	// 漏洞数。
	VulNum int64 `json:"VulNum"`
}
