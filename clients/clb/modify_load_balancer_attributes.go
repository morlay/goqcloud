package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改负载均衡实例的属性
// https://cloud.tencent.com/document/api/214/30680

type ModifyLoadBalancerAttributesRequest struct {
	// 负载均衡的唯一ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 负载均衡实例名称
	LoadBalancerName *string `name:"LoadBalancerName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyLoadBalancerAttributesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyLoadBalancerAttributesResponse, error) {
	resp := &ModifyLoadBalancerAttributesResponse{}
	err := client.Request("clb", "ModifyLoadBalancerAttributes", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyLoadBalancerAttributesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
