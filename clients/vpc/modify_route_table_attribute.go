package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改路由表属性
// https://cloud.tencent.com/document/api/215/15766

type ModifyRouteTableAttributeRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
	// 路由表名称。
	RouteTableName string `name:"RouteTableName"`
}

func (req *ModifyRouteTableAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyRouteTableAttributeResponse, error) {
	resp := &ModifyRouteTableAttributeResponse{}
	err := client.Request("vpc", "ModifyRouteTableAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyRouteTableAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
