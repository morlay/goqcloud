package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改监听器绑定的后端机器的转发权重
// https://cloud.tencent.com/document/api/214/30677

type ModifyTargetWeightRequest struct {
	// 目标规则的域名，提供LocationId参数时本参数不生效
	Domain *string `name:"Domain,omitempty"`
	// 负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一
	LocationId *string `name:"LocationId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要修改权重的后端机器列表
	Targets []*Target `name:"Targets,omitempty"`
	// 目标规则的URL，提供LocationId参数时本参数不生效
	Url *string `name:"Url,omitempty"`
	// 后端云服务器新的转发权重，取值范围：0~100。
	Weight int64 `name:"Weight"`
}

func (req *ModifyTargetWeightRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyTargetWeightResponse, error) {
	resp := &ModifyTargetWeightResponse{}
	err := client.Request("clb", "ModifyTargetWeight", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyTargetWeightResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
