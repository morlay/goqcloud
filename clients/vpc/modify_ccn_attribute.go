package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改CCN属性
// https://cloud.tencent.com/document/api/215/19195

type ModifyCcnAttributeRequest struct {
	// CCN描述信息，最大长度不能超过100个字节。
	CcnDescription *string `name:"CcnDescription,omitempty"`
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// CCN名称，最大长度不能超过60个字节。
	CcnName *string `name:"CcnName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyCcnAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyCcnAttributeResponse, error) {
	resp := &ModifyCcnAttributeResponse{}
	err := client.Request("vpc", "ModifyCcnAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyCcnAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
