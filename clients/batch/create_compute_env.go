package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建计算环境
// https://cloud.tencent.com/document/api/599/15891

type CreateComputeEnvRequest struct {
	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 计算环境信息
	ComputeEnv NamedComputeEnv `name:"ComputeEnv"`
	// 位置信息
	Placement Placement `name:"Placement"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateComputeEnvRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateComputeEnvResponse, error) {
	resp := &CreateComputeEnvResponse{}
	err := client.Request("batch", "CreateComputeEnv", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateComputeEnvResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 计算环境ID
	EnvId string `json:"EnvId"`
}
