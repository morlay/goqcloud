package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// HAVIP绑定EIP
// https://cloud.tencent.com/document/api/215/30649

type HaVipAssociateAddressIPRequest struct {
	// 弹性公网IP。必须是没有绑定HAVIP的EIP
	AddressIp string `name:"AddressIp"`
	// HAVIP唯一ID，形如：havip-9o233uri。必须是没有绑定EIP的HAVIP
	HaVipId string `name:"HaVipId"`
	// 区域
	Region string `name:"Region"`
}

func (req *HaVipAssociateAddressIPRequest) Invoke(client github_com_morlay_goqcloud.Client) (*HaVipAssociateAddressIPResponse, error) {
	resp := &HaVipAssociateAddressIPResponse{}
	err := client.Request("vpc", "HaVipAssociateAddressIp", "2017-03-12").Do(req, resp)
	return resp, err
}

type HaVipAssociateAddressIPResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
