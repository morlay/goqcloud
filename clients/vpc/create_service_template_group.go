package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建协议端口模板集合
// https://cloud.tencent.com/document/api/215/16711
type CreateServiceTemplateGroupRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板集合名称
	ServiceTemplateGroupName string `name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `name:"ServiceTemplateIds"`
}

func (req *CreateServiceTemplateGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateServiceTemplateGroupResponse, error) {
	resp := &CreateServiceTemplateGroupResponse{}
	err := client.Request("vpc", "CreateServiceTemplateGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateServiceTemplateGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 协议端口模板集合对象。
	ServiceTemplateGroup ServiceTemplateGroup `json:"ServiceTemplateGroup"`
}
