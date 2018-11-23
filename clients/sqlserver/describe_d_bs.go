package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库列表
// https://cloud.tencent.com/document/api/238/19968

type DescribeDBsRequest struct {
	// 实例ID
	InstanceIdSet []*string `name:"InstanceIdSet"`
	// 每页记录数，最大为100，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 页编号，从第0页开始
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDBsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDBsResponse, error) {
	resp := &DescribeDBsResponse{}
	err := client.Request("sqlserver", "DescribeDBs", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeDBsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例数据库列表
	DBInstances []*InstanceDBDetail `json:"DBInstances"`
	// 数据库数量
	TotalCount int64 `json:"TotalCount"`
}
