package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 执行迁移任务
// https://cloud.tencent.com/document/api/238/19939

type RunMigrationRequest struct {
	// 迁移任务ID
	MigrateId int64 `name:"MigrateId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RunMigrationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RunMigrationResponse, error) {
	resp := &RunMigrationResponse{}
	err := client.Request("sqlserver", "RunMigration", "2018-03-28").Do(req, resp)
	return resp, err
}

type RunMigrationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移流程启动后，返回流程ID
	FlowId int64 `json:"FlowId"`
}
