package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询备份数据库列表
// https://cloud.tencent.com/document/api/236/15840

type DescribeBackupDatabasesRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 分页大小，最小值为1，最大值为2000。
	Limit *int64 `name:"Limit,omitempty"`
	// 分页偏移量。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要查询的数据库名前缀。
	SearchDatabase *string `name:"SearchDatabase,omitempty"`
	// 开始时间，格式为：2017-07-12 10:29:20。
	StartTime string `name:"StartTime"`
}

func (req *DescribeBackupDatabasesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupDatabasesResponse, error) {
	resp := &DescribeBackupDatabasesResponse{}
	err := client.Request("cdb", "DescribeBackupDatabases", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeBackupDatabasesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合查询条件的数据库数组
	Items []*DatabaseName `json:"Items"`
	// 返回的数据个数
	TotalCount int64 `json:"TotalCount"`
}
