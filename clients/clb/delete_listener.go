package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除负载均衡监听器
// https://cloud.tencent.com/document/api/214/30690

type DeleteListenerRequest struct {
	// 要删除的监听器 ID
	ListenerId string `name:"ListenerId"`
	// 应用型负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteListenerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteListenerResponse, error) {
	resp := &DeleteListenerResponse{}
	err := client.Request("clb", "DeleteListener", "2018-03-17").Do(req, resp)
	return resp, err
}

type DeleteListenerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
