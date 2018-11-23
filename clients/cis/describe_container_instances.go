package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询容器实例列表
// https://cloud.tencent.com/document/api/858/17771

type DescribeContainerInstancesRequest struct {
	// 过滤条件。- Zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。- VpcId - String - 是否必填：否 -（过滤条件）按照VpcId过滤。- InstanceName - String - 是否必填：否 -（过滤条件）按照容器实例名称做模糊查询。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeContainerInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeContainerInstancesResponse, error) {
	resp := &DescribeContainerInstancesResponse{}
	err := client.Request("cis", "DescribeContainerInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeContainerInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 容器实例列表
	ContainerInstanceList []*ContainerInstance `json:"ContainerInstanceList"`
	// 容器实例总数
	TotalCount int64 `json:"TotalCount"`
}
