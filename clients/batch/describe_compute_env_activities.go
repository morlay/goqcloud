package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看计算环境活动信息
// https://cloud.tencent.com/document/api/599/15896

type DescribeComputeEnvActivitiesRequest struct {
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 过滤条件 compute-node-id - String - 是否必填：否 -（过滤条件）按照计算节点ID过滤。
	Filters *Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComputeEnvActivitiesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComputeEnvActivitiesResponse, error) {
	resp := &DescribeComputeEnvActivitiesResponse{}
	err := client.Request("batch", "DescribeComputeEnvActivities", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeComputeEnvActivitiesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 计算环境中的活动列表
	ActivitySet []*Activity `json:"ActivitySet"`
	// 活动数量
	TotalCount int64 `json:"TotalCount"`
}
