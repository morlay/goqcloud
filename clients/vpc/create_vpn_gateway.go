package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建VPN网关
// https://cloud.tencent.com/document/api/215/17521

type CreateVpnGatewayRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `name:"InstanceChargePrepaid,omitempty"`
	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps
	InternetMaxBandwidthOut int64 `name:"InternetMaxBandwidthOut"`
	// 区域
	Region string `name:"Region"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName string `name:"VpnGatewayName"`
	// 可用区，如：ap-guangzhou-2。
	Zone *string `name:"Zone,omitempty"`
}

func (req *CreateVpnGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateVpnGatewayResponse, error) {
	resp := &CreateVpnGatewayResponse{}
	err := client.Request("vpc", "CreateVpnGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateVpnGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// VPN网关对象
	VpnGateway VpnGateway `json:"VpnGateway"`
}
