package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取转发规则信息
// https://cloud.tencent.com/document/api/568/16461

type GetRuleRequest struct {
	// 区域
	Region string `name:"Region"`
	// 规则Id
	RuleId string `name:"RuleId"`
}

func (req *GetRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetRuleResponse, error) {
	resp := &GetRuleResponse{}
	err := client.Request("iot", "GetRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 规则
	Rule Rule `json:"Rule"`
}
