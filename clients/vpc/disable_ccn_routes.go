package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 禁用云联网路由
// https://cloud.tencent.com/document/api/215/19197

type DisableCcnRoutesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 区域
	Region string `name:"Region"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `name:"RouteIds"`
}

func (req *DisableCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DisableCcnRoutesResponse, error) {
	resp := &DisableCcnRoutesResponse{}
	err := client.Request("vpc", "DisableCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DisableCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
