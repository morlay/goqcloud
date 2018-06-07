package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除云数据库的账号
// https://cloud.tencent.com/document/api/236/17501

type DeleteAccountsRequest struct {
	// 云数据库账号。
	Accounts []*Account `name:"Accounts"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAccountsResponse, error) {
	resp := &DeleteAccountsResponse{}
	err := client.Request("cdb", "DeleteAccounts", "2017-03-20").Do(req, resp)
	return resp, err
}

type DeleteAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}
