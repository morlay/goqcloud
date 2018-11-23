package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询应用型负载均衡云服务器列表
// https://cloud.tencent.com/document/api/214/30684

type DescribeTargetsRequest struct {
	// 监听器 ID列表
	ListenerIds []*string `name:"ListenerIds,omitempty"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 负载均衡监听器端口
	Port *int64 `name:"Port,omitempty"`
	// 监听器协议类型
	Protocol *string `name:"Protocol,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeTargetsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTargetsResponse, error) {
	resp := &DescribeTargetsResponse{}
	err := client.Request("clb", "DescribeTargets", "2018-03-17").Do(req, resp)
	return resp, err
}

type DescribeTargetsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 监听器后端绑定的机器信息
	Listeners []*ListenerBackend `json:"Listeners"`
}
