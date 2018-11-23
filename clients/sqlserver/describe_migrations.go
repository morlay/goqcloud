package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询迁移任务列表
// https://cloud.tencent.com/document/api/238/19941

type DescribeMigrationsRequest struct {
	// 每页的记录数
	Limit *int64 `name:"Limit,omitempty"`
	// 迁移任务的名称，模糊匹配
	MigrateName *string `name:"MigrateName,omitempty"`
	// 查询第几页的记录
	Offset *int64 `name:"Offset,omitempty"`
	// 查询结果按照关键字排序，可选值为name、createTime、startTime，endTime，status
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序方式，可选值为desc、asc
	OrderByType *string `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 状态集合。只要符合集合中某一状态的迁移任务，就会查出来
	StatusSet []*int64 `name:"StatusSet,omitempty"`
}

func (req *DescribeMigrationsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMigrationsResponse, error) {
	resp := &DescribeMigrationsResponse{}
	err := client.Request("sqlserver", "DescribeMigrations", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeMigrationsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 查询结果的列表
	MigrateTaskSet []*MigrateTask `json:"MigrateTaskSet"`
	// 查询结果的总数
	TotalCount int64 `json:"TotalCount"`
}
