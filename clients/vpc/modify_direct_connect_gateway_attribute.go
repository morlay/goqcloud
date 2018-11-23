package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改专线网关属性
// https://cloud.tencent.com/document/api/215/30643

type ModifyDirectConnectGatewayAttributeRequest struct {
	// 云联网路由学习类型，可选值：BGP（自动学习）、STATIC（静态，即用户配置）。只有云联网类型专线网关且开启了BGP功能才支持修改CcnRouteType。
	CcnRouteType *string `name:"CcnRouteType,omitempty"`
	// 专线网关唯一ID，形如：dcg-9o233uri。
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 专线网关名称，可任意命名，但不得超过60个字符。
	DirectConnectGatewayName *string `name:"DirectConnectGatewayName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDirectConnectGatewayAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDirectConnectGatewayAttributeResponse, error) {
	resp := &ModifyDirectConnectGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyDirectConnectGatewayAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDirectConnectGatewayAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
