package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建IP地址模版
// https://cloud.tencent.com/document/api/215/16708

type CreateAddressTemplateRequest struct {
	// IP地址模版名称
	AddressTemplateName string `name:"AddressTemplateName"`
	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `name:"Addresses"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateAddressTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAddressTemplateResponse, error) {
	resp := &CreateAddressTemplateResponse{}
	err := client.Request("vpc", "CreateAddressTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateAddressTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// IP地址模板对象。
	AddressTemplate AddressTemplate `json:"AddressTemplate"`
}
