package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改七层转发规则的域名
// https://cloud.tencent.com/document/api/214/30682

type ModifyDomainRequest struct {
	// 监听器下的某个旧域名。
	Domain string `name:"Domain"`
	// 应用型负载均衡监听器 ID
	ListenerId string `name:"ListenerId"`
	// 负载均衡实例 ID
	LoadBalancerId string `name:"LoadBalancerId"`
	// 新域名， 长度限制为：1-80。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。
	NewDomain string `name:"NewDomain"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDomainRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDomainResponse, error) {
	resp := &ModifyDomainResponse{}
	err := client.Request("clb", "ModifyDomain", "2018-03-17").Do(req, resp)
	return resp, err
}

type ModifyDomainResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
