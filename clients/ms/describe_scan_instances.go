package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询扫描列表
// https://cloud.tencent.com/document/api/283/17756

type DescribeScanInstancesRequest struct {
	// 支持通过app名称，app包名进行筛选
	Filters []*Filter `name:"Filters,omitempty"`
	// 可以提供ItemId数组来查询一个或者多个结果。注意不可以同时指定ItemIds和Filters。
	ItemIds []*string `name:"ItemIds,omitempty"`
	// 数量限制，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 升序（asc）还是降序（desc），默认：desc。
	OrderDirection *string `name:"OrderDirection,omitempty"`
	// 按某个字段排序，目前仅支持TaskTime排序。
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeScanInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeScanInstancesResponse, error) {
	resp := &DescribeScanInstancesResponse{}
	err := client.Request("ms", "DescribeScanInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeScanInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 一个关于app详细信息的结构体，主要包括app的基本信息和扫描状态信息。
	ScanSet []*AppScanSet `json:"ScanSet"`
	// 符合要求的app数量
	TotalCount int64 `json:"TotalCount"`
}
