package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改IP地址模板集合
// https://cloud.tencent.com/document/api/215/16721

type ModifyAddressTemplateGroupAttributeRequest struct {
	// IP地址模板集合实例ID，例如：ipmg-2uw6ujo6。
	AddressTemplateGroupId string `name:"AddressTemplateGroupId"`
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `name:"AddressTemplateGroupName,omitempty"`
	// IP地址模板实例ID， 例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `name:"AddressTemplateIds,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAddressTemplateGroupAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAddressTemplateGroupAttributeResponse, error) {
	resp := &ModifyAddressTemplateGroupAttributeResponse{}
	err := client.Request("vpc", "ModifyAddressTemplateGroupAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
