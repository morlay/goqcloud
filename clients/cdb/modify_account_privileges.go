package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例账号的权限
// https://cloud.tencent.com/document/api/236/17496
type ModifyAccountPrivilegesRequest struct {
	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `name:"Accounts"`
	// 数据库表中列的权限。Privileges权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	ColumnPrivileges []*ColumnPrivilege `name:"ColumnPrivileges,omitempty"`
	// 数据库的权限。Privileges权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",   "DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	DatabasePrivileges []*DatabasePrivilege `name:"DatabasePrivileges,omitempty"`
	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",    "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	GlobalPrivileges []*string `name:"GlobalPrivileges,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 数据库中表的权限。Privileges权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE",  "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	TablePrivileges []*TablePrivilege `name:"TablePrivileges,omitempty"`
}

func (req *ModifyAccountPrivilegesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountPrivilegesResponse, error) {
	resp := &ModifyAccountPrivilegesResponse{}
	err := client.Request("cdb", "ModifyAccountPrivileges", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyAccountPrivilegesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}
