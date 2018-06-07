package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询IP地址模板集合
// https://cloud.tencent.com/document/api/215/16716

type DescribeAddressTemplateGroupsRequest struct {
	// 过滤条件。address-template-group-name - String - （过滤条件）IP地址模板集合名称。address-template-group-id - String - （过滤条件）IP地址模板实集合例ID，例如：ipmg-mdunqeb6。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAddressTemplateGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAddressTemplateGroupsResponse, error) {
	resp := &DescribeAddressTemplateGroupsResponse{}
	err := client.Request("vpc", "DescribeAddressTemplateGroups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAddressTemplateGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// IP地址模板。
	AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
