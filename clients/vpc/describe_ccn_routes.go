package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云联网路由策略
// https://cloud.tencent.com/document/api/215/19200

type DescribeCcnRoutesRequest struct {
	// CCN实例ID，形如：ccn-gree226l。
	CcnId string `name:"CcnId"`
	// 过滤条件，参数不支持同时指定RouteIds和Filters。ccn-id - String -（过滤条件）CCN实例ID。route-id - String -（过滤条件）路由策略ID。cidr-block - String -（过滤条件）目的端。instance-type - String -（过滤条件）下一跳类型。instance-region - String -（过滤条件）下一跳所属地域。instance-id - String -（过滤条件）下一跳实例ID。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `name:"RouteIds,omitempty"`
}

func (req *DescribeCcnRoutesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCcnRoutesResponse, error) {
	resp := &DescribeCcnRoutesResponse{}
	err := client.Request("vpc", "DescribeCcnRoutes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCcnRoutesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// CCN路由策略对象。
	RouteSet []*CcnRoute `json:"RouteSet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
