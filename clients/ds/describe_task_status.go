package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取任务状态
// https://cloud.tencent.com/document/api/869/17789

type DescribeTaskStatusRequest struct {
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
	// 任务ID
	TaskId int64 `name:"TaskId"`
}

func (req *DescribeTaskStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskStatusResponse, error) {
	resp := &DescribeTaskStatusResponse{}
	err := client.Request("ds", "DescribeTaskStatus", "2018-05-23").Do(req, resp)
	return resp, err
}

type DescribeTaskStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务结果
	TaskResult string `json:"TaskResult"`
	// 任务类型，010代表合同上传结果，020代表合同下载结果
	TaskType string `json:"TaskType"`
}
