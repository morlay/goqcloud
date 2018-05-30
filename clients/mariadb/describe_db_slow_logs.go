package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询慢查询日志列表
// https://cloud.tencent.com/document/api/237/16159
type DescribeDbSlowLogsRequest struct {
	// 要查询的具体数据库名称
	Db *string `name:"Db,omitempty"`
	// 查询的结束时间，形如2016-08-22 14:55:20
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 返回的结果条数
	Limit int64 `name:"Limit"`
	// 从结果的第几条数据开始返回
	Offset int64 `name:"Offset"`
	// 排序指标，取值为query_time_sum或者query_count
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序类型，desc或者asc
	OrderByType *string `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询的起始时间，形如2016-07-23 14:55:20
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbSlowLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbSlowLogsResponse, error) {
	resp := &DescribeDbSlowLogsResponse{}
	err := client.Request("mariadb", "DescribeDBSlowLogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbSlowLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 慢查询日志数据
	Data []*SlowLogData `json:"Data"`
	// 所有语句锁时间总和
	LockTimeSum string `json:"LockTimeSum"`
	// 所有语句查询总次数
	QueryCount string `json:"QueryCount"`
	// 所有语句查询时间总和
	QueryTimeSum string `json:"QueryTimeSum"`
	// 总记录数
	Total string `json:"Total"`
}
