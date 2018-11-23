package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云联网关联实例列表
// https://cloud.tencent.com/document/api/215/19202

type DescribeCcnAttachedInstancesRequest struct {
	// 云联网实例ID
	CcnId *string `name:"CcnId,omitempty"`
	// 过滤条件：ccn-id - String -（过滤条件）CCN实例ID。instance-type - String -（过滤条件）关联实例类型。instance-region - String -（过滤条件）关联实例所属地域。instance-id - String -（过滤条件）关联实例实例ID。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCcnAttachedInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCcnAttachedInstancesResponse, error) {
	resp := &DescribeCcnAttachedInstancesResponse{}
	err := client.Request("vpc", "DescribeCcnAttachedInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCcnAttachedInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 关联实例列表。
	InstanceSet []*CcnAttachedInstance `json:"InstanceSet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
