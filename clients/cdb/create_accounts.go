package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建云数据库的账户
// https://cloud.tencent.com/document/api/236/17502
type CreateAccountsRequest struct {
	// 云数据库账号。
	Accounts []*Account `name:"Accounts"`
	// 备注信息。
	Description *string `name:"Description,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 新账户的密码。
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAccountsResponse, error) {
	resp := &CreateAccountsResponse{}
	err := client.Request("cdb", "CreateAccounts", "2017-03-20").Do(req, resp)
	return resp, err
}

type CreateAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}
