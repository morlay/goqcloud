package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改监听器绑定的后端机器的端口
// https://cloud.tencent.com/document/api/214/30678

type ModifyTargetPortRequest struct {
	// 目标规则的域名，提供LocationId参数时本参数不生效
	Domain *string `name:"Domain,omitempty"`
	// 负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 转发规则的ID
	LocationId *string `name:"LocationId,omitempty"`
	// 后端机器绑定到监听器的新端口
	NewPort int64 `name:"NewPort"`
	// 区域
	Region string `name:"Region"`
	// 要修改端口的后端机器列表
	Targets []*Target `name:"Targets"`
	// 目标规则的URL，提供LocationId参数时本参数不生效
	Url *string `name:"Url,omitempty"`
}

func (req *ModifyTargetPortRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyTargetPortResponse, error) {
	resp := &ModifyTargetPortResponse{}
	err := client.Request("clb", "ModifyTargetPort", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyTargetPortResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
