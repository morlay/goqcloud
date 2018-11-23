package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询带宽包资源
// https://cloud.tencent.com/document/api/215/19209

type DescribeBandwidthPackagesRequest struct {
	// 带宽包Id，支持批量
	BandwidthPackageIds []*string `name:"BandwidthPackageIds,omitempty"`
	// 每次请求的Filters的上限为10。参数不支持同时指定BandwidthPackageIds和Filters。详细的过滤条件如下： bandwidth-package_id - String - 是否必填：否 - （过滤条件）按照带宽包的唯一标识ID过滤。 bandwidth-package-name - String - 是否必填：否 - （过滤条件）按照 带宽包名称过滤。不支持模糊过滤。 network-type - String - 是否必填：否 - （过滤条件）按照带宽包的类型过滤。类型包括'BGP','SINGLEISP'和'ANYCAST'。 charge-type - String - 是否必填：否 - （过滤条件）按照带宽包的计费类型过滤。计费类型包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH' resource.resource-type - String - 是否必填：否 - （过滤条件）按照带宽包资源类型过滤。资源类型包括'Address'和'LoadBalance' resource.resource-id - String - 是否必填：否 - （过滤条件）按照带宽包资源Id过滤。资源Id形如'eip-xxxx','lb-xxxx' resource.address-ip - String - 是否必填：否 - （过滤条件）按照带宽包资源Ip过滤。
	Filters []*Filter `name:"Filters,omitempty"`
	// 查询带宽包数量限制
	Limit *int64 `name:"Limit,omitempty"`
	// 查询带宽包偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBandwidthPackagesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBandwidthPackagesResponse, error) {
	resp := &DescribeBandwidthPackagesResponse{}
	err := client.Request("vpc", "DescribeBandwidthPackages", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeBandwidthPackagesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 描述带宽包详细信息
	BandwidthPackageSet []*BandwidthPackage `json:"BandwidthPackageSet"`
	// 符合条件的带宽包数量
	TotalCount int64 `json:"TotalCount"`
}
