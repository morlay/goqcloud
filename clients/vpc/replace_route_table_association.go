package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 替换路由表绑定关系
// https://cloud.tencent.com/document/api/215/15767

type ReplaceRouteTableAssociationRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId string `name:"SubnetId"`
}

func (req *ReplaceRouteTableAssociationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ReplaceRouteTableAssociationResponse, error) {
	resp := &ReplaceRouteTableAssociationResponse{}
	err := client.Request("vpc", "ReplaceRouteTableAssociation", "2017-03-12").Do(req, resp)
	return resp, err
}

type ReplaceRouteTableAssociationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
