package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建CCN
// https://cloud.tencent.com/document/api/215/19204

type CreateCcnRequest struct {
	// CCN描述信息，最大长度不能超过100个字节。
	CcnDescription *string `name:"CcnDescription,omitempty"`
	// CCN名称，最大长度不能超过60个字节。
	CcnName string `name:"CcnName"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateCcnRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateCcnResponse, error) {
	resp := &CreateCcnResponse{}
	err := client.Request("vpc", "CreateCcn", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateCcnResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云联网（CCN）对象。
	Ccn CCN `json:"Ccn"`
}
