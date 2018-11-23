package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// Redis查询任务结果
// https://cloud.tencent.com/document/api/239/30601

type DescribeTaskInfoRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务ID
	TaskId int64 `name:"TaskId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskInfoResponse, error) {
	resp := &DescribeTaskInfoResponse{}
	err := client.Request("redis", "DescribeTaskInfo", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeTaskInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例的ID
	InstanceId string `json:"InstanceId"`
	// 任务开始时间
	StartTime string `json:"StartTime"`
	// 任务状态preparing:待执行，running：执行中，succeed：成功，failed：失败，error 执行出错
	Status string `json:"Status"`
	// 任务信息，错误时显示错误信息。执行中与成功则为空
	TaskMessage string `json:"TaskMessage"`
	// 任务类型
	TaskType string `json:"TaskType"`
}
