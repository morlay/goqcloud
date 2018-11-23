package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询迁移任务详情
// https://cloud.tencent.com/document/api/238/19942

type DescribeMigrationDetailRequest struct {
	// 迁移任务ID
	MigrateId int64 `name:"MigrateId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeMigrationDetailRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMigrationDetailResponse, error) {
	resp := &DescribeMigrationDetailResponse{}
	err := client.Request("sqlserver", "DescribeMigrationDetail", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeMigrationDetailResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移任务所属的用户ID
	AppId int64 `json:"AppId"`
	// 迁移任务的创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 迁移任务的结束时间
	EndTime time.Time `json:"EndTime"`
	// 迁移DB对象 ，离线迁移（SourceType=4或SourceType=5）不使用。
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet"`
	// 迁移任务ID
	MigrateId int64 `json:"MigrateId"`
	// 迁移任务名称
	MigrateName string `json:"MigrateName"`
	// 迁移类型（1:结构迁移 2:数据迁移 3:增量同步）
	MigrateType int64 `json:"MigrateType"`
	// 迁移任务当前进度
	Progress int64 `json:"Progress"`
	// 迁移任务所属的地域
	Region string `json:"Region"`
	// 迁移源
	Source MigrateSource `json:"Source"`
	// 迁移源的类型 1:CDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType int64 `json:"SourceType"`
	// 迁移任务的开始时间
	StartTime time.Time `json:"StartTime"`
	// 迁移任务的状态（1:初始化,4:迁移中,5.迁移失败,6.迁移成功）
	Status int64 `json:"Status"`
	// 迁移目标
	Target MigrateTarget `json:"Target"`
}
