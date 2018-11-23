package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除实例账号
// https://cloud.tencent.com/document/api/238/19972

type DeleteAccountRequest struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 实例用户名数组
	UserNames []*string `name:"UserNames"`
}

func (req *DeleteAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAccountResponse, error) {
	resp := &DeleteAccountResponse{}
	err := client.Request("sqlserver", "DeleteAccount", "2018-03-28").Do(req, resp)
	return resp, err
}

type DeleteAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务流id
	FlowId int64 `json:"FlowId"`
}
