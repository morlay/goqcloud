package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询定时任务
// https://cloud.tencent.com/document/api/377/20450

type DescribeScheduledActionsRequest struct {
	// 过滤条件。 scheduled-action-id - String - 是否必填：否 -（过滤条件）按照定时任务ID过滤。 scheduled-action-name - String - 是否必填：否 - （过滤条件） 按照定时任务名称过滤。 auto-scaling-group-id - String - 是否必填：否 - （过滤条件） 按照伸缩组ID过滤。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 按照一个或者多个定时任务ID查询。实例ID形如：asst-am691zxo。每次请求的实例的上限为100。参数不支持同时指定ScheduledActionIds和Filters。
	ScheduledActionIds []*string `name:"ScheduledActionIds,omitempty"`
}

func (req *DescribeScheduledActionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeScheduledActionsResponse, error) {
	resp := &DescribeScheduledActionsResponse{}
	err := client.Request("as", "DescribeScheduledActions", "2018-04-19").Do(req, resp)
	return resp, err
}

type DescribeScheduledActionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 定时任务详细信息列表。
	ScheduledActionSet []*ScheduledAction `json:"ScheduledActionSet"`
	// 符合条件的定时任务数量。
	TotalCount int64 `json:"TotalCount"`
}
