package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看计算环境活动信息
// https://cloud.tencent.com/document/api/599/15896
type DescribeComputeEnvActivitiesRequest struct {
	// u8ba1u7b97u73afu5883ID
	EnvId string `name:"EnvId"`
	// 过滤条件
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
