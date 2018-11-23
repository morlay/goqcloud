package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取任务日志详情
// https://cloud.tencent.com/document/api/599/30318

type DescribeTaskLogsRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 最大任务实例数
	Limit *int64 `name:"Limit,omitempty"`
	// 起始任务实例
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 任务实例集合
	TaskInstanceIndexes []*int64 `name:"TaskInstanceIndexes,omitempty"`
	// 任务名称
	TaskName string `name:"TaskName"`
}

func (req *DescribeTaskLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskLogsResponse, error) {
	resp := &DescribeTaskLogsResponse{}
	err := client.Request("batch", "DescribeTaskLogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeTaskLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务实例日志详情集合
	TaskInstanceLogSet []*TaskInstanceLog `json:"TaskInstanceLogSet"`
	// 任务实例总数
	TotalCount int64 `json:"TotalCount"`
}
