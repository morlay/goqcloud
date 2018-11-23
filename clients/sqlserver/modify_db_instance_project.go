package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库实例所属项目
// https://cloud.tencent.com/document/api/238/19957

type ModifyDbInstanceProjectRequest struct {
	// 实例ID数组，形如mssql-j8kv137v
	InstanceIdSet []*string `name:"InstanceIdSet"`
	// 项目ID，为0的话表示默认项目
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceProjectResponse, error) {
	resp := &ModifyDbInstanceProjectResponse{}
	err := client.Request("sqlserver", "ModifyDBInstanceProject", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 修改成功的实例个数
	Count int64 `json:"Count"`
}
