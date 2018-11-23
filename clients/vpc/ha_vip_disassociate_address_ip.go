package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// HAVIP解绑EIP
// https://cloud.tencent.com/document/api/215/30648

type HaVipDisassociateAddressIPRequest struct {
	// HAVIP唯一ID，形如：havip-9o233uri。必须是已绑定EIP的HAVIP。
	HaVipId string `name:"HaVipId"`
	// 区域
	Region string `name:"Region"`
}

func (req *HaVipDisassociateAddressIPRequest) Invoke(client github_com_morlay_goqcloud.Client) (*HaVipDisassociateAddressIPResponse, error) {
	resp := &HaVipDisassociateAddressIPResponse{}
	err := client.Request("vpc", "HaVipDisassociateAddressIp", "2017-03-12").Do(req, resp)
	return resp, err
}

type HaVipDisassociateAddressIPResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
