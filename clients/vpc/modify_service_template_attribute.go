package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改协议端口模板
// https://cloud.tencent.com/document/api/215/16722
type ModifyServiceTemplateAttributeRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板实例ID，例如：ppm-529nwwj8。
	ServiceTemplateId string `name:"ServiceTemplateId"`
	// 协议端口模板名称。
	ServiceTemplateName *string `name:"ServiceTemplateName,omitempty"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `name:"Services,omitempty"`
}

func (req *ModifyServiceTemplateAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyServiceTemplateAttributeResponse, error) {
	resp := &ModifyServiceTemplateAttributeResponse{}
	err := client.Request("vpc", "ModifyServiceTemplateAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyServiceTemplateAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
