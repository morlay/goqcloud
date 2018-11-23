package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改负载均衡监听器属性
// https://cloud.tencent.com/document/api/214/30681

type ModifyListenerRequest struct {
	// 证书相关信息，此参数仅适用于HTTPS/TCP_SSL监听器
	Certificate *CertificateInput `name:"Certificate,omitempty"`
	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器
	HealthCheck *HealthCheck `name:"HealthCheck,omitempty"`
	// 负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 新的监听器名称
	ListenerName *string `name:"ListenerName,omitempty"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 区域
	Region string `name:"Region"`
	// 监听器转发的方式。可选值：WRR、LEAST_CONN分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `name:"Scheduler,omitempty"`
	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `name:"SessionExpireTime,omitempty"`
}

func (req *ModifyListenerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyListenerResponse, error) {
	resp := &ModifyListenerResponse{}
	err := client.Request("clb", "ModifyListener", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyListenerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
