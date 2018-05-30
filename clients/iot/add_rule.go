package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增规则
// https://cloud.tencent.com/document/api/568/16458
type AddRuleRequest struct {
	// 转发
	Actions []*Object `name:"Actions,omitempty"`
	// 描述
	Description string `name:"Description"`
	// 名称
	Name string `name:"Name"`
	// 查询
	Query *RuleQuery `name:"Query,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *AddRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddRuleResponse, error) {
	resp := &AddRuleResponse{}
	err := client.Request("iot", "AddRule", "2018-01-23").Do(req, resp)
	return resp, err
}

type AddRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 规则
	Rule Rule `json:"Rule"`
}
