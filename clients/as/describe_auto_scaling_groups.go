package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询伸缩组
// https://cloud.tencent.com/document/api/377/20438

type DescribeAutoScalingGroupsRequest struct {
	// 按照一个或者多个伸缩组ID查询。伸缩组ID形如：asg-nkdwoui0。每次请求的上限为100。参数不支持同时指定AutoScalingGroups和Filters。
	AutoScalingGroupIds []*string `name:"AutoScalingGroupIds,omitempty"`
	// 过滤条件。 auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。 auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称过滤。 launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定AutoScalingGroupIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAutoScalingGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAutoScalingGroupsResponse, error) {
	resp := &DescribeAutoScalingGroupsResponse{}
	err := client.Request("as", "DescribeAutoScalingGroups", "2018-04-19").Do(req, resp)
	return resp, err
}

type DescribeAutoScalingGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 伸缩组详细信息列表。
	AutoScalingGroupSet []*AutoScalingGroup `json:"AutoScalingGroupSet"`
	// 符合条件的伸缩组数量。
	TotalCount int64 `json:"TotalCount"`
}
