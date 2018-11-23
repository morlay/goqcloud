package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询Redis实例列表
// https://cloud.tencent.com/document/api/239/20018

type DescribeInstancesRequest struct {
	// 实例Id
	InstanceId *string `name:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `name:"InstanceName,omitempty"`
	// 实例列表大小
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，取Limit整数倍
	Offset *int64 `name:"Offset,omitempty"`
	// 枚举范围： projectId,createtime,instancename,type,curDeadline
	OrderBy *string `name:"OrderBy,omitempty"`
	// 1倒序，0顺序，默认倒序
	OrderType *int64 `name:"OrderType,omitempty"`
	// 项目ID 组成的数组，数组下标从0开始
	ProjectIds []*int64 `name:"ProjectIds,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询的Region的列表。
	RegionIds []*int64 `name:"RegionIds,omitempty"`
	// 查找实例的ID。
	SearchKey *string `name:"SearchKey,omitempty"`
	// 子网ID数组，数组下标从0开始
	SubnetIds []*string `name:"SubnetIds,omitempty"`
	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络
	VpcIds []*string `name:"VpcIds,omitempty"`
}

func (req *DescribeInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("redis", "DescribeInstances", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表
	InstanceSet []*InstanceSet `json:"InstanceSet"`
	// 实例数
	TotalCount int64 `json:"TotalCount"`
}
