package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除应用型七层负载均衡监听器规则
// https://cloud.tencent.com/document/api/214/30688

type DeleteRuleRequest struct {
	// 要删除的转发规则的域名，已提供LocationIds参数时本参数不生效
	Domain *string `name:"Domain,omitempty"`
	// 应用型负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 要删除的转发规则的ID组成的数组
	LocationIds []*string `name:"LocationIds,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要删除的转发规则的转发路径，已提供LocationIds参数时本参数不生效
	Url *string `name:"Url,omitempty"`
}

func (req *DeleteRuleRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteRuleResponse, error) {
	resp := &DeleteRuleResponse{}
	err := client.Request("clb", "DeleteRule", "2018-03-17").Do(req, resp)
	return resp, err
}

type DeleteRuleResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
