package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改对端网关
// https://cloud.tencent.com/document/api/215/17509
type ModifyCustomerGatewayAttributeRequest struct {
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId string `name:"CustomerGatewayId"`
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName string `name:"CustomerGatewayName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyCustomerGatewayAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyCustomerGatewayAttributeResponse, error) {
	resp := &ModifyCustomerGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyCustomerGatewayAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyCustomerGatewayAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
