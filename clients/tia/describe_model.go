package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 描述Model
// https://cloud.tencent.com/document/api/851/18313

type DescribeModelRequest struct {
	// 模型所在集群名称
	Cluster *string `name:"Cluster,omitempty"`
	// 模型名称
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
	// 模型类型
	ServType *string `name:"ServType,omitempty"`
}

func (req *DescribeModelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeModelResponse, error) {
	resp := &DescribeModelResponse{}
	err := client.Request("tia", "DescribeModel", "2018-02-26").Do(req, resp)
	return resp, err
}

type DescribeModelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 模型信息
	Model Model `json:"Model"`
}
