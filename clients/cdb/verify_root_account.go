package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 验证root账号权限
// https://cloud.tencent.com/document/api/236/17495
type VerifyRootAccountRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 实例ROOT账号的密码。
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
}

func (req *VerifyRootAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*VerifyRootAccountResponse, error) {
	resp := &VerifyRootAccountResponse{}
	err := client.Request("cdb", "VerifyRootAccount", "2017-03-20").Do(req, resp)
	return resp, err
}

type VerifyRootAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果
	AsyncRequestId string `json:"AsyncRequestId"`
}
