package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 购买负载均衡实例
// https://cloud.tencent.com/document/api/214/30692

type CreateLoadBalancerRequest struct {
	// 负载均衡实例。1：应用型，0：传统型，默认为应用型负载均衡实例。
	Forward *int64 `name:"Forward,omitempty"`
	// 负载均衡实例的名称，只用来创建一个的时候生效。规则：1-50 个英文、汉字、数字、连接线“-”或下划线“_”。注意：如果名称与系统中已有负载均衡实例的名称重复的话，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `name:"LoadBalancerName,omitempty"`
	// 负载均衡实例的网络类型：OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType string `name:"LoadBalancerType"`
	// 负载均衡实例所属的项目 ID，可以通过 DescribeProject 接口获取。不填则属于默认项目。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 在私有网络内购买内网负载均衡实例的时候需要指定子网 ID，内网负载均衡实例的 VIP 将从这个子网中产生。其他情况不用填写该字段。
	SubnetId *string `name:"SubnetId,omitempty"`
	// 负载均衡后端实例所属网络 ID，可以通过 DescribeVpcEx 接口获取。 不填则默认为基础网络。
	VpcId *string `name:"VpcId,omitempty"`
}

func (req *CreateLoadBalancerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateLoadBalancerResponse, error) {
	resp := &CreateLoadBalancerResponse{}
	err := client.Request("clb", "CreateLoadBalancer", "2018-03-17").Do(req, resp)
	return resp, err
}

type CreateLoadBalancerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 由负载均衡实例统一 ID 组成的数组。
	LoadBalancerIds []*string `json:"LoadBalancerIds"`
}
