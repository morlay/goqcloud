package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看安全组
// https://cloud.tencent.com/document/api/215/15808

type DescribeSecurityGroupsRequest struct {
	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。project-id - Integer - （过滤条件）项目id。security-group-name - String - （过滤条件）安全组名称。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `name:"SecurityGroupIds,omitempty"`
}

func (req *DescribeSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSecurityGroupsResponse, error) {
	resp := &DescribeSecurityGroupsResponse{}
	err := client.Request("vpc", "DescribeSecurityGroups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组对象。
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
