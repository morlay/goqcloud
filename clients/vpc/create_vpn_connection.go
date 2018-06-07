package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建VPN通道
// https://cloud.tencent.com/document/api/215/17522

type CreateVpnConnectionRequest struct {
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId string `name:"CustomerGatewayId"`
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自保护机制，用户配置网络安全协议
	IKEOptionsSpecification *IKEOptionsSpecification `name:"IKEOptionsSpecification,omitempty"`
	// IPSec配置，腾讯云提供IPSec安全会话设置
	IPSECOptionsSpecification *IPSECOptionsSpecification `name:"IPSECOptionsSpecification,omitempty"`
	// 预共享密钥。
	PreShareKey string `name:"PreShareKey"`
	// 区域
	Region string `name:"Region"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `name:"SecurityPolicyDatabases"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
	// 通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName string `name:"VpnConnectionName"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *CreateVpnConnectionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateVpnConnectionResponse, error) {
	resp := &CreateVpnConnectionResponse{}
	err := client.Request("vpc", "CreateVpnConnection", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateVpnConnectionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 通道实例对象。
	VpnConnection VpnConnection `json:"VpnConnection"`
}
