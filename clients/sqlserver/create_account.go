package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建实例账号
// https://cloud.tencent.com/document/api/238/19975

type CreateAccountRequest struct {
	// 数据库实例账户信息
	Accounts []*AccountCreateInfo `name:"Accounts"`
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := client.Request("sqlserver", "CreateAccount", "2018-03-28").Do(req, resp)
	return resp, err
}

type CreateAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务流id
	FlowId int64 `json:"FlowId"`
}
