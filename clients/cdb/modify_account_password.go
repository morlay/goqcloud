package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例账号的密码
// https://cloud.tencent.com/document/api/236/17497

type ModifyAccountPasswordRequest struct {
	// 云数据库账号。
	Accounts []*Account `name:"Accounts"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 数据库账号的新密码。
	NewPassword string `name:"NewPassword"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAccountPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountPasswordResponse, error) {
	resp := &ModifyAccountPasswordResponse{}
	err := client.Request("cdb", "ModifyAccountPassword", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyAccountPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}
