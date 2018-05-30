package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除弹性网卡
// https://cloud.tencent.com/document/api/215/15822
type DeleteNetworkInterfaceRequest struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteNetworkInterfaceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteNetworkInterfaceResponse, error) {
	resp := &DeleteNetworkInterfaceResponse{}
	err := client.Request("vpc", "DeleteNetworkInterface", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteNetworkInterfaceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
