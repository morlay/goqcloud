package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例名字
// https://cloud.tencent.com/document/api/238/19958

type ModifyDbInstanceNameRequest struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 新的数据库实例名字
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceNameRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceNameResponse, error) {
	resp := &ModifyDbInstanceNameResponse{}
	err := client.Request("sqlserver", "ModifyDBInstanceName", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceNameResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
