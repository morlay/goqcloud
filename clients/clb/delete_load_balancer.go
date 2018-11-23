package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除负载均衡实例
// https://cloud.tencent.com/document/api/214/30689

type DeleteLoadBalancerRequest struct {
	// 要删除的负载均衡实例 ID数组
	LoadBalancerIds []*string `name:"LoadBalancerIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteLoadBalancerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteLoadBalancerResponse, error) {
	resp := &DeleteLoadBalancerResponse{}
	err := client.Request("clb", "DeleteLoadBalancer", "2018-03-17").Do(req, resp)
	return resp, err
}

type DeleteLoadBalancerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
