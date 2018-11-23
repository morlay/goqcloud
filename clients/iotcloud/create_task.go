package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建任务
// https://cloud.tencent.com/document/api/634/19483

type CreateTaskRequest struct {
	// 执行任务的设备名的正则表达式
	DeviceNameFilter string `name:"DeviceNameFilter"`
	// 最长执行时间，单位秒，被调度后超过此时间仍未有结果则视为任务失败。取值为0-86400，默认为86400
	MaxExecutionTimeInSeconds *int64 `name:"MaxExecutionTimeInSeconds,omitempty"`
	// 执行任务的产品ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 任务开始执行的时间。 取值为 Unix 时间戳，单位秒，且需大于等于当前时间时间戳，0为系统当前时间时间戳，即立即执行，最大为当前时间86400秒后，超过则取值为当前时间86400秒后
	ScheduleTimeInSeconds int64 `name:"ScheduleTimeInSeconds"`
	// 任务类型，取值为 “UpdateShadow” 或者 “PublishMessage”
	TaskType string `name:"TaskType"`
	// 任务描述细节，描述见下 Task
	Tasks Task `name:"Tasks"`
}

func (req *CreateTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateTaskResponse, error) {
	resp := &CreateTaskResponse{}
	err := client.Request("iotcloud", "CreateTask", "2018-06-14").Do(req, resp)
	return resp, err
}

type CreateTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建的任务ID
	TaskId string `json:"TaskId"`
}
