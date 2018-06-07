package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询可支持的对端网关厂商信息
// https://cloud.tencent.com/document/api/215/17517

type DescribeCustomerGatewayVendorsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCustomerGatewayVendorsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCustomerGatewayVendorsResponse, error) {
	resp := &DescribeCustomerGatewayVendorsResponse{}
	err := client.Request("vpc", "DescribeCustomerGatewayVendors", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCustomerGatewayVendorsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 对端网关厂商信息对象。
	CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet"`
}
