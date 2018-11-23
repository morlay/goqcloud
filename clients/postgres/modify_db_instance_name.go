package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例名字
// https://cloud.tencent.com/document/api/409/18101

type ModifyDbInstanceNameRequest struct {
	// 数据库实例ID，形如postgres-6fego161
	DBInstanceId string `name:"DBInstanceId"`
	// 新的数据库实例名字
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceNameRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceNameResponse, error) {
	resp := &ModifyDbInstanceNameResponse{}
	err := client.Request("postgres", "ModifyDBInstanceName", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceNameResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
