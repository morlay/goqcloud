package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改带宽包属性
// https://cloud.tencent.com/document/api/215/19208

type ModifyBandwidthPackageAttributeRequest struct {
	// 带宽包唯一标识ID
	BandwidthPackageId string `name:"BandwidthPackageId"`
	// 带宽包名称
	BandwidthPackageName string `name:"BandwidthPackageName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyBandwidthPackageAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyBandwidthPackageAttributeResponse, error) {
	resp := &ModifyBandwidthPackageAttributeResponse{}
	err := client.Request("vpc", "ModifyBandwidthPackageAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyBandwidthPackageAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
