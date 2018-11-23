package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除带宽包资源
// https://cloud.tencent.com/document/api/215/19207

type RemoveBandwidthPackageResourcesRequest struct {
	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `name:"BandwidthPackageId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 资源Id，形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `name:"ResourceIds,omitempty"`
	// 资源类型，包括‘Address’, ‘LoadBalance’
	ResourceType *string `name:"ResourceType,omitempty"`
}

func (req *RemoveBandwidthPackageResourcesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RemoveBandwidthPackageResourcesResponse, error) {
	resp := &RemoveBandwidthPackageResourcesResponse{}
	err := client.Request("vpc", "RemoveBandwidthPackageResources", "2017-03-12").Do(req, resp)
	return resp, err
}

type RemoveBandwidthPackageResourcesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
