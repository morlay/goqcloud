package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建迁移任务
// https://cloud.tencent.com/document/api/238/19945

type CreateMigrationRequest struct {
	// 迁移DB对象 ，离线迁移不使用（SourceType=4或SourceType=5）。
	MigrateDBSet []*MigrateDB `name:"MigrateDBSet,omitempty"`
	// 迁移任务的名称
	MigrateName string `name:"MigrateName"`
	// 迁移类型（1:结构迁移 2:数据迁移 3:增量同步）
	MigrateType int64 `name:"MigrateType"`
	// 区域
	Region string `name:"Region"`
	// 迁移源
	Source MigrateSource `name:"Source"`
	// 迁移源的类型 1:CDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType int64 `name:"SourceType"`
	// 迁移目标
	Target MigrateTarget `name:"Target"`
}

func (req *CreateMigrationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateMigrationResponse, error) {
	resp := &CreateMigrationResponse{}
	err := client.Request("sqlserver", "CreateMigration", "2018-03-28").Do(req, resp)
	return resp, err
}

type CreateMigrationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移任务ID
	MigrateId int64 `json:"MigrateId"`
}
