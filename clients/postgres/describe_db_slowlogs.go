package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取慢查询日志
// https://cloud.tencent.com/document/api/409/18092

type DescribeDbSlowlogsRequest struct {
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId string `name:"DBInstanceId"`
	// 数据库名字
	DatabaseName *string `name:"DatabaseName,omitempty"`
	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime time.Time `name:"EndTime"`
	// 分页返回结果，每页最大返回数量，取值为1-100，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 分页返回结果，返回结果的第几页，从0开始计数
	Offset *int64 `name:"Offset,omitempty"`
	// 按照何种指标排序，取值为sum_calls或者sum_cost_time。sum_calls-总调用次数；sum_cost_time-总的花费时间
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序规则。desc-降序；asc-升序
	OrderByType *string `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询起始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbSlowlogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbSlowlogsResponse, error) {
	resp := &DescribeDbSlowlogsResponse{}
	err := client.Request("postgres", "DescribeDBSlowlogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbSlowlogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 慢查询日志详情
	Detail SlowlogDetail `json:"Detail"`
	// 本次返回多少条数据
	TotalCount int64 `json:"TotalCount"`
}
