package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 将实例转至其他项目
// https://cloud.tencent.com/document/api/409/18100

type ModifyDbInstancesProjectRequest struct {
	// postgresql实例ID数组
	DBInstanceIdSet []*string `name:"DBInstanceIdSet"`
	// postgresql实例所属新项目的ID
	ProjectId string `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstancesProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstancesProjectResponse, error) {
	resp := &ModifyDbInstancesProjectResponse{}
	err := client.Request("postgres", "ModifyDBInstancesProject", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDbInstancesProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 转移项目成功的实例个数
	Count int64 `json:"Count"`
}
