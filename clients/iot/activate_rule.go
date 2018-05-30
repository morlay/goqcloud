package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 启用规则
// https://cloud.tencent.com/document/api/568/16457
type ActivateRuleRequest struct {
	// 区域
	Region string `name:"Region"`
	// 规则Id
	RuleId string `name:"RuleId"`
}

func (req *ActivateRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ActivateRuleResponse, error) {
	resp := &ActivateRuleResponse{}
	err := client.Request("iot", "ActivateRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type ActivateRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
