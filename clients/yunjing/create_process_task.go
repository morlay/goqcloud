package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建实时拉取进程任务
// https://cloud.tencent.com/document/api/296/30341

type CreateProcessTaskRequest struct {
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。
	Uuid string `name:"Uuid"`
}

func (req *CreateProcessTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateProcessTaskResponse, error) {
	resp := &CreateProcessTaskResponse{}
	err := client.Request("yunjing", "CreateProcessTask", "2018-02-28").Do(req, resp)
	return resp, err
}

type CreateProcessTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
