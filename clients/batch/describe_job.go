package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看作业详情
// https://cloud.tencent.com/document/api/599/15904
type DescribeJobRequest struct {
	// 作业标识
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeJobResponse, error) {
	resp := &DescribeJobResponse{}
	err := client.Request("batch", "DescribeJob", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 任务间依赖信息
	DependenceSet []*Dependence `json:"DependenceSet"`
	// 结束时间
	EndTime string `json:"EndTime"`
	// 作业ID
	JobId string `json:"JobId"`
	// 作业名称
	JobName string `json:"JobName"`
	// 作业状态
	JobState string `json:"JobState"`
	// 作业优先级
	Priority int64 `json:"Priority"`
	// 作业失败原因
	StateReason string `json:"StateReason"`
	// 任务实例统计指标
	TaskInstanceMetrics TaskInstanceView `json:"TaskInstanceMetrics"`
	// 任务统计指标
	TaskMetrics TaskMetrics `json:"TaskMetrics"`
	// 任务视图信息
	TaskSet []*TaskView `json:"TaskSet"`
	// 可用区信息
	Zone string `json:"Zone"`
}
