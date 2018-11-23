package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询CCN列表
// https://cloud.tencent.com/document/api/215/19199

type DescribeCcnsRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定CcnIds和Filters。
	CcnIds []*string `name:"CcnIds,omitempty"`
	// 过滤条件，参数不支持同时指定CcnIds和Filters。ccn-id - String - （过滤条件）CCN唯一ID，形如：vpc-f49l6u0z。ccn-name - String - （过滤条件）CCN名称。ccn-description - String - （过滤条件）CCN描述。state - String - （过滤条件）实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCcnsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCcnsResponse, error) {
	resp := &DescribeCcnsResponse{}
	err := client.Request("vpc", "DescribeCcns", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCcnsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// CCN对象。
	CcnSet []*CCN `json:"CcnSet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
