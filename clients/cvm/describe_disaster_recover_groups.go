package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询分散置放群组信息
// https://cloud.tencent.com/document/api/213/17810

type DescribeDisasterRecoverGroupsRequest struct {
	// 分散置放群组ID列表。
	DisasterRecoverGroupIds []*string `name:"DisasterRecoverGroupIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 分散置放群组名称，支持模糊匹配。
	Name *string `name:"Name,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDisasterRecoverGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDisasterRecoverGroupsResponse, error) {
	resp := &DescribeDisasterRecoverGroupsResponse{}
	err := client.Request("cvm", "DescribeDisasterRecoverGroups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDisasterRecoverGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 分散置放群组信息列表。
	DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet"`
	// 用户置放群组总量。
	TotalCount int64 `json:"TotalCount"`
}
