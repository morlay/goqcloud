package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库切换记录
// https://cloud.tencent.com/document/api/236/17490

type DescribeDbSwitchRecordsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 分页参数，单次请求数量限制。
	Limit *int64 `name:"Limit,omitempty"`
	// 分页参数，偏移量。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbSwitchRecordsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbSwitchRecordsResponse, error) {
	resp := &DescribeDbSwitchRecordsResponse{}
	err := client.Request("cdb", "DescribeDBSwitchRecords", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbSwitchRecordsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例切换记录详情。
	Items []*DBSwitchInfo `json:"Items"`
	// 实例切换记录的总数。
	TotalCount int64 `json:"TotalCount"`
}
