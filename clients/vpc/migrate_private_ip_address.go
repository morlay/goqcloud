package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 弹性网卡内网IP迁移
// https://cloud.tencent.com/document/api/215/15820

type MigratePrivateIPAddressRequest struct {
	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId string `name:"DestinationNetworkInterfaceId"`
	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress string `name:"PrivateIpAddress"`
	// 区域
	Region string `name:"Region"`
	// 当内网IP绑定的弹性网卡实例ID，例如：eni-m6dyj72l。
	SourceNetworkInterfaceId string `name:"SourceNetworkInterfaceId"`
}

func (req *MigratePrivateIPAddressRequest) Invoke(client github_com_morlay_goqcloud.Client) (*MigratePrivateIPAddressResponse, error) {
	resp := &MigratePrivateIPAddressResponse{}
	err := client.Request("vpc", "MigratePrivateIpAddress", "2017-03-12").Do(req, resp)
	return resp, err
}

type MigratePrivateIPAddressResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
