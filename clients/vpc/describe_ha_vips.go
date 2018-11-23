package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询HAVIP列表
// https://cloud.tencent.com/document/api/215/30650

type DescribeHaVipsRequest struct {
	// 过滤条件，参数不支持同时指定HaVipIds和Filters。havip-id - String - HAVIP唯一ID，形如：havip-9o233uri。havip-name - String - HAVIP名称。vpc-id - String - HAVIP所在私有网络ID。subnet-id - String - HAVIP所在子网ID。address-ip - String - HAVIP绑定的弹性公网IP。
	Filters []*Filter `name:"Filters,omitempty"`
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipIds []*string `name:"HaVipIds,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeHaVipsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeHaVipsResponse, error) {
	resp := &DescribeHaVipsResponse{}
	err := client.Request("vpc", "DescribeHaVips", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeHaVipsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// HAVIP对象数组。
	HaVipSet []*HaVip `json:"HaVipSet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
