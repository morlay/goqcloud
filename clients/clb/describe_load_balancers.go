package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询负载均衡实例列表
// https://cloud.tencent.com/document/api/214/30685

type DescribeLoadBalancersRequest struct {
	// 后端云服务器的内网 IP。
	BackendPrivateIps []*string `name:"BackendPrivateIps,omitempty"`
	// 后端云服务器的外网 IP。
	BackendPublicIps []*string `name:"BackendPublicIps,omitempty"`
	// 腾讯云为负载均衡实例分配的域名，应用型负载均衡该字段无意义。
	Domain *string `name:"Domain,omitempty"`
	// 1：应用型，0：传统型，-1：全部类型。
	Forward *int64 `name:"Forward,omitempty"`
	// 返回负载均衡个数，默认为 20。
	Limit *int64 `name:"Limit,omitempty"`
	// 负载均衡实例 ID。
	LoadBalancerIds []*string `name:"LoadBalancerIds,omitempty"`
	// 负载均衡实例名称。
	LoadBalancerName *string `name:"LoadBalancerName,omitempty"`
	// 负载均衡实例的网络类型：OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `name:"LoadBalancerType,omitempty"`
	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `name:"LoadBalancerVips,omitempty"`
	// 数据偏移量，默认为 0。
	Offset *int64 `name:"Offset,omitempty"`
	// 排序字段，支持以下字段：LoadBalancerName，CreateTime，Domain，LoadBalancerType。
	OrderBy *string `name:"OrderBy,omitempty"`
	// 1：倒序，0：顺序，默认按照创建时间倒序。
	OrderType *int64 `name:"OrderType,omitempty"`
	// 负载均衡实例所属的项目 ID，可以通过 DescribeProject 接口获取。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 搜索字段，模糊匹配名称、域名、VIP。
	SearchKey *string `name:"SearchKey,omitempty"`
	// 查询的负载均衡是否绑定后端服务器，0：没有绑定云服务器，1：绑定云服务器，-1：查询全部。
	WithRs *int64 `name:"WithRs,omitempty"`
}

func (req *DescribeLoadBalancersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLoadBalancersResponse, error) {
	resp := &DescribeLoadBalancersResponse{}
	err := client.Request("clb", "DescribeLoadBalancers", "2018-03-17").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的负载均衡实例数组。
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet"`
	// 满足过滤条件的负载均衡实例总数。
	TotalCount int64 `json:"TotalCount"`
}
