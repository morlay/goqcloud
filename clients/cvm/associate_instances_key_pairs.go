package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 绑定密钥对
// https://cloud.tencent.com/document/api/213/15698

type AssociateInstancesKeyPairsRequest struct {
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：TRUE：表示在正常关机失败后进行强制关机。FALSE：表示在正常关机失败后不进行强制关机。默认取值：FALSE。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。可以通过以下方式获取可用的实例ID：通过登录控制台查询实例ID。通过调用接口 DescribeInstances ，取返回信息中的InstanceId获取实例ID。
	InstanceIds []*string `name:"InstanceIds"`
	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。密钥对ID形如：skey-3glfot13。可以通过以下方式获取可用的密钥ID：通过登录控制台查询密钥ID。通过调用接口 DescribeKeyPairs ，取返回信息中的KeyId获取密钥对ID。
	KeyIds []*string `name:"KeyIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *AssociateInstancesKeyPairsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AssociateInstancesKeyPairsResponse, error) {
	resp := &AssociateInstancesKeyPairsResponse{}
	err := client.Request("cvm", "AssociateInstancesKeyPairs", "2017-03-12").Do(req, resp)
	return resp, err
}

type AssociateInstancesKeyPairsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
