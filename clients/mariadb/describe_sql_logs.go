package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取SQL日志
// https://cloud.tencent.com/document/api/237/20256

type DescribeSQLLogsRequest struct {
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 拉取数量（0-1000，为0时拉取总数信息）。
	Limit *int64 `name:"Limit,omitempty"`
	// SQL日志偏移。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSQLLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSQLLogsResponse, error) {
	resp := &DescribeSQLLogsResponse{}
	err := client.Request("mariadb", "DescribeSqlLogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSQLLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的sql日志数量。
	Count int64 `json:"Count"`
	// 消息队列中的sql日志结束偏移。
	EndOffset int64 `json:"EndOffset"`
	// 返回的第一条sql日志的偏移。
	Offset int64 `json:"Offset"`
	// Sql日志列表。
	SqlItems []*SqlLogItem `json:"SqlItems"`
	// 消息队列中的sql日志起始偏移。
	StartOffset int64 `json:"StartOffset"`
	// 当前消息队列中的sql日志条目数。
	TotalCount int64 `json:"TotalCount"`
}
