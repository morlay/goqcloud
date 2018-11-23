package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询指定数据库的备份数据表
// https://cloud.tencent.com/document/api/236/15846

type DescribeBackupTablesRequest struct {
	// 指定的数据库名。
	DatabaseName string `name:"DatabaseName"`
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 分页大小，最小值为1，最大值为2000。
	Limit *int64 `name:"Limit,omitempty"`
	// 分页偏移。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要查询的数据表名前缀。
	SearchTable *string `name:"SearchTable,omitempty"`
	// 开始时间，格式为：2017-07-12 10:29:20。
	StartTime string `name:"StartTime"`
}

func (req *DescribeBackupTablesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupTablesResponse, error) {
	resp := &DescribeBackupTablesResponse{}
	err := client.Request("cdb", "DescribeBackupTables", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeBackupTablesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合条件的数据表数组。
	Items []*TableName `json:"Items"`
	// 返回的数据个数。
	TotalCount int64 `json:"TotalCount"`
}
