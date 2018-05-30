package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建协议端口模板
// https://cloud.tencent.com/document/api/215/16710
type CreateServiceTemplateRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板名称
	ServiceTemplateName string `name:"ServiceTemplateName"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `name:"Services"`
}

func (req *CreateServiceTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateServiceTemplateResponse, error) {
	resp := &CreateServiceTemplateResponse{}
	err := client.Request("vpc", "CreateServiceTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateServiceTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 协议端口模板对象。
	ServiceTemplate ServiceTemplate `json:"ServiceTemplate"`
}
