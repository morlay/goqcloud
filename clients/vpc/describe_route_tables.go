package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询路由列表
// https://cloud.tencent.com/document/api/215/15763

type DescribeRouteTablesRequest struct {
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。route-table-id - String - （过滤条件）路由表实例ID。route-table-name - String - （过滤条件）路由表名称。vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。association.main - String - （过滤条件）是否主路由表。
	Filters []*Filter `name:"Filters,omitempty"`
	// 请求对象个数。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `name:"RouteTableIds,omitempty"`
}

func (req *DescribeRouteTablesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRouteTablesResponse, error) {
	resp := &DescribeRouteTablesResponse{}
	err := client.Request("vpc", "DescribeRouteTables", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeRouteTablesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 路由表对象。
	RouteTableSet []*RouteTable `json:"RouteTableSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
