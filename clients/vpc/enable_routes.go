package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 启用子网路由
// https://cloud.tencent.com/document/api/215/19215

type EnableRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由策略唯一ID。
	RouteIds []*int64 `name:"RouteIds"`
	// 路由表唯一ID。
	RouteTableId string `name:"RouteTableId"`
}

func (req *EnableRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*EnableRoutesResponse, error) {
	resp := &EnableRoutesResponse{}
	err := client.Request("vpc", "EnableRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type EnableRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
