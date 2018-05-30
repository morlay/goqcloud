package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 绑定弹性公网IP
// https://cloud.tencent.com/document/api/215/16700
type AssociateAddressRequest struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId string `name:"AddressId"`
	// 要绑定的实例 ID。实例 ID 形如：ins-11112222。可通过登录控制台查询，也可通过 DescribeInstances 接口返回值中的InstanceId获取。
	InstanceId *string `name:"InstanceId,omitempty"`
	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：eni-11112222。NetworkInterfaceId 与 InstanceId 不可同时指定。弹性网卡 ID 可通过登录控制台查询，也可通过DescribeNetworkInterfaces接口返回值中的networkInterfaceId获取。
	NetworkInterfaceId *string `name:"NetworkInterfaceId,omitempty"`
	// 要绑定的内网 IP。如果指定了 NetworkInterfaceId 则也必须指定 PrivateIpAddress ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 PrivateIpAddress 是指定的 NetworkInterfaceId 上的一个内网 IP。指定弹性网卡的内网 IP 可通过登录控制台查询，也可通过DescribeNetworkInterfaces接口返回值中的privateIpAddress获取。
	PrivateIpAddress *string `name:"PrivateIpAddress,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *AssociateAddressRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AssociateAddressResponse, error) {
	resp := &AssociateAddressResponse{}
	err := client.Request("vpc", "AssociateAddress", "2017-03-12").Do(req, resp)
	return resp, err
}

type AssociateAddressResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
