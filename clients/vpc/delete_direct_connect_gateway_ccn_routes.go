package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除专线网关云联网路由
// https://cloud.tencent.com/document/api/215/19190

type DeleteDirectConnectGatewayCcnRoutesRequest struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 区域
	Region string `name:"Region"`
	// 路由ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `name:"RouteIds"`
}

func (req *DeleteDirectConnectGatewayCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDirectConnectGatewayCcnRoutesResponse, error) {
	resp := &DeleteDirectConnectGatewayCcnRoutesResponse{}
	err := client.Request("vpc", "DeleteDirectConnectGatewayCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteDirectConnectGatewayCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
