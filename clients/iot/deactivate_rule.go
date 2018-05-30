package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 禁用规则
// https://cloud.tencent.com/document/api/568/16459
type DeactivateRuleRequest struct {
	// 区域
	Region string `name:"Region"`
	// 规则Id
	RuleId string `name:"RuleId"`
}

func (req *DeactivateRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeactivateRuleResponse, error) {
	resp := &DeactivateRuleResponse{}
	err := client.Request("iot", "DeactivateRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type DeactivateRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
