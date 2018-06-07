package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 弹性网卡退还内网 IP
// https://cloud.tencent.com/document/api/215/15814

type UnassignPrivateIPAddressesRequest struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `name:"PrivateIpAddresses"`
	// 区域
	Region string `name:"Region"`
}

func (req *UnassignPrivateIPAddressesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UnassignPrivateIPAddressesResponse, error) {
	resp := &UnassignPrivateIPAddressesResponse{}
	err := client.Request("vpc", "UnassignPrivateIpAddresses", "2017-03-12").Do(req, resp)
	return resp, err
}

type UnassignPrivateIPAddressesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
