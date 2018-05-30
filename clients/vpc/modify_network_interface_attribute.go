package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改弹性网卡
// https://cloud.tencent.com/document/api/215/15815
type ModifyNetworkInterfaceAttributeRequest struct {
	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `name:"NetworkInterfaceDescription,omitempty"`
	// 弹性网卡实例ID，例如：eni-pxir56ns。
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `name:"NetworkInterfaceName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 指定绑定的安全组，例如:['sg-1dd51d']。
	SecurityGroupIds []*string `name:"SecurityGroupIds,omitempty"`
}

func (req *ModifyNetworkInterfaceAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyNetworkInterfaceAttributeResponse, error) {
	resp := &ModifyNetworkInterfaceAttributeResponse{}
	err := client.Request("vpc", "ModifyNetworkInterfaceAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyNetworkInterfaceAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
