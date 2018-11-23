package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改账户备注
// https://cloud.tencent.com/document/api/238/19935

type ModifyAccountRemarkRequest struct {
	// 修改备注的账户信息
	Accounts []*AccountRemark `name:"Accounts"`
	// 实例ID，形如mssql-j8kv137v
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAccountRemarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountRemarkResponse, error) {
	resp := &ModifyAccountRemarkResponse{}
	err := client.Request("sqlserver", "ModifyAccountRemark", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyAccountRemarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
