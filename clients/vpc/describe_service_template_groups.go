package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询协议端口模板集合
// https://cloud.tencent.com/document/api/215/16718
type DescribeServiceTemplateGroupsRequest struct {
	// 过滤条件。service-template-group-name - String - （过滤条件）协议端口模板集合名称。service-template-group-id - String - （过滤条件）协议端口模板集合实例ID，例如：ppmg-e6dy460g。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeServiceTemplateGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeServiceTemplateGroupsResponse, error) {
	resp := &DescribeServiceTemplateGroupsResponse{}
	err := client.Request("vpc", "DescribeServiceTemplateGroups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeServiceTemplateGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 协议端口模板集合。
	ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
