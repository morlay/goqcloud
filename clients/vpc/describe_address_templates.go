package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询IP地址模板
// https://cloud.tencent.com/document/api/215/16717

type DescribeAddressTemplatesRequest struct {
	// 过滤条件。address-template-name - String - （过滤条件）IP地址模板名称。address-template-id - String - （过滤条件）IP地址模板实例ID，例如：ipm-mdunqeb6。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAddressTemplatesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAddressTemplatesResponse, error) {
	resp := &DescribeAddressTemplatesResponse{}
	err := client.Request("vpc", "DescribeAddressTemplates", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAddressTemplatesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// IP地址模版。
	AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
