package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置路由表名称和路由策略
// https://cloud.tencent.com/document/api/215/15769
type ResetRoutesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId string `name:"RouteTableId"`
	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName string `name:"RouteTableName"`
	// 路由策略。
	Routes []*Route `name:"Routes"`
}

func (req *ResetRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetRoutesResponse, error) {
	resp := &ResetRoutesResponse{}
	err := client.Request("vpc", "ResetRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
