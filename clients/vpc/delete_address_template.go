package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除IP地址模板
// https://cloud.tencent.com/document/api/215/16712

type DeleteAddressTemplateRequest struct {
	// IP地址模板实例ID，例如：ipm-09o5m8kc。
	AddressTemplateId string `name:"AddressTemplateId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteAddressTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAddressTemplateResponse, error) {
	resp := &DeleteAddressTemplateResponse{}
	err := client.Request("vpc", "DeleteAddressTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteAddressTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
