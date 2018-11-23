package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新数据库名
// https://cloud.tencent.com/document/api/238/19955

type ModifyDbNameRequest struct {
	// 实例id
	InstanceId string `name:"InstanceId"`
	// 新数据库名
	NewDBName string `name:"NewDBName"`
	// 旧数据库名
	OldDBName string `name:"OldDBName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbNameRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbNameResponse, error) {
	resp := &ModifyDbNameResponse{}
	err := client.Request("sqlserver", "ModifyDBName", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyDbNameResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务流id
	FlowId int64 `json:"FlowId"`
}
