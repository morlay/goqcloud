package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 取消任务
// https://cloud.tencent.com/document/api/634/19484

type CancelTaskRequest struct {
	// 任务 ID
	Id string `name:"Id"`
	// 区域
	Region string `name:"Region"`
}

func (req *CancelTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CancelTaskResponse, error) {
	resp := &CancelTaskResponse{}
	err := client.Request("iotcloud", "CancelTask", "2018-06-14").Do(req, resp)
	return resp, err
}

type CancelTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
