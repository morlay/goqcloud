package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建专线网关云联网路由
// https://cloud.tencent.com/document/api/215/19191

type CreateDirectConnectGatewayCcnRoutesRequest struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId string `name:"DirectConnectGatewayId"`
	// 区域
	Region string `name:"Region"`
	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `name:"Routes"`
}

func (req *CreateDirectConnectGatewayCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDirectConnectGatewayCcnRoutesResponse, error) {
	resp := &CreateDirectConnectGatewayCcnRoutesResponse{}
	err := client.Request("vpc", "CreateDirectConnectGatewayCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDirectConnectGatewayCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
