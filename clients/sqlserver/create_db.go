package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建数据库
// https://cloud.tencent.com/document/api/238/19974

type CreateDbRequest struct {
	// 数据库创建信息
	DBs []*DBCreateInfo `name:"DBs"`
	// 实例id
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateDbRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbResponse, error) {
	resp := &CreateDbResponse{}
	err := client.Request("sqlserver", "CreateDB", "2018-03-28").Do(req, resp)
	return resp, err
}

type CreateDbResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务流id
	FlowId int64 `json:"FlowId"`
}
