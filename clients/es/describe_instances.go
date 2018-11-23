package es

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询ES集群实例
// https://cloud.tencent.com/document/api/845/30631

type DescribeInstancesRequest struct {
	// 一个或多个集群实例ID
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 一个或多个集群实例名称
	InstanceNames []*string `name:"InstanceNames,omitempty"`
	// 分页大小，默认值20
	Limit *int64 `name:"Limit,omitempty"`
	// 分页起始值, 默认值0
	Offset *int64 `name:"Offset,omitempty"`
	// 排序字段：1，实例ID；2，实例名称；3，可用区；4，创建时间，若orderKey未传递则按创建时间降序排序
	OrderByKey *int64 `name:"OrderByKey,omitempty"`
	// 排序方式：0，升序；1，降序；若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *int64 `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 集群实例所属可用区，不传则默认所有可用区
	Zone *string `name:"Zone,omitempty"`
}

func (req *DescribeInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("es", "DescribeInstances", "2018-04-16").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表
	InstanceList []*InstanceInfo `json:"InstanceList"`
	// 返回的实例个数
	TotalCount int64 `json:"TotalCount"`
}
