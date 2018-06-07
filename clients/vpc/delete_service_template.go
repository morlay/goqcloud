package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除协议端口模板
// https://cloud.tencent.com/document/api/215/16714

type DeleteServiceTemplateRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板实例ID，例如：ppm-e6dy460g。
	ServiceTemplateId string `name:"ServiceTemplateId"`
}

func (req *DeleteServiceTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteServiceTemplateResponse, error) {
	resp := &DeleteServiceTemplateResponse{}
	err := client.Request("vpc", "DeleteServiceTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteServiceTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
