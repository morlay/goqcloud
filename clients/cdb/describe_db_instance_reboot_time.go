package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库实例的预期重启时间
// https://cloud.tencent.com/document/api/236/15874

type DescribeDbInstanceRebootTimeRequest struct {
	// 实例的ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceRebootTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceRebootTimeResponse, error) {
	resp := &DescribeDbInstanceRebootTimeResponse{}
	err := client.Request("cdb", "DescribeDBInstanceRebootTime", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceRebootTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的参数信息。
	Items []*InstanceRebootTime `json:"Items"`
	// 符合查询条件的实例总数。
	TotalCount int64 `json:"TotalCount"`
}
