package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改预授权规则
// https://cloud.tencent.com/document/api/386/30140

type ModifyPsaRegulationRequest struct {
	// 预授权规则备注
	PsaDescription *string `name:"PsaDescription,omitempty"`
	// 预授权规则ID
	PsaId string `name:"PsaId"`
	// 预授权规则别名
	PsaName *string `name:"PsaName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 维修中的实例上限
	RepairLimit *int64 `name:"RepairLimit,omitempty"`
	// 预授权规则关联故障类型ID列表
	TaskTypeIds []*int64 `name:"TaskTypeIds,omitempty"`
}

func (req *ModifyPsaRegulationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPsaRegulationResponse, error) {
	resp := &ModifyPsaRegulationResponse{}
	err := client.Request("bm", "ModifyPsaRegulation", "2018-04-23").Do(req, resp)
	return resp, err
}

type ModifyPsaRegulationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
