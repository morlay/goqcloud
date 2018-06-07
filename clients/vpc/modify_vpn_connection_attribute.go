package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改VPN通道
// https://cloud.tencent.com/document/api/215/17508

type ModifyVpnConnectionAttributeRequest struct {
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `name:"IKEOptionsSpecification,omitempty"`
	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `name:"IPSECOptionsSpecification,omitempty"`
	// 预共享密钥。
	PreShareKey *string `name:"PreShareKey,omitempty"`
	// 区域
	Region string `name:"Region"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `name:"SecurityPolicyDatabases,omitempty"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId string `name:"VpnConnectionId"`
	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `name:"VpnConnectionName,omitempty"`
}

func (req *ModifyVpnConnectionAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyVpnConnectionAttributeResponse, error) {
	resp := &ModifyVpnConnectionAttributeResponse{}
	err := client.Request("vpc", "ModifyVpnConnectionAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyVpnConnectionAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
