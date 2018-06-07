package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询任务详情
// https://cloud.tencent.com/document/api/599/15905

type DescribeTaskRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
	// 任务名称
	TaskName string `name:"TaskName"`
}

func (req *DescribeTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskResponse, error) {
	resp := &DescribeTaskResponse{}
	err := client.Request("batch", "DescribeTask", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 结束时间
	EndTime string `json:"EndTime"`
	// 作业ID
	JobId string `json:"JobId"`
	// 任务实例统计指标
	TaskInstanceMetrics TaskInstanceMetrics `json:"TaskInstanceMetrics"`
	// 任务实例信息
	TaskInstanceSet []*TaskInstanceView `json:"TaskInstanceSet"`
	// 任务实例总数
	TaskInstanceTotalCount int64 `json:"TaskInstanceTotalCount"`
	// 任务名称
	TaskName string `json:"TaskName"`
	// 任务状态
	TaskState string `json:"TaskState"`
}
