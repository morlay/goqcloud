package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建弹性网卡
// https://cloud.tencent.com/document/api/215/15818

type CreateNetworkInterfaceRequest struct {
	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `name:"NetworkInterfaceDescription,omitempty"`
	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName string `name:"NetworkInterfaceName"`
	// 指定内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `name:"PrivateIpAddresses,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 新申请的内网IP地址个数。
	SecondaryPrivateIpAddressCount *int64 `name:"SecondaryPrivateIpAddressCount,omitempty"`
	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `name:"SecurityGroupIds,omitempty"`
	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId string `name:"SubnetId"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
}

func (req *CreateNetworkInterfaceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateNetworkInterfaceResponse, error) {
	resp := &CreateNetworkInterfaceResponse{}
	err := client.Request("vpc", "CreateNetworkInterface", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateNetworkInterfaceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 弹性网卡实例。
	NetworkInterface NetworkInterface `json:"NetworkInterface"`
}
