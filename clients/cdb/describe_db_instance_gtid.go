package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据实例的GTID是否开通
// https://cloud.tencent.com/document/api/236/15862

type DescribeDbInstanceGtidRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceGtidRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceGtidResponse, error) {
	resp := &DescribeDbInstanceGtidResponse{}
	err := client.Request("cdb", "DescribeDBInstanceGTID", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceGtidResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// GTID是否开通的标记：0-未开通，1-已开通。
	IsGTIDOpen int64 `json:"IsGTIDOpen"`
}
