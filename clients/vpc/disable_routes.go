package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 禁用子网路由
// https://cloud.tencent.com/document/api/215/19216

type DisableRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由策略唯一ID。
	RouteIds []*int64 `name:"RouteIds"`
	// 路由表唯一ID。
	RouteTableId string `name:"RouteTableId"`
}

func (req *DisableRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DisableRoutesResponse, error) {
	resp := &DisableRoutesResponse{}
	err := client.Request("vpc", "DisableRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DisableRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
