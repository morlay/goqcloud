package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建专线网关
// https://cloud.tencent.com/document/api/215/19192

type CreateDirectConnectGatewayRequest struct {
	// 专线网关名称
	DirectConnectGatewayName string `name:"DirectConnectGatewayName"`
	// 网关类型，可选值：NORMAL - （默认）标准型，注：云联网只支持标准型NAT - NAT型NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关
	GatewayType *string `name:"GatewayType,omitempty"`
	// NetworkType 为 VPC 时，这里传值为私有网络实例IDNetworkType 为 NAT 时，这里传值为云联网实例ID
	NetworkInstanceId string `name:"NetworkInstanceId"`
	// 关联网络类型，可选值：VPC - 私有网络CCN - 云联网
	NetworkType string `name:"NetworkType"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateDirectConnectGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDirectConnectGatewayResponse, error) {
	resp := &CreateDirectConnectGatewayResponse{}
	err := client.Request("vpc", "CreateDirectConnectGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDirectConnectGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 专线网关对象。
	DirectConnectGateway DirectConnectGateway `json:"DirectConnectGateway"`
}
