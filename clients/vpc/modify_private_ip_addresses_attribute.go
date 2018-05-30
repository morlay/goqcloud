package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改弹性网卡内网IP信息
// https://cloud.tencent.com/document/api/215/15823
type ModifyPrivateIPAddressesAttributeRequest struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `name:"PrivateIpAddresses"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyPrivateIPAddressesAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPrivateIPAddressesAttributeResponse, error) {
	resp := &ModifyPrivateIPAddressesAttributeResponse{}
	err := client.Request("vpc", "ModifyPrivateIpAddressesAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyPrivateIPAddressesAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
