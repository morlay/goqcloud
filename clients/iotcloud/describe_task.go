package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取任务详情
// https://cloud.tencent.com/document/api/634/19482

type DescribeTaskRequest struct {
	// 任务ID
	Id string `name:"Id"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskResponse, error) {
	resp := &DescribeTaskResponse{}
	err := client.Request("iotcloud", "DescribeTask", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 匹配到的需执行任务的设备数目
	AllDeviceCnt int64 `json:"AllDeviceCnt"`
	// 任务创建时间，Unix 时间戳
	CreateTime int64 `json:"CreateTime"`
	// 已完成任务的设备数目
	DoneDeviceCnt int64 `json:"DoneDeviceCnt"`
	// 任务完成时间，Unix 时间戳
	DoneTime int64 `json:"DoneTime"`
	// 返回的错误信息
	ErrMsg string `json:"ErrMsg"`
	// 任务 ID
	Id string `json:"Id"`
	// 完成任务的设备比例
	Percent int64 `json:"Percent"`
	// 产品 ID
	ProductId string `json:"ProductId"`
	// 返回的错误码
	RetCode int64 `json:"RetCode"`
	// 被调度时间，Unix 时间戳
	ScheduleTime int64 `json:"ScheduleTime"`
	// 状态。1表示等待处理，2表示调度处理中，3表示已完成，4表示失败，5表示已取消
	Status int64 `json:"Status"`
	// 任务类型，目前取值为 “UpdateShadow” 或者 “PublishMessage”
	Type string `json:"Type"`
	// 最后任务更新时间，Unix 时间戳
	UpdateTime int64 `json:"UpdateTime"`
}
