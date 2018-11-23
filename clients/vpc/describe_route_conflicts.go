package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询路由策略冲突列表
// https://cloud.tencent.com/document/api/215/19217

type DescribeRouteConflictsRequest struct {
	// 要检查的与之冲突的目的端列表
	DestinationCidrBlocks []*string `name:"DestinationCidrBlocks"`
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
}

func (req *DescribeRouteConflictsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRouteConflictsResponse, error) {
	resp := &DescribeRouteConflictsResponse{}
	err := client.Request("vpc", "DescribeRouteConflicts", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeRouteConflictsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 路由策略冲突列表
	RouteConflictSet []*RouteConflict `json:"RouteConflictSet"`
}
