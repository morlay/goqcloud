package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置密码
// https://cloud.tencent.com/document/api/239/20014

type ResetPasswordRequest struct {
	// Redis实例ID
	InstanceId string `name:"InstanceId"`
	// 重置的密码
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResetPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetPasswordResponse, error) {
	resp := &ResetPasswordResponse{}
	err := client.Request("redis", "ResetPassword", "2018-04-12").Do(req, resp)
	return resp, err
}

type ResetPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId int64 `json:"TaskId"`
}
