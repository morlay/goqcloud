package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云联网各地域出带宽上限
// https://cloud.tencent.com/document/api/215/19201

type DescribeCcnRegionBandwidthLimitsRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCcnRegionBandwidthLimitsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCcnRegionBandwidthLimitsResponse, error) {
	resp := &DescribeCcnRegionBandwidthLimitsResponse{}
	err := client.Request("vpc", "DescribeCcnRegionBandwidthLimits", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCcnRegionBandwidthLimitsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云联网（CCN）各地域出带宽上限
	CcnRegionBandwidthLimitSet []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimitSet"`
}
