package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除迁移任务
// https://cloud.tencent.com/document/api/238/19944

type DeleteMigrationRequest struct {
	// 迁移任务ID
	MigrateId int64 `name:"MigrateId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteMigrationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteMigrationResponse, error) {
	resp := &DeleteMigrationResponse{}
	err := client.Request("sqlserver", "DeleteMigration", "2018-03-28").Do(req, resp)
	return resp, err
}

type DeleteMigrationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
