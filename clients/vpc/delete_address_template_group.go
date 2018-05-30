package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除IP地址模板集合
// https://cloud.tencent.com/document/api/215/16713
type DeleteAddressTemplateGroupRequest struct {
	// IP地址模板集合实例ID，例如：ipmg-90cex8mq。
	AddressTemplateGroupId string `name:"AddressTemplateGroupId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteAddressTemplateGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAddressTemplateGroupResponse, error) {
	resp := &DeleteAddressTemplateGroupResponse{}
	err := client.Request("vpc", "DeleteAddressTemplateGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteAddressTemplateGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
