package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 绑定后端机器到监听器上
// https://cloud.tencent.com/document/api/214/30676

type RegisterTargetsRequest struct {
	// 目标规则的域名，提供LocationId参数时本参数不生效
	Domain *string `name:"Domain,omitempty"`
	// 负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 转发规则的ID，当注册机器到七层转发规则时，必须提供此参数或Domain+Url两者之一
	LocationId *string `name:"LocationId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要注册的后端机器列表
	Targets []*Target `name:"Targets"`
	// 目标规则的URL，提供LocationId参数时本参数不生效
	Url *string `name:"Url,omitempty"`
}

func (req *RegisterTargetsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RegisterTargetsResponse, error) {
	resp := &RegisterTargetsResponse{}
	err := client.Request("clb", "RegisterTargets", "2018-03-17").Do(req, resp)
	return resp, err
}

type RegisterTargetsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
