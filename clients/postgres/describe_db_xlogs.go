package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取实例Xlog列表
// https://cloud.tencent.com/document/api/409/18091

type DescribeDbXlogsRequest struct {
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId string `name:"DBInstanceId"`
	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime time.Time `name:"EndTime"`
	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `name:"Limit,omitempty"`
	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbXlogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbXlogsResponse, error) {
	resp := &DescribeDbXlogsResponse{}
	err := client.Request("postgres", "DescribeDBXlogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbXlogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 表示此次返回结果有多少条数据。
	TotalCount int64 `json:"TotalCount"`
	// Xlog列表
	XlogList []*Xlog `json:"XlogList"`
}
