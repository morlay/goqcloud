package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除数据库
// https://cloud.tencent.com/document/api/238/19971

type DeleteDbRequest struct {
	// 实例ID，形如mssql-rljoi3bf
	InstanceId string `name:"InstanceId"`
	// 数据库名数组
	Names []*string `name:"Names"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDbRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDbResponse, error) {
	resp := &DeleteDbResponse{}
	err := client.Request("sqlserver", "DeleteDB", "2018-03-28").Do(req, resp)
	return resp, err
}

type DeleteDbResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务流id
	FlowId int64 `json:"FlowId"`
}
