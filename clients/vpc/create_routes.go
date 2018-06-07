package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建路由策略
// https://cloud.tencent.com/document/api/215/16724

type CreateRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID。
	RouteTableId string `name:"RouteTableId"`
	// 路由策略对象。
	Routes []*string `name:"Routes"`
}

func (req *CreateRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateRoutesResponse, error) {
	resp := &CreateRoutesResponse{}
	err := client.Request("vpc", "CreateRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
