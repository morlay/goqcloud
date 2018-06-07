package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询备份日志
// https://cloud.tencent.com/document/api/236/15842

type DescribeBackupsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 单次请求返回的数量，默认值为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，最小值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBackupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupsResponse, error) {
	resp := &DescribeBackupsResponse{}
	err := client.Request("cdb", "DescribeBackups", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeBackupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合查询条件的备份信息详情
	Items []*BackupInfo `json:"Items"`
	// 符合查询条件的实例总数
	TotalCount int64 `json:"TotalCount"`
}
