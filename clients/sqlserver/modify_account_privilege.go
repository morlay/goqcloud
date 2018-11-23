package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例账户权限
// https://cloud.tencent.com/document/api/238/19959

type ModifyAccountPrivilegeRequest struct {
	// 账号权限变更信息
	Accounts []*AccountPrivilegeModifyInfo `name:"Accounts"`
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAccountPrivilegeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountPrivilegeResponse, error) {
	resp := &ModifyAccountPrivilegeResponse{}
	err := client.Request("sqlserver", "ModifyAccountPrivilege", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyAccountPrivilegeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务流程ID
	FlowId int64 `json:"FlowId"`
}
