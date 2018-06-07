package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例的属性
// https://cloud.tencent.com/document/api/213/15739

type ModifyInstancesAttributeRequest struct {
	// 一个或多个待操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `name:"InstanceIds"`
	// 实例名称。可任意命名，但不得超过60个字符。
	InstanceName *string `name:"InstanceName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyInstancesAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstancesAttributeResponse, error) {
	resp := &ModifyInstancesAttributeResponse{}
	err := client.Request("cvm", "ModifyInstancesAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyInstancesAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
