package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建IP地址模版集合
// https://cloud.tencent.com/document/api/215/16709

type CreateAddressTemplateGroupRequest struct {
	// IP地址模版集合名称。
	AddressTemplateGroupName string `name:"AddressTemplateGroupName"`
	// IP地址模版实例ID，例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `name:"AddressTemplateIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateAddressTemplateGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAddressTemplateGroupResponse, error) {
	resp := &CreateAddressTemplateGroupResponse{}
	err := client.Request("vpc", "CreateAddressTemplateGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateAddressTemplateGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// IP地址模板集合对象。
	AddressTemplateGroup AddressTemplateGroup `json:"AddressTemplateGroup"`
}
