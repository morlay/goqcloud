package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费VPN网关询价
// https://cloud.tencent.com/document/api/215/17511
type InquiryPriceRenewVpnGatewayRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid InstanceChargePrepaid `name:"InstanceChargePrepaid"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *InquiryPriceRenewVpnGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceRenewVpnGatewayResponse, error) {
	resp := &InquiryPriceRenewVpnGatewayResponse{}
	err := client.Request("vpc", "InquiryPriceRenewVpnGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceRenewVpnGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 商品价格。
	Price Price `json:"Price"`
}
