package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 终止任务实例
// https://cloud.tencent.com/document/api/599/15908
type TerminateTaskInstanceRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
	// 任务实例索引
	TaskInstanceIndex int64 `name:"TaskInstanceIndex"`
	// 任务名称
	TaskName string `name:"TaskName"`
}

func (req *TerminateTaskInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TerminateTaskInstanceResponse, error) {
	resp := &TerminateTaskInstanceResponse{}
	err := client.Request("batch", "TerminateTaskInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type TerminateTaskInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
