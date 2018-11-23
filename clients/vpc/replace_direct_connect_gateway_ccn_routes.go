package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 替换专线网关云联网路由
// https://cloud.tencent.com/document/api/215/19188

type ReplaceDirectConnectGatewayCcnRoutesRequest struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 区域
	Region string `name:"Region"`
	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `name:"Routes"`
}

func (req *ReplaceDirectConnectGatewayCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ReplaceDirectConnectGatewayCcnRoutesResponse, error) {
	resp := &ReplaceDirectConnectGatewayCcnRoutesResponse{}
	err := client.Request("vpc", "ReplaceDirectConnectGatewayCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type ReplaceDirectConnectGatewayCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
