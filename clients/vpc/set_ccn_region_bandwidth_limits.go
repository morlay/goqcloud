package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置云联网各地域出带宽上限
// https://cloud.tencent.com/document/api/215/19194

type SetCcnRegionBandwidthLimitsRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 云联网（CCN）各地域出带宽上限。
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `name:"CcnRegionBandwidthLimits"`
	// 区域
	Region string `name:"Region"`
}

func (req *SetCcnRegionBandwidthLimitsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SetCcnRegionBandwidthLimitsResponse, error) {
	resp := &SetCcnRegionBandwidthLimitsResponse{}
	err := client.Request("vpc", "SetCcnRegionBandwidthLimits", "2017-03-12").Do(req, resp)
	return resp, err
}

type SetCcnRegionBandwidthLimitsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
