package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取用户的所有资源信息
// https://cloud.tencent.com/document/api/283/18578

type DescribeResourceInstancesRequest struct {
	// 支持通过资源id，pid进行查询
	Filters []*Filter `name:"Filters,omitempty"`
	// 数量限制，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `name:"OrderDirection,omitempty"`
	// 按某个字段排序，目前支持CreateTime、ExpireTime其中的一个排序。
	OrderField *string `name:"OrderField,omitempty"`
	// 资源类别id数组，13624：加固专业版，12750：企业版。空数组表示返回全部资源。
	Pids []*int64 `name:"Pids"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeResourceInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeResourceInstancesResponse, error) {
	resp := &DescribeResourceInstancesResponse{}
	err := client.Request("ms", "DescribeResourceInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeResourceInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合要求的资源数组
	ResourceSet []*ResourceInfo `json:"ResourceSet"`
	// 符合要求的资源数量
	TotalCount int64 `json:"TotalCount"`
}
