package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据迁移任务
// https://cloud.tencent.com/document/api/571/18130

type DescribeMigrateJobsRequest struct {
	// 数据迁移任务ID
	JobId *string `name:"JobId,omitempty"`
	// 数据迁移任务名称
	JobName *string `name:"JobName,omitempty"`
	// 返回实例数量，默认20，有效区间[1,100]
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 排序字段，可以取值为JobId、Status、JobName、MigrateType、RunMode、CreateTime
	Order *string `name:"Order,omitempty"`
	// 排序方式，升序为ASC，降序为DESC
	OrderSeq *string `name:"OrderSeq,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeMigrateJobsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMigrateJobsResponse, error) {
	resp := &DescribeMigrateJobsResponse{}
	err := client.Request("dts", "DescribeMigrateJobs", "2018-03-30").Do(req, resp)
	return resp, err
}

type DescribeMigrateJobsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务详情数组
	JobList []*MigrateJobInfo `json:"JobList"`
	// 任务数目
	TotalCount int64 `json:"TotalCount"`
}
