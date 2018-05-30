package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取计算环境详情
// https://cloud.tencent.com/document/api/599/15892
type DescribeComputeEnvRequest struct {
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComputeEnvRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComputeEnvResponse, error) {
	resp := &DescribeComputeEnvResponse{}
	err := client.Request("batch", "DescribeComputeEnv", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeComputeEnvResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 计算节点统计指标
	ComputeNodeMetrics ComputeNodeMetrics `json:"ComputeNodeMetrics"`
	// 计算节点列表信息
	ComputeNodeSet []*ComputeNode `json:"ComputeNodeSet"`
	// 计算环境创建时间
	CreateTime string `json:"CreateTime"`
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `json:"DesiredComputeNodeCount"`
	// 计算环境ID
	EnvId string `json:"EnvId"`
	// 计算环境名称
	EnvName string `json:"EnvName"`
	// 计算环境类型
	EnvType string `json:"EnvType"`
	// 位置信息
	Placement Placement `json:"Placement"`
}
