package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看CDH实例列表
// https://cloud.tencent.com/document/api/213/16474

type DescribeHostsRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 - （过滤条件）按照可用区过滤。 project-id - Integer - 是否必填：否 - （过滤条件）按照项目ID过滤。可通过调用 DescribeProject 查询已创建的项目列表或登录控制台进行查看；也可以调用 AddProject 创建新的项目。 host-id - String - 是否必填：否 - （过滤条件）按照CDH ID过滤。CDH ID形如：host-11112222。 host-name - String - 是否必填：否 - （过滤条件）按照CDH实例名称过滤。 host-state - String - 是否必填：否 - （过滤条件）按照CDH实例状态进行过滤。（PENDING：创建中|LAUNCH_FAILURE：创建失败|RUNNING：运行中|EXPIRED：已过期）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeHostsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeHostsResponse, error) {
	resp := &DescribeHostsResponse{}
	err := client.Request("cvm", "DescribeHosts", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeHostsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// cdh实例详细信息列表
	HostSet []*HostItem `json:"HostSet"`
	// 符合查询条件的cdh实例总数
	TotalCount int64 `json:"TotalCount"`
}
