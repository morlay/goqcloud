package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除共享带宽包
// https://cloud.tencent.com/document/api/215/19211

type DeleteBandwidthPackageRequest struct {
	// 待删除带宽包bwpId
	BandwidthPackageId string `name:"BandwidthPackageId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteBandwidthPackageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteBandwidthPackageResponse, error) {
	resp := &DeleteBandwidthPackageResponse{}
	err := client.Request("vpc", "DeleteBandwidthPackage", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteBandwidthPackageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
