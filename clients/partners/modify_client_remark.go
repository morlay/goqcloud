package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改客户备注
// https://cloud.tencent.com/document/api/563/16045
type ModifyClientRemarkRequest struct {
	// 客户备注名称
	ClientRemark string `name:"ClientRemark"`
	// 客户账号ID
	ClientUin string `name:"ClientUin"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyClientRemarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyClientRemarkResponse, error) {
	resp := &ModifyClientRemarkResponse{}
	err := client.Request("partners", "ModifyClientRemark", "2018-03-21").Do(req, resp)
	return resp, err
}

type ModifyClientRemarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
