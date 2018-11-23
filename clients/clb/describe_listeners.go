package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询应用型负载均衡监听器列表
// https://cloud.tencent.com/document/api/214/30686

type DescribeListenersRequest struct {
	// 要查询的应用型负载均衡监听器 ID数组
	ListenerIds []*string `name:"ListenerIds,omitempty"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 要查询的监听器的端口
	Port *int64 `name:"Port,omitempty"`
	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL
	Protocol *string `name:"Protocol,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeListenersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeListenersResponse, error) {
	resp := &DescribeListenersResponse{}
	err := client.Request("clb", "DescribeListeners", "2018-03-17").Do(req, resp)
	return resp, err
}

type DescribeListenersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 监听器列表
	Listeners []*Listener `json:"Listeners"`
}
