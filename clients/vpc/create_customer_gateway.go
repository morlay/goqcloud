package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建对端网关
// https://cloud.tencent.com/document/api/215/17523

type CreateCustomerGatewayRequest struct {
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName string `name:"CustomerGatewayName"`
	// 对端网关公网IP。
	IpAddress string `name:"IpAddress"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateCustomerGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateCustomerGatewayResponse, error) {
	resp := &CreateCustomerGatewayResponse{}
	err := client.Request("vpc", "CreateCustomerGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateCustomerGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 对端网关对象
	CustomerGateway CustomerGateway `json:"CustomerGateway"`
}
