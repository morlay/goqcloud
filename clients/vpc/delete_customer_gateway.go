package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除对端网关
// https://cloud.tencent.com/document/api/215/17520

type DeleteCustomerGatewayRequest struct {
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId string `name:"CustomerGatewayId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteCustomerGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteCustomerGatewayResponse, error) {
	resp := &DeleteCustomerGatewayResponse{}
	err := client.Request("vpc", "DeleteCustomerGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteCustomerGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
