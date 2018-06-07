package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新规则
// https://cloud.tencent.com/document/api/568/16463

type UpdateRuleRequest struct {
	// 转发
	Actions []*Object `name:"Actions,omitempty"`
	// 描述
	Description *string `name:"Description,omitempty"`
	// 名称
	Name *string `name:"Name,omitempty"`
	// 查询
	Query *RuleQuery `name:"Query,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 规则Id
	RuleId string `name:"RuleId"`
}

func (req *UpdateRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateRuleResponse, error) {
	resp := &UpdateRuleResponse{}
	err := client.Request("iot", "UpdateRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type UpdateRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
