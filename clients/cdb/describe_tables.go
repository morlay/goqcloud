package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库表
// https://cloud.tencent.com/document/api/236/18727

type DescribeTablesRequest struct {
	// 数据库的名称。
	Database string `name:"Database"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 单次请求返回的数量，默认值为20，最大值为2000。
	Limit *int64 `name:"Limit,omitempty"`
	// 记录偏移量，默认值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 匹配数据库表名的正则表达式，规则同MySQL官网
	TableRegexp *string `name:"TableRegexp,omitempty"`
}

func (req *DescribeTablesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTablesResponse, error) {
	resp := &DescribeTablesResponse{}
	err := client.Request("cdb", "DescribeTables", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeTablesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的数据库表信息。
	Items []*string `json:"Items"`
	// 符合查询条件的数据库表总数。
	TotalCount int64 `json:"TotalCount"`
}
