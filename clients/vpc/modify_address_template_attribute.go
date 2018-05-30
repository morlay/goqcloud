package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改IP地址模板
// https://cloud.tencent.com/document/api/215/16720
type ModifyAddressTemplateAttributeRequest struct {
	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateId string `name:"AddressTemplateId"`
	// IP地址模板名称。
	AddressTemplateName *string `name:"AddressTemplateName,omitempty"`
	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `name:"Addresses,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAddressTemplateAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAddressTemplateAttributeResponse, error) {
	resp := &ModifyAddressTemplateAttributeResponse{}
	err := client.Request("vpc", "ModifyAddressTemplateAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAddressTemplateAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
