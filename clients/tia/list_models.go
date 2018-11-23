package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 列举Model
// https://cloud.tencent.com/document/api/851/18312

type ListModelsRequest struct {
	// 部署模型的集群
	Cluster *string `name:"Cluster,omitempty"`
	// 分页参数，返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 分页参数，起始位置
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 模型类型
	ServType *string `name:"ServType,omitempty"`
}

func (req *ListModelsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ListModelsResponse, error) {
	resp := &ListModelsResponse{}
	err := client.Request("tia", "ListModels", "2018-02-26").Do(req, resp)
	return resp, err
}

type ListModelsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Model数组，用以显示所有模型的信息
	Models []*Model `json:"Models"`
}
