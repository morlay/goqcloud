package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改应用型负载均衡监听器转发规则
// https://cloud.tencent.com/document/api/214/30679

type ModifyRuleRequest struct {
	// 健康检查信息
	HealthCheck *HealthCheck `name:"HealthCheck,omitempty"`
	// 应用型负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 要修改的转发规则的 ID。
	LocationId string `name:"LocationId"`
	// 区域
	Region string `name:"Region"`
	// 规则的请求转发方式
	Scheduler *string `name:"Scheduler,omitempty"`
	// 会话保持时间
	SessionExpireTime *int64 `name:"SessionExpireTime,omitempty"`
	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数
	Url *string `name:"Url,omitempty"`
}

func (req *ModifyRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyRuleResponse, error) {
	resp := &ModifyRuleResponse{}
	err := client.Request("clb", "ModifyRule", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
