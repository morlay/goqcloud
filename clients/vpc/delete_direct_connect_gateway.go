package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除专线网关
// https://cloud.tencent.com/document/api/215/30645

type DeleteDirectConnectGatewayRequest struct {
	// 专线网关唯一ID，形如：dcg-9o233uri。
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDirectConnectGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDirectConnectGatewayResponse, error) {
	resp := &DeleteDirectConnectGatewayResponse{}
	err := client.Request("vpc", "DeleteDirectConnectGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteDirectConnectGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
