package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建VPN网关询价
// https://cloud.tencent.com/document/api/215/17512

type InquiryPriceCreateVpnGatewayRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `name:"InstanceChargePrepaid,omitempty"`
	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut int64 `name:"InternetMaxBandwidthOut"`
	// 区域
	Region string `name:"Region"`
}

func (req *InquiryPriceCreateVpnGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceCreateVpnGatewayResponse, error) {
	resp := &InquiryPriceCreateVpnGatewayResponse{}
	err := client.Request("vpc", "InquiryPriceCreateVpnGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceCreateVpnGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 商品价格。
	Price Price `json:"Price"`
}
