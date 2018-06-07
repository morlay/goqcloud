package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 弹性网卡绑定云主机
// https://cloud.tencent.com/document/api/215/15819

type AttachNetworkInterfaceRequest struct {
	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId string `name:"InstanceId"`
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *AttachNetworkInterfaceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AttachNetworkInterfaceResponse, error) {
	resp := &AttachNetworkInterfaceResponse{}
	err := client.Request("vpc", "AttachNetworkInterface", "2017-03-12").Do(req, resp)
	return resp, err
}

type AttachNetworkInterfaceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
