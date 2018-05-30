package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除计算环境
// https://cloud.tencent.com/document/api/599/15889
type DeleteComputeEnvRequest struct {
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteComputeEnvRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteComputeEnvResponse, error) {
	resp := &DeleteComputeEnvResponse{}
	err := client.Request("batch", "DeleteComputeEnv", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteComputeEnvResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
