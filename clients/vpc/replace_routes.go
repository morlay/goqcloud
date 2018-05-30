package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 替换路由策略
// https://cloud.tencent.com/document/api/215/15764
type ReplaceRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
	// 路由策略对象。只需要指定路由策略ID（RouteId）。
	Routes []*Route `name:"Routes"`
}

func (req *ReplaceRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ReplaceRoutesResponse, error) {
	resp := &ReplaceRoutesResponse{}
	err := client.Request("vpc", "ReplaceRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type ReplaceRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
