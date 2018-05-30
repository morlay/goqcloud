package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除路由策略
// https://cloud.tencent.com/document/api/215/16725
type DeleteRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID。
	RouteTableId string `name:"RouteTableId"`
	// 路由策略对象。
	Routes []*string `name:"Routes"`
}

func (req *DeleteRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteRoutesResponse, error) {
	resp := &DeleteRoutesResponse{}
	err := client.Request("vpc", "DeleteRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
