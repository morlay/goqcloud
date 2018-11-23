package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询专线网关云联网路由
// https://cloud.tencent.com/document/api/215/19189

type DescribeDirectConnectGatewayCcnRoutesRequest struct {
	// 云联网路由学习类型，可选值：BGP - 自动学习。STATIC - 静态，即用户配置，默认值。
	CcnRouteType *string `name:"CcnRouteType,omitempty"`
	// 专线网关ID，形如：dcg-prpqlmg1。
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 返回数量。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDirectConnectGatewayCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDirectConnectGatewayCcnRoutesResponse, error) {
	resp := &DescribeDirectConnectGatewayCcnRoutesResponse{}
	err := client.Request("vpc", "DescribeDirectConnectGatewayCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDirectConnectGatewayCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云联网路由（IDC网段）列表。
	RouteSet []*DirectConnectGatewayCcnRoute `json:"RouteSet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
