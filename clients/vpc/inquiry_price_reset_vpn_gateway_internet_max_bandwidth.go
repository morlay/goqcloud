package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 调整VPN网关带宽上限询价
// https://cloud.tencent.com/document/api/215/17510
type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut int64 `name:"InternetMaxBandwidthOut"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse, error) {
	resp := &InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse{}
	err := client.Request("vpc", "InquiryPriceResetVpnGatewayInternetMaxBandwidth", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 商品价格。
	Price Price `json:"Price"`
}
