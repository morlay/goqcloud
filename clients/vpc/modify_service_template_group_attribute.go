package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改协议端口模板集合
// https://cloud.tencent.com/document/api/215/16723
type ModifyServiceTemplateGroupAttributeRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板集合实例ID，例如：ppmg-ei8hfd9a。
	ServiceTemplateGroupId string `name:"ServiceTemplateGroupId"`
	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `name:"ServiceTemplateGroupName,omitempty"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `name:"ServiceTemplateIds,omitempty"`
}

func (req *ModifyServiceTemplateGroupAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyServiceTemplateGroupAttributeResponse, error) {
	resp := &ModifyServiceTemplateGroupAttributeResponse{}
	err := client.Request("vpc", "ModifyServiceTemplateGroupAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
