package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询弹性网卡列表
// https://cloud.tencent.com/document/api/215/15817

type DescribeNetworkInterfacesRequest struct {
	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。network-interface-name - String - （过滤条件）网卡实例名称。network-interface-description - String - （过滤条件）网卡实例描述。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `name:"NetworkInterfaceIds,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeNetworkInterfacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeNetworkInterfacesResponse, error) {
	resp := &DescribeNetworkInterfacesResponse{}
	err := client.Request("vpc", "DescribeNetworkInterfaces", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeNetworkInterfacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表。
	NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
