package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除规则
// https://cloud.tencent.com/document/api/568/16460

type DeleteRuleRequest struct {
	// 区域
	Region string `name:"Region"`
	// 规则Id
	RuleId string `name:"RuleId"`
}

func (req *DeleteRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteRuleResponse, error) {
	resp := &DeleteRuleResponse{}
	err := client.Request("iot", "DeleteRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type DeleteRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
