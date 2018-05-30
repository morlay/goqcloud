package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 弹性网卡迁移
// https://cloud.tencent.com/document/api/215/15821
type MigrateNetworkInterfaceRequest struct {
	// 待迁移的目的CVM实例ID。
	DestinationInstanceId string `name:"DestinationInstanceId"`
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 区域
	Region string `name:"Region"`
	// 弹性网卡当前绑定的CVM实例ID。形如：ins-r8hr2upy。
	SourceInstanceId string `name:"SourceInstanceId"`
}

func (req *MigrateNetworkInterfaceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*MigrateNetworkInterfaceResponse, error) {
	resp := &MigrateNetworkInterfaceResponse{}
	err := client.Request("vpc", "MigrateNetworkInterface", "2017-03-12").Do(req, resp)
	return resp, err
}

type MigrateNetworkInterfaceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
