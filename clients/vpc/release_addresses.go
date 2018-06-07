package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 释放弹性公网IP
// https://cloud.tencent.com/document/api/215/16705

type ReleaseAddressesRequest struct {
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：eip-11112222。
	AddressIds []*string `name:"AddressIds,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ReleaseAddressesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ReleaseAddressesResponse, error) {
	resp := &ReleaseAddressesResponse{}
	err := client.Request("vpc", "ReleaseAddresses", "2017-03-12").Do(req, resp)
	return resp, err
}

type ReleaseAddressesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
