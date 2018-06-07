package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例vpc属性
// https://cloud.tencent.com/document/api/213/15750

type UpdateInstanceVpcConfigRequest struct {
	// 是否对运行中的实例选择强制关机。默认为TRUE。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 私有网络相关信息配置。通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。
	VirtualPrivateCloud VirtualPrivateCloud `name:"VirtualPrivateCloud"`
}

func (req *UpdateInstanceVpcConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateInstanceVpcConfigResponse, error) {
	resp := &UpdateInstanceVpcConfigResponse{}
	err := client.Request("cvm", "UpdateInstanceVpcConfig", "2017-03-12").Do(req, resp)
	return resp, err
}

type UpdateInstanceVpcConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
