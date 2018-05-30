package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 销毁计算节点
// https://cloud.tencent.com/document/api/599/15895
type TerminateComputeNodeRequest struct {
	// 计算节点ID
	ComputeNodeId string `name:"ComputeNodeId"`
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *TerminateComputeNodeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TerminateComputeNodeResponse, error) {
	resp := &TerminateComputeNodeResponse{}
	err := client.Request("batch", "TerminateComputeNode", "2017-03-12").Do(req, resp)
	return resp, err
}

type TerminateComputeNodeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
