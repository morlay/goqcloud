package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改计算环境
// https://cloud.tencent.com/document/api/599/15890
type ModifyComputeEnvRequest struct {
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `name:"DesiredComputeNodeCount"`
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyComputeEnvRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyComputeEnvResponse, error) {
	resp := &ModifyComputeEnvResponse{}
	err := client.Request("batch", "ModifyComputeEnv", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyComputeEnvResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
