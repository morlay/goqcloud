package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改子网属性
// https://cloud.tencent.com/document/api/215/15781
type ModifySubnetAttributeRequest struct {
	// 子网是否开启广播。
	EnableBroadcast *string `name:"EnableBroadcast,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId string `name:"SubnetId"`
	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `name:"SubnetName,omitempty"`
}

func (req *ModifySubnetAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySubnetAttributeResponse, error) {
	resp := &ModifySubnetAttributeResponse{}
	err := client.Request("vpc", "ModifySubnetAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifySubnetAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
