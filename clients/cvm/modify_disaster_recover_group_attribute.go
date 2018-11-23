package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改分散置放群组属性
// https://cloud.tencent.com/document/api/213/17809

type ModifyDisasterRecoverGroupAttributeRequest struct {
	// 分散置放群组ID，可使用DescribeDisasterRecoverGroups接口获取。
	DisasterRecoverGroupId string `name:"DisasterRecoverGroupId"`
	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDisasterRecoverGroupAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDisasterRecoverGroupAttributeResponse, error) {
	resp := &ModifyDisasterRecoverGroupAttributeResponse{}
	err := client.Request("cvm", "ModifyDisasterRecoverGroupAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDisasterRecoverGroupAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
