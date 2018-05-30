package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 弹性网卡申请内网 IP
// https://cloud.tencent.com/document/api/215/15813
type AssignPrivateIPAddressesRequest struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `name:"PrivateIpAddresses,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 新申请的内网IP地址个数。
	SecondaryPrivateIpAddressCount *int64 `name:"SecondaryPrivateIpAddressCount,omitempty"`
}

func (req *AssignPrivateIPAddressesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AssignPrivateIPAddressesResponse, error) {
	resp := &AssignPrivateIPAddressesResponse{}
	err := client.Request("vpc", "AssignPrivateIpAddresses", "2017-03-12").Do(req, resp)
	return resp, err
}

type AssignPrivateIPAddressesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 内网IP详细信息。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet"`
}
