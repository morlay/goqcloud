package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看实例状态列表
// https://cloud.tencent.com/document/api/213/15738

type DescribeInstancesStatusRequest struct {
	// 按照一个或者多个实例ID查询。实例ID形如：ins-11112222。此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstancesStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstancesStatusResponse, error) {
	resp := &DescribeInstancesStatusResponse{}
	err := client.Request("cvm", "DescribeInstancesStatus", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstancesStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例状态 列表。
	InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet"`
	// 符合条件的实例状态数量。
	TotalCount int64 `json:"TotalCount"`
}
