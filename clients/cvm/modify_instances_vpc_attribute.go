package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例vpc属性
// https://cloud.tencent.com/document/api/213/15750

type ModifyInstancesVpcAttributeRequest struct {
	// 是否对运行中的实例选择强制关机。默认为TRUE。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 待操作的实例ID数组。可通过DescribeInstances接口返回值中的InstanceId获取。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 私有网络相关信息配置。通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。当指定私有网络ID和子网ID（子网必须在实例所在的可用区）与指定实例所在私有网络不一致时，会将实例迁移至指定的私有网络的子网下。可通过PrivateIpAddresses指定私有网络子网IP，若需指定则所有已指定的实例均需要指定子网IP，此时InstanceIds与PrivateIpAddresses一一对应。不指定PrivateIpAddresses时随机分配私有网络子网IP。
	VirtualPrivateCloud VirtualPrivateCloud `name:"VirtualPrivateCloud"`
}

func (req *ModifyInstancesVpcAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstancesVpcAttributeResponse, error) {
	resp := &ModifyInstancesVpcAttributeResponse{}
	err := client.Request("cvm", "ModifyInstancesVpcAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyInstancesVpcAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
