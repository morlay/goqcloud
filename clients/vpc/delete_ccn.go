package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除CCN
// https://cloud.tencent.com/document/api/215/19203

type DeleteCcnRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteCcnRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteCcnResponse, error) {
	resp := &DeleteCcnResponse{}
	err := client.Request("vpc", "DeleteCcn", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteCcnResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
