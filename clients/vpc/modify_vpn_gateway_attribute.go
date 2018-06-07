package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改VPN网关属性
// https://cloud.tencent.com/document/api/215/17507

type ModifyVpnGatewayAttributeRequest struct {
	// VPN网关计费模式，目前只支持预付费（即包年包月）到后付费（即按量计费）的转换。即参数只支持：POSTPAID_BY_HOUR。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `name:"VpnGatewayName,omitempty"`
}

func (req *ModifyVpnGatewayAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyVpnGatewayAttributeResponse, error) {
	resp := &ModifyVpnGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyVpnGatewayAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyVpnGatewayAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
