package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 添加带宽包资源
// https://cloud.tencent.com/document/api/215/19213

type AddBandwidthPackageResourcesRequest struct {
	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `name:"BandwidthPackageId,omitempty"`
	// 带宽包类型，包括'BGP', 'SINGLEISP', 'ANYCAST'
	NetworkType *string `name:"NetworkType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 资源Id，形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `name:"ResourceIds"`
	// 资源类型，包括'Address', 'LoadBalance'
	ResourceType *string `name:"ResourceType,omitempty"`
}

func (req *AddBandwidthPackageResourcesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddBandwidthPackageResourcesResponse, error) {
	resp := &AddBandwidthPackageResourcesResponse{}
	err := client.Request("vpc", "AddBandwidthPackageResources", "2017-03-12").Do(req, resp)
	return resp, err
}

type AddBandwidthPackageResourcesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
