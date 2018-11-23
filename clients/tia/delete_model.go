package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除Model
// https://cloud.tencent.com/document/api/851/18314

type DeleteModelRequest struct {
	// 要删除的模型所在的集群名称
	Cluster *string `name:"Cluster,omitempty"`
	// 要删除的模型名称
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
	// 模型类型
	ServType *string `name:"ServType,omitempty"`
}

func (req *DeleteModelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteModelResponse, error) {
	resp := &DeleteModelResponse{}
	err := client.Request("tia", "DeleteModel", "2018-02-26").Do(req, resp)
	return resp, err
}

type DeleteModelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
