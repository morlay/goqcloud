package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改Redis密码
// https://cloud.tencent.com/document/api/239/20025

type ModfiyInstancePasswordRequest struct {
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 实例旧密码
	OldPassword string `name:"OldPassword"`
	// 实例新密码
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModfiyInstancePasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModfiyInstancePasswordResponse, error) {
	resp := &ModfiyInstancePasswordResponse{}
	err := client.Request("redis", "ModfiyInstancePassword", "2018-04-12").Do(req, resp)
	return resp, err
}

type ModfiyInstancePasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId int64 `json:"TaskId"`
}
