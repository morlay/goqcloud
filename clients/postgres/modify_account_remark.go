package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改帐号备注
// https://cloud.tencent.com/document/api/409/18108

type ModifyAccountRemarkRequest struct {
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId string `name:"DBInstanceId"`
	// 区域
	Region string `name:"Region"`
	// 用户UserName对应的新备注
	Remark string `name:"Remark"`
	// 实例用户名
	UserName string `name:"UserName"`
}

func (req *ModifyAccountRemarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountRemarkResponse, error) {
	resp := &ModifyAccountRemarkResponse{}
	err := client.Request("postgres", "ModifyAccountRemark", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAccountRemarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
