package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取创建多设备任务状态
// https://cloud.tencent.com/document/api/634/19492

type DescribeMultiDevTaskRequest struct {
	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 任务 ID，由批量创建设备接口返回
	TaskId string `name:"TaskId"`
}

func (req *DescribeMultiDevTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMultiDevTaskResponse, error) {
	resp := &DescribeMultiDevTaskResponse{}
	err := client.Request("iotcloud", "DescribeMultiDevTask", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeMultiDevTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务 ID
	TaskId string `json:"TaskId"`
	// 任务是否完成。0 代表任务未开始，1 代表任务正在执行，2 代表任务已完成
	TaskStatus int64 `json:"TaskStatus"`
}
