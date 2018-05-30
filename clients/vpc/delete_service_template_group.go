package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除协议端口模板集合
// https://cloud.tencent.com/document/api/215/16715
type DeleteServiceTemplateGroupRequest struct {
	// 区域
	Region string `name:"Region"`
	// 协议端口模板集合实例ID，例如：ppmg-n17uxvve。
	ServiceTemplateGroupId string `name:"ServiceTemplateGroupId"`
}

func (req *DeleteServiceTemplateGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteServiceTemplateGroupResponse, error) {
	resp := &DeleteServiceTemplateGroupResponse{}
	err := client.Request("vpc", "DeleteServiceTemplateGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteServiceTemplateGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
