package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建负载均衡七层监听器转发规则
// https://cloud.tencent.com/document/api/214/30691

type CreateRuleRequest struct {
	// 监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 区域
	Region string `name:"Region"`
	// 新建转发规则的信息
	Rules []*RuleInput `name:"Rules"`
}

func (req *CreateRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateRuleResponse, error) {
	resp := &CreateRuleResponse{}
	err := client.Request("clb", "CreateRule", "2018-03-17").Do(req, resp)
	return resp, err
}

type CreateRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
