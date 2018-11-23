package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改迁移任务
// https://cloud.tencent.com/document/api/238/19940

type ModifyMigrationRequest struct {
	// 迁移DB对象 ，离线迁移（SourceType=4或SourceType=5）不使用，若不填则不修改
	MigrateDBSet []*MigrateDB `name:"MigrateDBSet,omitempty"`
	// 迁移任务ID
	MigrateId int64 `name:"MigrateId"`
	// 新的迁移任务的名称，若不填则不修改
	MigrateName *string `name:"MigrateName,omitempty"`
	// 新的迁移类型（1:结构迁移 2:数据迁移 3:增量同步），若不填则不修改
	MigrateType *int64 `name:"MigrateType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 迁移源，若不填则不修改
	Source *MigrateSource `name:"Source,omitempty"`
	// 迁移源的类型 1:CDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式），若不填则不修改
	SourceType *int64 `name:"SourceType,omitempty"`
	// 迁移目标，若不填则不修改
	Target *MigrateTarget `name:"Target,omitempty"`
}

func (req *ModifyMigrationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyMigrationResponse, error) {
	resp := &ModifyMigrationResponse{}
	err := client.Request("sqlserver", "ModifyMigration", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyMigrationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移任务ID
	MigrateId int64 `json:"MigrateId"`
}
