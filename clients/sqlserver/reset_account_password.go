package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置账户密码
// https://cloud.tencent.com/document/api/238/19952

type ResetAccountPasswordRequest struct {
	// 更新后的账户密码信息数组
	Accounts []*AccountPassword `name:"Accounts"`
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResetAccountPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetAccountPasswordResponse, error) {
	resp := &ResetAccountPasswordResponse{}
	err := client.Request("sqlserver", "ResetAccountPassword", "2018-03-28").Do(req, resp)
	return resp, err
}

type ResetAccountPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 修改帐号密码的异步任务流程ID
	FlowId int64 `json:"FlowId"`
}
