package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取错误日志
// https://cloud.tencent.com/document/api/409/18093

type DescribeDbErrlogsRequest struct {
	// 实例ID，形如postgres-5bq3wfjd
	DBInstanceId string `name:"DBInstanceId"`
	// 数据库名字
	DatabaseName *string `name:"DatabaseName,omitempty"`
	// 查询结束时间，形如2018-01-01 00:00:00
	EndTime time.Time `name:"EndTime"`
	// 分页返回，每页返回的最大数量。取值为1-100
	Limit *int64 `name:"Limit,omitempty"`
	// 分页返回，返回第几页的数据，从第0页开始计数
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 搜索关键字
	SearchKeys []*string `name:"SearchKeys,omitempty"`
	// 查询起始时间，形如2018-01-01 00:00:00，起始时间不得小于7天以前
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbErrlogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbErrlogsResponse, error) {
	resp := &DescribeDbErrlogsResponse{}
	err := client.Request("postgres", "DescribeDBErrlogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbErrlogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 错误日志列表
	Details []*ErrLogDetail `json:"Details"`
	// 本次调用返回了多少条数据
	TotalCount int64 `json:"TotalCount"`
}
