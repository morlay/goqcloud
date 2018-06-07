package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除路由表
// https://cloud.tencent.com/document/api/215/15771

type DeleteRouteTableRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
}

func (req *DeleteRouteTableRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteRouteTableResponse, error) {
	resp := &DeleteRouteTableResponse{}
	err := client.Request("vpc", "DeleteRouteTable", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteRouteTableResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
