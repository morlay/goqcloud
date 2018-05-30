package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 下载VPN通道配置
// https://cloud.tencent.com/document/api/215/17513
type DownloadCustomerGatewayConfigurationRequest struct {
	// 对端网关厂商信息对象，可通过DescribeCustomerGatewayVendors获取。
	CustomerGatewayVendor CustomerGatewayVendor `name:"CustomerGatewayVendor"`
	// 通道接入设备物理接口名称。
	InterfaceName string `name:"InterfaceName"`
	// 区域
	Region string `name:"Region"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId string `name:"VpnConnectionId"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *DownloadCustomerGatewayConfigurationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DownloadCustomerGatewayConfigurationResponse, error) {
	resp := &DownloadCustomerGatewayConfigurationResponse{}
	err := client.Request("vpc", "DownloadCustomerGatewayConfiguration", "2017-03-12").Do(req, resp)
	return resp, err
}

type DownloadCustomerGatewayConfigurationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// XML格式配置信息。
	CustomerGatewayConfiguration string `json:"CustomerGatewayConfiguration"`
}
