package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询慢查询日志
// https://cloud.tencent.com/document/api/236/15845

type DescribeSlowLogsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 返回记录数量，默认值为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，最小值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSlowLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSlowLogsResponse, error) {
	resp := &DescribeSlowLogsResponse{}
	err := client.Request("cdb", "DescribeSlowLogs", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeSlowLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合查询条件的慢查询日志详情
	Items []*SlowLogInfo `json:"Items"`
	// 符合查询条件的慢查询日志总数
	TotalCount int64 `json:"TotalCount"`
}
