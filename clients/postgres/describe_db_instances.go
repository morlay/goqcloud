package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例列表
// https://cloud.tencent.com/document/api/409/16773

type DescribeDbInstancesRequest struct {
	// 过滤条件，目前支持：db-instance-id、db-instance-name两种。
	Filters []*Filter `name:"Filters,omitempty"`
	// 分页序号，从1开始。
	PageNumber *int64 `name:"PageNumber,omitempty"`
	// 每页显示数量，默认返回10条。
	PageSize *int64 `name:"PageSize,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstancesResponse, error) {
	resp := &DescribeDbInstancesResponse{}
	err := client.Request("postgres", "DescribeDBInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息集合。
	DBInstanceSet []*DBInstance `json:"DBInstanceSet"`
	// 查询到的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
