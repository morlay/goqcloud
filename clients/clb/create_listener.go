package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建负载均衡监听器
// https://cloud.tencent.com/document/api/214/30693

type CreateListenerRequest struct {
	// 证书相关信息，此参数仅适用于HTTPS/TCP_SSL监听器
	Certificate *CertificateInput `name:"Certificate,omitempty"`
	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器
	HealthCheck *HealthCheck `name:"HealthCheck,omitempty"`
	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数
	ListenerNames []*string `name:"ListenerNames,omitempty"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器
	Ports []*int64 `name:"Ports"`
	// 监听器协议：HTTP | HTTPS | TCP | TCP_SSL
	Protocol string `name:"Protocol"`
	// 区域
	Region string `name:"Region"`
	// 监听器转发的方式。可选值：WRR、LEAST_CONN分别表示按权重轮询、最小连接数， 默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL监听器。
	Scheduler *string `name:"Scheduler,omitempty"`
	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `name:"SessionExpireTime,omitempty"`
	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。
	SniSwitch *int64 `name:"SniSwitch,omitempty"`
}

func (req *CreateListenerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateListenerResponse, error) {
	resp := &CreateListenerResponse{}
	err := client.Request("clb", "CreateListener", "2018-03-17").Do(req, resp)
	return resp, err
}

type CreateListenerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建的监听器的唯一标识数组
	ListenerIds []*string `json:"ListenerIds"`
}
