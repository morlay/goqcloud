package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量销毁计算节点
// https://cloud.tencent.com/document/api/599/17372
type TerminateComputeNodesRequest struct {
	// 计算节点ID列表
	ComputeNodeIds []*string `name:"ComputeNodeIds"`
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *TerminateComputeNodesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TerminateComputeNodesResponse, error) {
	resp := &TerminateComputeNodesResponse{}
	err := client.Request("batch", "TerminateComputeNodes", "2017-03-12").Do(req, resp)
	return resp, err
}

type TerminateComputeNodesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
