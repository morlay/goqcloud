package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建默认VPC和默认子网
// https://cloud.tencent.com/document/api/215/17876

type CreateDefaultVpcRequest struct {
	// 是否强制返回默认VPC
	Force *bool `name:"Force,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 子网所在的可用区ID，不指定将随机选择可用区
	Zone *string `name:"Zone,omitempty"`
}

func (req *CreateDefaultVpcRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDefaultVpcResponse, error) {
	resp := &CreateDefaultVpcResponse{}
	err := client.Request("vpc", "CreateDefaultVpc", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDefaultVpcResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 默认VPC和子网ID
	Vpc DefaultVpcSubnet `json:"Vpc"`
}
