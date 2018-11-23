package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建预授权规则
// https://cloud.tencent.com/document/api/386/30143

type CreatePsaRegulationRequest struct {
	// 规则备注
	PsaDescription *string `name:"PsaDescription,omitempty"`
	// 规则别名
	PsaName string `name:"PsaName"`
	// 区域
	Region string `name:"Region"`
	// 维修实例上限，默认为5
	RepairLimit *int64 `name:"RepairLimit,omitempty"`
	// 关联的故障类型ID列表
	TaskTypeIds []*int64 `name:"TaskTypeIds"`
}

func (req *CreatePsaRegulationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreatePsaRegulationResponse, error) {
	resp := &CreatePsaRegulationResponse{}
	err := client.Request("bm", "CreatePsaRegulation", "2018-04-23").Do(req, resp)
	return resp, err
}

type CreatePsaRegulationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建的预授权规则ID
	PsaId string `json:"PsaId"`
}
