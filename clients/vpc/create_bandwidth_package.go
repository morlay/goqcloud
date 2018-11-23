package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建共享带宽包
// https://cloud.tencent.com/document/api/215/19212

type CreateBandwidthPackageRequest struct {
	// 带宽包数量(非上移账户只能填1)
	BandwidthPackageCount *int64 `name:"BandwidthPackageCount,omitempty"`
	// 带宽包名字
	BandwidthPackageName *string `name:"BandwidthPackageName,omitempty"`
	// 带宽包计费类型，包括‘TOP5_POSTPAID_BY_MONTH’，‘PERCENT95_POSTPAID_BY_MONTH’
	ChargeType *string `name:"ChargeType,omitempty"`
	// 带宽包限速大小。单位：Mbps，-1表示不限速。
	InternetMaxBandwidth *int64 `name:"InternetMaxBandwidth,omitempty"`
	// 带宽包类型，包括'BGP'，'SINGLEISP'，'ANYCAST'
	NetworkType *string `name:"NetworkType,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateBandwidthPackageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateBandwidthPackageResponse, error) {
	resp := &CreateBandwidthPackageResponse{}
	err := client.Request("vpc", "CreateBandwidthPackage", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateBandwidthPackageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 带宽包Id
	BandwidthPackageId string `json:"BandwidthPackageId"`
	// 带宽包Ids(申请数量大于1时有效)
	BandwidthPackageIds []*string `json:"BandwidthPackageIds"`
}
